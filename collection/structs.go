package collection

//// Definitions

// Card contains the shorthand definition of a card
type Card struct {
	CardName string
	CardSet  string
}

//// Example Values
var (
	exampleCard = Card{
		CardName: "Mogis, God of Slaughter",
		CardSet:  "bng",
	}

	exampleCollection = []Card{exampleCard, exampleCard, exampleCard}
)

//// Example Accessors

// GetExampleCollection returns an example collection of cards
func GetExampleCollection() []Card {
	return exampleCollection
}

// GetInvalidExampleCollection returns an example collection of invalid cards
func GetInvalidExampleCollection() []Card {
	invalidCard := Card{
		CardName: "Brett Bretterson",
		CardSet:  "bab",
	}

	return []Card{invalidCard, invalidCard}
}
