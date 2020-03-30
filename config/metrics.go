package config

import (
	"github.com/barrebre/goGetMTGPrices/prices"
)

var (
	priceChannel chan prices.CardPrice
)

// GetPriceChannel returns a channel to write prices to
func GetPriceChannel() *chan prices.CardPrice {
	if priceChannel != nil {
		// log.Println("INFO - priceChannel exists")
		return &priceChannel
	}

	// log.Println("INFO - Making new priceChannel")
	priceChannel = make(chan prices.CardPrice)
	return &priceChannel
}
