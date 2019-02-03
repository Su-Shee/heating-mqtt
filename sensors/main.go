package main

import (
	"fmt"

	"github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions()
  opts.AddBroker("0.0.0.0:1883")
  opts.SetClientID("sensor nr one")

	client := mqtt.NewClient(opts)

  token := client.Connect();

  if token.Wait() && token.Error() != nil {

		panic(token.Error())
	}

	for i := 0; i < 5; i++ {
		msg   := fmt.Sprintf("19", i)
		token := client.Publish("readings/temperature", 0, false, msg)
		token.Wait()
	}

}


