FROM golang:latest

RUN go get golang.org/x/net/websocket
RUN go get golang.org/x/net/proxy
RUN go get github.com/eclipse/paho.mqtt.golang

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o heater .
CMD ["/app/heater"]
