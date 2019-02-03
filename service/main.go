package main

import (
	"fmt"
  "os"
	"github.com/eclipse/paho.mqtt.golang"
)

var tempHandle = make(chan bool)

func TempHandler(client mqtt.Client, msg mqtt.Message) {
    tempHandle <- true
    fmt.Printf("temperature handler")
    fmt.Printf("[%s]  ", msg.Topic())
    fmt.Printf("%s\n", msg.Payload())
}

func main() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker("0.0.0.0:1883")
    opts.SetClientID("regulator")
    opts.SetCleanSession(true)

    client := mqtt.NewClient(opts)

    token := client.Connect();

    if token.Wait() && token.Error() != nil {
          fmt.Println(token.Error())
          os.Exit(1)
    }

    subscription := client.Subscribe("readings/temperature/#", 0, TempHandler)

    if subscription.Wait() && subscription.Error() != nil {
          fmt.Println(subscription.Error())
          os.Exit(1)
    }


    select {
        case <-tempHandle:
            fmt.Printf("temperature read")
    }

    client.Disconnect(250)
}


