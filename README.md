## How does it work?
* A little fake sensor publishes its (randomly generated within a range)
  temperature readings to mosquitto every couple of seconds.

* A service "in the middle" determines temperature by subscribing
  to temperature readings. If the temperature is above or below 22°C, it opens
  or closes the valve of the heater by publishing a level to adjust the valve.

* The fake heater is subscribed to the valve level and heats accordingly. ;)

I like to go by "make it work, make it right, make it fast/secure/pretty" - so
I've aimed for a working heating cycle above else and publishers and
subscribers actually reacting to each other via mosquitto as message broker. I
left out actually calculating valve openings by percentage; I'm just opening
the heater within mock-up range of temperature.

Next steps are clearly "better code" in the sense of improved abstractions so
that publishing, subscribing, generating (or rather reading values in a more
realistic scenario) is more easily pluggable. Also, just now (22:56) I've realized
that I could have made a channel instead of an endless loop...

The connection to the broker needs very loose and gracious retries and timeout
handling; otherwise using docker-compose (or Kubernetes) will mostly fail (no
build order, no waiting until a container has finished...)

The service in the middle could serve all kinds of things as a HTTP-based API:
a commandline client, web-based UIs, other tools...

I'm also not yet capturing any errors properly and I'm not yet making good use
of the asynchronous nature of the entire scenario. I also didn't care yet for
quality of service level or message retaining; for the most part it's just
"fire and forget".

Faking proper floating point data, the data conversion and JSON handling is
also quite horrible; a bigger prototype or test setups with lots of mocking
would definetely profit from a range of util/helper functions.

Also, finally deciding on a proper way how to balance the temperature readings
and valve opening with more rooms and many sensors so that it actually makes
sense. Add proper calculations for whatever scenario chosen.

A real prototype should have encrypted transport right from the start (SSL?
VPN? IPSec?).

Considering that we are talking about tiny sensors, IPv6 might actually be
really useful here.

## Installation/Build
### Manually:
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
4) run "go build" in the service/, sensors/ and heater/ subdirectory
5) start ``./service``, ``./sensors``, ``./heater`` in the respective
subdirectory, it'll just run. Watch them in three different terminal
publishing, subscribing and heating. :)

### Use the Makefile to build locally:
You'd still nedd the broker up and running:
``docker run -it -p 1883:1883 --name=mosquitto  toke/mosquitto`` as recommended!

``make deps`` to install the libraries

``make service`` to build the service

``make heater`` to build the heater

``make sensors`` to build the fake sensor binary

Lastly, start ``./service``, ``./sensors``, ``./heater`` in the respective
subdirectory, it'll just run. After that, you can just watch the three apps
subscribing, publishing and heating. :)


### Build and run via Docker containers
In order to allow connections to the broker properly, the containers need to
be started with --net=host:

```
docker run -it -p 1883:1883 -p 9001:9001 --name=mosquitto --net=host toke/mosquitto
```

The heating service containers can be build like this:

```
docker build -t service .
docker run -ti --net=host service <containerid>
```

```
docker build -t heater .
docker run -ti --net=host heater <containerid>
```

```
docker build -t sensors .
docker run -ti --net=host service <containerid>
```

If you start them in three different terminals, you can watch them publishing,
subscribing and heating. :)

There is also a basic docker-compose file available which at the moment has
the usal "distributed" problem: there is no quarantee which container gets
started and run in what order and they are not waiting for each other to
finish. The three tiny apps are not yet prepared with proper delayed retries
to actually wait themselves until they reach the broker.

(Deployment to Kubernetes has the same problem - there is however the concept
of a pre- and post hook in order to do something before and after a pod has
been set up properly. Use cases are e.g. database seeds which of course
require to actually HAVE a database before you can seed it.)

## More Ideas
* Arduinos with a cheap temperature sensor + mqtt library actually reading
  temperatures
* Save data points in database (nice calculations of mean, average, highest,
  lowest temp, historical data)
* Listen on web socket for streaming data (charts etc) (not so fond of that; a broker
  imho shouldn't be a service for websockets..)
* Don't use average as temperature measurements, but discard top 5% and bottom
  5% readings as outliers (scenario: in winter, the sun is shining on one
  sensor and another one is sitting close to a drafty window - readings too
  high/low)
* Make measurements more "sloppy" by accounting for the famous problems of
  distributed systems ("network is always available" ;)): Add more cheap
  sensors, if one or more isn't working, doesn't matter, if one or more
  doesn't send data for a couple of minutes, doesn't matter either. Solve
  problem of "good readings" with plenty of cheap sensors instead of more
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



