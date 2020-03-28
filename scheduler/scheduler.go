package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/barrebre/goGetMTGPrices/collection"
	"github.com/barrebre/goGetMTGPrices/metrics"
	"github.com/barrebre/goGetMTGPrices/prices"
)

const (
	// rateLimiterTime is how frequently we query for new data
	// Scryfall says they only update once per day
	// rateLimiterTime = 86400000 * time.Millisecond

	// Overwritten to every 5 mins for now, anyway
	rateLimiterTime = 300000 * time.Millisecond
)

// CreatePriceReader sets up all necessary hooks into the PriceReader chan
func CreatePriceReader(prices chan prices.CardPrice) {
	metrics.SendPriceMetrics(prices)
}

// StartLookupScheduler starts the infinite loop which queries prices after the rateLimiter
func StartLookupScheduler(priceChannel chan prices.CardPrice) {
	for {
		err := lookup(priceChannel)
		if err != nil {
			log.Printf("Error running at %v - %v.\n", time.Now(), err.Error())
		} else {
			log.Printf("Successful run at %v.\n", time.Now())
		}

		// wait for the desired time to query again
		time.Sleep(rateLimiterTime)
	}
}

// we pass the priceChannel we want to write to
func lookup(priceChannel chan prices.CardPrice) error {
	cards, err := collection.GetCards("config/cardList.json")
	if err != nil {
		return fmt.Errorf("couldn't read card list - %v", err.Error())
	}

	prices.GetCardPrices(cards, priceChannel)
	if err != nil {
		return fmt.Errorf("couldn't read price from scryfall - %v", err.Error())
	}

	return nil
}
