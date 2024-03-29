package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/barrebre/goGetMTGPrices/collection"
	"github.com/barrebre/goGetMTGPrices/prices"
)

const (
	// rateLimiterTime is how frequently we query for new data
	// Scryfall says they only update once per day
	// rateLimiterTime = 86400000 * time.Millisecond

	// Overwritten to every 5 mins for now, anyway
	rateLimiterTime = 300000 * time.Millisecond
)

// StartLookupScheduler starts the infinite loop which queries prices after the rateLimiter
func StartLookupScheduler() {
	for {
		RunLookup()

		// wait for the desired time to query again
		time.Sleep(rateLimiterTime)
	}
}

// RunLookup actually performs the price gathering
func RunLookup() {
	err := lookupCardsAndPrices()
	if err != nil {
		log.Printf("Error running at %v - %v.\n", time.Now(), err.Error())
	} else {
		log.Printf("Successful run at %v.\n", time.Now())
	}
}

// we pass the priceChannel we want to write to
func lookupCardsAndPrices() error {
	var cards collection.Collection
	cards, err := collection.GetCards("config/cardList.json")
	if err != nil {
		log.Printf("INFO - Couldn't find real card source. Using example card list.")
		cards, err = collection.GetCards("config/exampleCardList.json")
		if err != nil {
			return fmt.Errorf("couldn't read card list - %v", err.Error())
		}
	}

	prices.GetCardPrices(cards)
	if err != nil {
		return fmt.Errorf("couldn't read price from scryfall - %v", err.Error())
	}

	return nil
}
