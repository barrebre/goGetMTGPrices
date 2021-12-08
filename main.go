package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

	runOnce, err := strconv.ParseBool(os.Getenv("runOnce"))
	if err != nil {
		fmt.Printf("WARN - Couldn't parse runOnce env variable. runOnce: %v. error: %v.\n", runOnce, err)
	}

	if !runOnce {
		// Create the scheduler which writes to the priceChannel
		scheduler.StartLookupScheduler()

		log.Println("goGetMTGPrices started!")
	} else {
		scheduler.RunLookup()
	}
}
