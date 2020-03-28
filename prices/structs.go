package prices

import "github.com/barrebre/goGetMTGPrices/collection"

//// Definitions

// CardPrice contains a card name, set, and price
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

//// Accessor

// MakeCardPrice makes a CardPrice object
func MakeCardPrice(card collection.Card, price string) CardPrice {
	return CardPrice{
		Card:  card,
		Price: price,
	}
}

//// Example Accessors

// MakeExampleCardPrice creates an example CardPrice object
func MakeExampleCardPrice() CardPrice {
	return MakeCardPrice(collection.MakeExampleCard(), "7.21")
}
