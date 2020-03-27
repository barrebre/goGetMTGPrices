package collection

//// Definitions

// Card contains the shorthand definition of a card
type Card struct {
	CardName string
	CardSet  string
	Foil     bool
}

// Collection contains an array of cards
type Collection struct {
	Cards []Card `json:"cards"`
}

//// Accessors

// MakeCard creates a card with the given parameters
func MakeCard(cardName string, cardSet string, foil bool) Card {
	return Card{
		CardName: cardName,
		CardSet:  cardSet,
		Foil:     foil,
	}
}

// MakeCollection creates a collection with the given parameters
func MakeCollection(cards []Card) Collection {
	return Collection{cards}
}

//// Example Accessors

// MakeExampleCard returns an example card
func MakeExampleCard() Card {
	return MakeCard("Mogis, God of Slaughter", "bng", false)
}

// MakeExampleCollection returns an example collection of cards
func MakeExampleCollection() Collection {
	return MakeCollection([]Card{MakeExampleCard(), MakeExampleCard(), MakeExampleCard()})
}

// MakeInvalidExampleCollection returns an example collection of invalid cards
func MakeInvalidExampleCollection() Collection {
	invalidCard := MakeCard("Brett Bretterson", "bab", false)

	return Collection{[]Card{invalidCard, invalidCard}}
}
