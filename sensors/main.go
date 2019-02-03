package main

import (
	"fmt"
  "time"
  "os"
  "encoding/json"
  "math/rand"
	"github.com/eclipse/paho.mqtt.golang"
)

type Payload struct {
  SensorID string  `json:"sensorID"`
  Type     string  `json:"type"`
  Value    float64 `json:"value"`
}

func generatePayload() []byte {
    rand_temp := float64(int((15.0 + rand.Float64() * 20.0) * 10.0)) / 10.0
    payload   := Payload{Value:    rand_temp,
                         Type:     "temperature",
                         SensorID: "sensor-1"}

    msg, _ := json.Marshal(payload)

    return msg
}

func main() {
	opts := mqtt.NewClientOptions()
  opts.AddBroker("0.0.0.0:1883")
  opts.SetClientID("sensor nr one")
	client := mqtt.NewClient(opts)

  token := client.Connect();

  if token.Wait() && token.Error() != nil {
      fmt.Println(token.Error())
      os.Exit(1)
	}

	for {
      msg := generatePayload()
      token := client.Publish("readings/temperature", 0, false, msg)
      token.Wait()
      time.Sleep(5 * time.Second)
	}

  client.Disconnect(250)
}


