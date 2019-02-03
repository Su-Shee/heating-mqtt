deps:
	go get golang.org/x/net/websocket
	go get golang.org/x/net/proxy
	go get github.com/eclipse/paho.mqtt.golang

service:
	go build -o service/service service/*

heater:
	go build -o heater/heater heater/*

sensors:
	go build -o sensors/sensors sensors/*

.PHONY: deps heater service sensors
