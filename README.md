## How does it work?
* A little fake sensors sends its temperature readings to mosquitto every
  couple of seconds
* A service "in the middle" determines temperature by subscribing
  to temperature readings. If the temperature is above or below 22°C, it opens
  or closes the valve of the heater by publishing a level to adjust the valve.
* The heater is subscribed to the valve level and heats accordingly
* The fake sensors fake reading the temperature and continue publishing their
  readings. :) They just do their thing continuously.

## Installation/Build

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

## Nice Tools
* mosquitto_pub, mosquitto_sub
* go lib "mqtt" for low-level decoding of mqtt packages
* paho-go as "normal" library usage

## Deployment
* adafruit industries has a IoT cloud for non-commercial use
* mosquitto has a testing instance (with and without ssl)



