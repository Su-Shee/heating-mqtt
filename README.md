## How does it work?
* A little fake sensors publishes its (randomly generated within a range)
  temperature readings to mosquitto every couple of seconds.

* A service "in the middle" determines temperature by subscribing
  to temperature readings. If the temperature is above or below 22°C, it opens
  or closes the valve of the heater by publishing a level to adjust the valve.

* The fake heater is subscribed to the valve level and heats accordingly

I like to go by "make it work, make it right, make it fast/secure/pretty" - so
I've aimed for a working heating cycle above else and publishers and
subscribers actually reacting to each other via mosquitto as message broker.

Next steps are clearly "better code" in the sense of improved abstractions so
that publishing, subscribing, generating (or rather reading values in a more
realistic scenario) is more easily pluggable. I'm also not capturing any
errors properly and I'm not yet making good use of the asynchronous nature of
the entire scenario. I also didn't care for quality of service level or
message retaining; for the most part it's just "fire and forget".

Faking proper floating point data, the data conversion and JSON handling is
also quite horrible; a bigger prototype or test setups with lots of mocking
would definetely profit from a range of util/helper functions.

## Installation/Build
### manually:
1) docker run -it -p 1883:1883 --name=mosquitto  toke/mosquitto as recommended
2) go get dependencies for the eclipse "paho" mqtt library:
```
go get golang.org/x/net/websocket
go get golang.org/x/net/proxy
```
3) get the main library to handle MQTT:
```
go get github.com/eclipse/paho.mqtt.golang
```
4) "go build" in the service/, sensors/ and heater/ subdirectory
5) start ./service, ./sensors, ./heater in the respective subdirectory, it'll
just run.

### use the Makefile to build locally
``make deps`` to install the libraries
``make service`` to build the service
``make heater`` to build the heater
``make sensors`` to build the fake sensor binary


## More Ideas
* Arduino with a cheap temperature sensor + mqtt library
* Save data points in database (average etc, historical data)
* Listen on web socket for streaming data (charts etc) (not so great; a broker
  imho shouldn't be a service for websockets..)
* Don't use average as temperature measurements, but discard top 5% and bottom
  5% readings as outliers (scenario: in winter, the sun is shining on one
  sensor and another one is sitting close to a drafty window - readings to
  high/low)
* Make measurements more "sloppy" by accounting for the famous problems of
  distributed systems ("network is always available" ;)): add more cheap
  sensors, if one or more isn't working, doesn't matter, if one or more
  doesn't send data for a couple of minutes, doesn't matter either. solve
  problem of "good readings" with moar cheap sensors instead of more
  reliability.
* Set up a threshold how many sensor readings must exist (2 out of 4 sensors
  didn't send data - do we still want to heat up the room?)
* Make proper build environment for extra tiny client-side go (the binary
  which would actually sit in the sensors and send the readings)
* Decide between go modules or go dep whatever convention is used

## Nice Tools
* mosquitto_pub, mosquitto_sub
* go lib "mqtt" for low-level decoding of mqtt packages
* paho-go as "normal" library usage
* adafruit industries has a IoT cloud for non-commercial use
* mosquitto has a testing instance (with and without ssl)



