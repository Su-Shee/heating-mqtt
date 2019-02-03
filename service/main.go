package main

import (
	"fmt"
  "os"
  "time"
  "encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
)

var tempHandle = make(chan float64)

func TempHandler(client mqtt.Client, msg mqtt.Message) {

    var temp map[string]interface{}
    _ = json.Unmarshal([]byte(msg.Payload()), &temp)

    val := temp["value"].(float64)
    fmt.Printf("temperature read: %s\n", msg.Payload())
    tempHandle <- val
}

type valveLevel struct {
  Level int `json:"level"`
}

func generateValveLevel(agv_temp float64) []byte {
    level := 0
    fmt.Printf("setting level for temp: %.1f\n", agv_temp)
    switch {
        case agv_temp < 18.0:
          level = 75
        case agv_temp < 22.0 && agv_temp > 18.0:
          level = 50
        default:
          level = 0
    }

    fmt.Printf("setting level for valve at %d!\n", level)
    pl_level := valveLevel{Level: level}
    msg, _   := json.Marshal(pl_level)
    return msg
}

func main() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker("0.0.0.0:1883")
    opts.SetClientID("regulator")
    opts.SetCleanSession(false)

    client := mqtt.NewClient(opts)

    token := client.Connect();

    if token.Wait() && token.Error() != nil {
          fmt.Println(token.Error())
          os.Exit(1)
    }


    subscription := client.Subscribe("readings/temperature/#", 1, TempHandler)

    if subscription.Wait() && subscription.Error() != nil {
          fmt.Println(subscription.Error())
          os.Exit(1)
    }

    for {
        select {
            case temp := <-tempHandle:
            level := generateValveLevel(temp)
            publication := client.Publish("actuators/room-1", 0, false, level)
            publication.Wait()
            time.Sleep(5 * time.Second)
        }

    }

    client.Disconnect(250)
}


