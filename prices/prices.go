package prices

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/barrebre/goGetMTGPrices/collection"
)

const (
	queryPricingDelay = 100 * time.Millisecond
)

var (
	priceChannel chan CardPrice
)

// GetPriceChannel returns a channel to write prices to
func GetPriceChannel() *chan CardPrice {
	if priceChannel != nil {
		// log.Println("INFO - priceChannel exists")
		return &priceChannel
	}

	// log.Println("INFO - Making new priceChannel")
	priceChannel = make(chan CardPrice)
	return &priceChannel
}

// GetCardPrices iterates through a collection and writes each card's price into a channel
func GetCardPrices(cards collection.Collection) {
	priceChannel := *GetPriceChannel()

	for _, card := range cards.Cards {
		time.Sleep(queryPricingDelay)

		go func(card collection.Card, priceChannel chan CardPrice) {
			price, err := getCardPrice(card)
			if err != nil {
				log.Printf("ERROR - Unable to get pricing for %v - %v.\n", card, err.Error())
			} else {
				newCardPrice := CardPrice{
					Card:  card,
					Price: price,
				}

				priceChannel <- newCardPrice
			}
		}(card, priceChannel)
	}
}

// getCardPrice is what actually queries Scryfall to get the current price
func getCardPrice(card collection.Card) (string, error) {
	cardNameEscaped := url.QueryEscape(card.CardName)
	cardEncoded := fmt.Sprintf("https://api.scryfall.com/cards/named?exact=%s&set=%s", cardNameEscaped, card.CardSet)

	// log.Println("Querying for ", cardEncoded)
	response, err := http.Get(cardEncoded)
	if err != nil {
		return "", fmt.Errorf("query unsuccessful for card - %v", card.CardName)
	}

	data, _ := ioutil.ReadAll(response.Body)
	// log.Println("output is: ", string(data))

	var scryfallCard ScryfallCard
	err = json.Unmarshal(data, &scryfallCard)
	if err != nil {
		return "", fmt.Errorf("Couldn't read price information for card %v - %v", card.CardName, err.Error())
	}

	if card.Foil && scryfallCard.Prices.USDFoil != "" {
		return scryfallCard.Prices.USDFoil, nil
	} else if scryfallCard.Prices.USD != "" {
		return scryfallCard.Prices.USD, nil
	}

	return "", fmt.Errorf("Couldn't get the price from Scryfall")
}
