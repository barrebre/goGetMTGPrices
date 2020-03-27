package metrics

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/barrebre/goGetMTGPrices/prices"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	myDB = "CardPrices"
)

// SendPriceMetrics reads the prices channel and sends the info to influx
func SendPriceMetrics(prices chan prices.CardPrice) {
	go func() {
		for {
			price := <-prices
			log.Printf("Read in price: %v.\n", price)

			err := sendInfluxData(price)
			if err != nil {
				log.Printf("Couldn't send price info to influx for %v.\n", err.Error())
			}
		}
	}()
}

func sendInfluxData(price prices.CardPrice) error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://influxdb-service:8086",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a point and add to batch
	tags := map[string]string{
		"cardName": price.Card.CardName,
		"cardSet":  price.Card.CardSet,
	}

	priceFloat, err := strconv.ParseFloat(price.Price, 2)
	if err != nil {
		return fmt.Errorf("couldn't parse pricing data into float - %v", err.Error())
	}
	fields := map[string]interface{}{
		"value": priceFloat,
	}

	pt, err := client.NewPoint("price", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err.Error())
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
