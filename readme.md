This app reads prices from Scryfall and sends them to a metrics data store, with New Relic and InfluxDB being supported.

# Environment Variables
* `goGetMTGPricesBackend`
    * NewRelic - sends data to New Relic. Requires `NEW_RELIC_INSIGHTS_INSERT_API_KEY`
    * Default - sends data to InfluxDB. This assumes the InfluxDB is available at `http://influxdb-service:8086`
* `NEW_RELIC_INSIGHTS_INSERT_API_KEY` - New Relic Insights API Key
* `runOnce` - Use this flag to only run the price check once

## Set up Hosts file
sudo vi /etc/hosts
127.0.0.1 influxdb-service

## Start an Influx
docker run -d -expose -p 8086:8086 --name influxdb -e INFLUXDB_ADMIN_PASSWORD=PASSWORD influxdb/influxdb

## Setting up DB
curl -G http://influxdb-service:8086/query --data-urlencode "q=CREATE DATABASE CardPrices"

## Create the K8s Configmap
kubectl create configmap card-list-configmap --from-file=cardList.json