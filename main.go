package main

import (
	"github.com/barrebre/goGetMTGPrices/prices"
	"github.com/barrebre/goGetMTGPrices/scheduler"
)

func main() {
	// Make a channel to write prices to
	priceChannel := make(chan prices.CardPrice)

	// Create the reader which consumes from the priceChannel
	scheduler.CreatePriceReader(priceChannel)

	// Create the scheduler which writes to the priceChannel
	scheduler.StartLookupScheduler(priceChannel)
}
