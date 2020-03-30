package metrics

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/barrebre/goGetMTGPrices/collection"
	"github.com/barrebre/goGetMTGPrices/prices"

	influx "github.com/influxdata/influxdb/client/v2"
)

// SetupMetrics sets up reqs for sending metrics to influx
func SetupMetrics() error {
	client, err := createInfluxClient()
	if err != nil {
		return fmt.Errorf("error creating Influx client - %v", err)
	}

	bp, err := createBatchPointsStruct(client)
	if err != nil {
		return fmt.Errorf("couldn't create batch points - %v", err)
	}

	ReadPriceMetrics(bp)
	return nil
}

// ReadPriceMetrics reads the prices channel and sends the info to influx
func ReadPriceMetrics(bp *batchPointsStruct) {
	prices := prices.GetPriceChannel()
	go func() {
		for {
			price := <-*prices
			// log.Printf("Read in price: %v.\n", price)

			pt, err := createCardInfluxDataPoint(price)
			if err != nil {
				log.Printf("Couldn't send price info to influx for %v.\n", err.Error())
			}

			bp.addPoint(*pt)
		}
	}()
}

func createCardInfluxDataPoint(price prices.CardPrice) (*influx.Point, error) {
	tags := getCardTags(price)
	fields, err := getCardFields(price)
	if err != nil {
		return &influx.Point{}, fmt.Errorf("Couldn't build fields - %v", price)
	}

	pt, err := influx.NewPoint("price", tags, fields, time.Now())
	if err != nil {
		return &influx.Point{}, fmt.Errorf("couldn't create influx point - %v", err)
	}
	// log.Printf("Made influx point: %v.\n", spew.Sdump(pt))

	return pt, nil
}

// Sends the metric data to Influx
func createInfluxDataPoint(price prices.CardPrice, client influx.Client, bp influx.BatchPoints) error {
	// Write the batch
	if err := client.Write(bp); err != nil {
		log.Fatal(err.Error())
	}

	// Close client resources
	if err := client.Close(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

// Builds the Tags for a Card
func getCardTags(price prices.CardPrice) map[string]string {
	foilCard := isFoil(price.Card)

	// Create a point and add to batch
	return map[string]string{
		"cardName": price.Card.CardName,
		"cardSet":  price.Card.CardSet,
		"foil":     foilCard,
		"quantity": fmt.Sprintf("%v", price.Card.Quantity),
		"deck":     price.Card.Deck,
	}
}

// Check if foil
func isFoil(card collection.Card) string {
	foilCard := "false"
	if card.Foil {
		foilCard = "true"
	}

	return foilCard
}

// Builds the Fields for a card
func getCardFields(price prices.CardPrice) (map[string]interface{}, error) {
	// Turn the price into a float
	priceFloat, err := strconv.ParseFloat(price.Price, 2)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse pricing data into float - %v", err.Error())
	}

	fields := map[string]interface{}{
		"value":      priceFloat,
		"quantity":   price.Card.Quantity,
		"totalValue": priceFloat * float64(price.Card.Quantity),
	}

	return fields, nil
}
