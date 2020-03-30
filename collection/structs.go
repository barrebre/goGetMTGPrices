package collection

//// Definitions

// Card contains the shorthand definition of a card
type Card struct {
	Quantity int `json:"quantity"`
	CardName string
	CardSet  string
	Foil     bool
	Deck     string
}

// Collection contains an array of cards
type Collection struct {
	Cards []Card `json:"cards"`
}

//// Accessors

// MakeCard creates a card with the given parameters, otherwise defaults
func MakeCard(quantity int, cardName string, cardSet string, foil bool, deck string) Card {
	newCard := Card{
		Quantity: 1,
		CardName: "",
		CardSet:  "",
		Foil:     false,
		Deck:     "",
	}

	if quantity != 0 {
		newCard.Quantity = quantity
	}
	if foil == true {
		newCard.Foil = true
	}

	newCard.CardName = cardName
	newCard.CardSet = cardSet
	newCard.Deck = deck

	return newCard
}

// MakeCollection creates a collection with the given parameters
func MakeCollection(cards []Card) Collection {
	return Collection{cards}
}

//// Example Accessors

// MakeDefaultCard returns an example card
func MakeDefaultCard() Card {
	return MakeCard(1, "", "", false, "")
}

// MakeExampleCard returns an example card
func MakeExampleCard() Card {
	return MakeCard(1, "Mogis, God of Slaughter", "bng", false, "Mogis")
}

// MakeExampleCollection returns an example collection of cards
func MakeExampleCollection() Collection {
	return MakeCollection([]Card{MakeExampleCard(), MakeExampleCard(), MakeExampleCard()})
}

// MakeInvalidExampleCollection returns an example collection of invalid cards
func MakeInvalidExampleCollection() Collection {
	invalidCard := MakeCard(1, "Brett Bretterson", "bab", false, "")

	return Collection{[]Card{invalidCard, invalidCard}}
}
