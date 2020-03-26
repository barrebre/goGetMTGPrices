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

// GetCardPrices asdf
func GetCardPrices(cards []collection.Card, priceChannel chan CardPrice) ([]CardPrice, error) {
	prices := []CardPrice{}

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

	return prices, nil
}

// ScryfallCard describes
type ScryfallCard struct {
	Prices ScryfallPrices `json:"prices"`
}

// ScryfallPrices asdf
type ScryfallPrices struct {
	USD     string `json:"usd"`
	USDFoil string `json:"usd_foil"`
}

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
