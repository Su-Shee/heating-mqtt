package main

import (
	"fmt"
  "os"
  "encoding/json"
  "time"
	"github.com/eclipse/paho.mqtt.golang"
)

var valveHandle = make(chan int)

func ValveHandler(client mqtt.Client, msg mqtt.Message) {

    var temp map[string]interface{}
    _ = json.Unmarshal([]byte(msg.Payload()), &temp)

    //fmt.Printf("%s\n", msg.Payload())
    val := int(temp["level"].(float64))
    fmt.Printf("valve opening level requested: %d\n", val)
    valveHandle <- val
}

func main() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker("0.0.0.0:1883")
    opts.SetClientID("heater")

    client := mqtt.NewClient(opts)

    token := client.Connect();

    if token.Wait() && token.Error() != nil {
          fmt.Println(token.Error())
          os.Exit(1)
    }

    subscription := client.Subscribe("actuators/room-1/#", 0, ValveHandler)

    if subscription.Wait() && subscription.Error() != nil {
          fmt.Println(subscription.Error())
          os.Exit(1)
    }

    for {
      select {
          case valve := <-valveHandle:
              fmt.Println("HEATING with valve open at: ", valve)
              time.Sleep(5 * time.Second)
        }
    }

    client.Disconnect(250)
}


