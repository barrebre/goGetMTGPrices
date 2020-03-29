package main

import (
	"log"
	"os"

	"github.com/barrebre/goGetMTGPrices/metrics"
	"github.com/barrebre/goGetMTGPrices/prices"
	"github.com/barrebre/goGetMTGPrices/scheduler"
)

func main() {
	// Make a channel to write prices to
	priceChannel := make(chan prices.CardPrice)

	// Setup necessary metric components
	err := metrics.SetupMetrics(priceChannel)
	if err != nil {
		log.Println("FATAL - Couldn't set up metrics")
		os.Exit(1)
	}

	// Create the scheduler which writes to the priceChannel
	scheduler.StartLookupScheduler(priceChannel)

	log.Println("goGetMTGPrices started!")
}
