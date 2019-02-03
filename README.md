## How does it work?
* four little fake sensors send their temperature readings to mosquitto every
  couple of seconds
* a service "in the middle" calculates the average temperature by subscribing
  to temperature readings. if the temperature is above or below 22°C, it opens
  or closes the valve of the heater by publishing a level to adjust the valve.
  everything which needs data or wants to change anything needs to talk to
  this service (cli tools, GUIs, tools..)
* the heater is subscribed to the valve level and heats accordingly
* the fake sensors fake reading the temperature and continue publishing their
  readings. :) they just do their thing continuously

## Installation/Build

## More Ideas
* arduino with a cheap temperature sensor + mqtt library
* save data points in database (average etc, historical data)
* listen on web socket for streaming data (charts etc) (not so great; a broker
  imho shouldn't be a service for websockets..)
* don't use average as temperature measurements, but discard top 5% and bottom
  5% readings as outliers (scenario: in winter, the sun is shining on one
  sensor and another one is sitting close to a drafty window - readings to
  high/low)
* make measurements more "sloppy" by accounting for the famous problems of
  distributed systems ("network is always available" ;)): add more cheap
  sensors, if one or more isn't working, doesn't matter, if one or more
  doesn't send data for a couple of minutes, doesn't matter either. solve
  problem of "good readings" with moar cheap sensors instead of more
  reliability.
* set up a threshold how many sensor readings must exist (2 out of 4 sensors
  didn't send data - do we still want to heat up the room?)
* make proper build environment for extra tiny client-side go (the binary
  which would actually sit in the sensors and send the readings)

## Nice Tools
* mosquitto_pub, mosquitto_sub
* go lib "mqtt" for low-level decoding of mqtt packages
* paho-go as "normal" library usage

## Deployment
* adafruit industries has a IoT cloud for non-commercial use
* mosquitto has a testing instance (with and without ssl)



