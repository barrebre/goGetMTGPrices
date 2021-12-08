package metrics

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/barrebre/goGetMTGPrices/collection"
	"github.com/barrebre/goGetMTGPrices/prices"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
)

// SetupMetrics sets up reqs for sending metrics to influx
func SetupMetrics() error {
	switch backend := os.Getenv("goGetMTGPricesBackend"); backend {
	case "NewRelic":
		log.Println("Using NR backend")
		h, err := telemetry.NewHarvester(telemetry.ConfigAPIKey(os.Getenv("NEW_RELIC_INSIGHTS_INSERT_API_KEY")))
		if err != nil {
			return fmt.Errorf("error creating harvester: %s", err.Error())
		}

		NRReadPriceMetrics(h)
	default:
		log.Printf("Using backend: %s\n", backend)
		client, err := createInfluxClient()
		if err != nil {
			return fmt.Errorf("error creating Influx client - %v", err)
		}

		bp, err := createBatchPointsStruct(client)
		if err != nil {
			return fmt.Errorf("couldn't create batch points - %v", err)
		}

		InfluxReadPriceMetrics(bp)
	}
	return nil
}

// NRReadPriceMetrics reads the prices channel and sends the info to NR
func NRReadPriceMetrics(h *telemetry.Harvester) {
	prices := prices.GetPriceChannel()
	go func() {
		loc, _ := time.LoadLocation("America/Chicago")

		for {
			price := <-*prices

			metric := telemetry.Summary{
				Name:       "Card Price",
				Attributes: getCardTagsInterface(price),
				Count:      float64(price.Card.Quantity),
				Max:        price.Price,
				Timestamp:  time.Now().In(loc),
			}

			// log.Printf("Created metric: %v", spew.Sdump(metric))

			h.RecordMetric(metric)
		}
	}()
}

// InfluxReadPriceMetrics reads the prices channel and sends the info to influx
func InfluxReadPriceMetrics(bp *batchPointsStruct) {
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

	return pt, nil
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

// Builds the Tags for a Card
func getCardTagsInterface(price prices.CardPrice) map[string]interface{} {
	foilCard := isFoil(price.Card)

	// Create a point and add to batch
	return map[string]interface{}{
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
	fields := map[string]interface{}{
		"value":      price.Price,
		"quantity":   price.Card.Quantity,
		"totalValue": price.Price * float64(price.Card.Quantity),
	}

	return fields, nil
}
