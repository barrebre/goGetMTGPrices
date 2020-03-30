package main

import (
	"log"
	"os"

	"github.com/barrebre/goGetMTGPrices/metrics"
	"github.com/barrebre/goGetMTGPrices/scheduler"
)

func main() {
	// Setup necessary metric components
	err := metrics.SetupMetrics()
	if err != nil {
		log.Println("FATAL - Couldn't set up metrics")
		os.Exit(1)
	}

	// Create the scheduler which writes to the priceChannel
	scheduler.StartLookupScheduler()

	log.Println("goGetMTGPrices started!")
}
