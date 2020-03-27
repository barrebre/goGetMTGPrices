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
