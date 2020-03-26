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

// CardPrice is
type CardPrice struct {
	Card  collection.Card
	Price string
}

// ScryfallCard includes the relevant fields we're looking for when querying Scryfall
type ScryfallCard struct {
	Prices ScryfallPrices `json:"prices"`
}

// ScryfallPrices lets us access the prices from the json
type ScryfallPrices struct {
	USD     string `json:"usd"`
	USDFoil string `json:"usd_foil"`
}

// GetCardPrices iterates through a collection and writes each card's price into a channel
func GetCardPrices(cards []collection.Card, priceChannel chan CardPrice) {
	for _, card := range cards {
		time.Sleep(100 * time.Millisecond)

		go func(card collection.Card, priceChannel chan CardPrice) {
			price, err := getCardPrice(card)
			if err != nil {
				fmt.Printf("ERROR - Unable to get pricing for %v - %v.\n", card, err)
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
	var scryfallCard ScryfallCard

	log.Println("Querying for ", cardEncoded)
	response, err := http.Get(cardEncoded)
	if err != nil {
		return "", fmt.Errorf("query unsuccessful for card - %v", card.CardName)
	}

	data, _ := ioutil.ReadAll(response.Body)
	// log.Println("output is: ", string(data))
	err = json.Unmarshal(data, &scryfallCard)
	if err != nil {
		return "", fmt.Errorf("Couldn't read price information for card %v - %v", card.CardName, err)
	}

	if scryfallCard.Prices.USD != "" {
		// log.Println("Found value ", scryfallCard.Prices.USD)
		return scryfallCard.Prices.USD, nil
	}

	return scryfallCard.Prices.USDFoil, nil
}
