This app reads prices from Scryfall and sends them to a metrics data store, with New Relic and InfluxDB being supported.

# Environment Variables
* `goGetMTGPricesBackend`
    * NewRelic - sends data to New Relic. Requires `NEW_RELIC_INSIGHTS_INSERT_API_KEY`
    * Default - sends data to InfluxDB. This assumes the InfluxDB is available at `http://influxdb-service:8086`
* `NEW_RELIC_INSIGHTS_INSERT_API_KEY` - New Relic Insights API Key
* `runOnce` - Use this flag to only run the price check once

# Running this app
To run this app, either `go run main.go` or build a docker container, `docker build -t go-get-mtg-prices .` and run that `docker run go-get-mtg-prices:latest`.

# Requirement - Card List
The app requires a `cardList.json` file in the `config` folder. The formatting should be an array of cards in a `cards` value.

```
type Card struct {
	Quantity int `json:"quantity"`
	CardName string
	CardSet  string
	Foil     bool
	Deck     string
}
```

## Example
```
{
  "cards": [
    {
      "Quantity": 1, 
      "CardName": "Adriana, Captain of the Guard", 
      "CardSet": "C20"
    },
    {
      "Quantity": 1, 
      "CardName": "Aegis of the Gods", 
      "CardSet": "JOU",
      "Foil": true
    }
  ]
}
```

An example card list can be found in the `config` folder.

# Using a Local Influx
## Set up Hosts file
sudo vi /etc/hosts
127.0.0.1 influxdb-service

## Start an Influx
docker run -d -expose -p 8086:8086 --name influxdb -e INFLUXDB_ADMIN_PASSWORD=PASSWORD influxdb/influxdb

## Setting up DB
curl -G http://influxdb-service:8086/query --data-urlencode "q=CREATE DATABASE CardPrices"