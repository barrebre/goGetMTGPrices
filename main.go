package main

import (
	"fmt"
	"log"
	"time"

	"github.com/barrebre/goGetMTGPrices/collection"
	"github.com/barrebre/goGetMTGPrices/prices"
)

var (
	rateLimiterTime = 86400000 * time.Millisecond
)

// Create the paths to access the APIs
func main() {
	priceChannel := make(chan prices.CardPrice)
	createChannelReader(priceChannel)

	for {
		err := scheduledLookup(priceChannel)
		if err != nil {
			log.Printf("Error running at %v - %v.\n", time.Now(), err)
		} else {
			log.Printf("Successful run at %v.\n", time.Now())
		}
		time.Sleep(rateLimiterTime)
	}
}

func createChannelReader(prices chan prices.CardPrice) {
	go func() {
		for {
			price := <-prices
			log.Printf("Read in price: %v.\n", price)
		}
	}()
}

func scheduledLookup(priceChannel chan prices.CardPrice) error {
	cards, err := collection.GetCards("example/cardList.txt")
	if err != nil {
		return fmt.Errorf("couldn't read card list - %v", err)
	}

	prices, err := prices.GetCardPrices(cards, priceChannel)
	if err != nil {
		return fmt.Errorf("couldn't read price from scryfall - %v", err)
	}

	fmt.Println(prices)

	return nil
}
