package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetCards takes a pointer to a file to read a list of cards
func GetCards(loc string) (Collection, error) {
	cardStrings, err := readCardList(loc)
	if err != nil {
		return Collection{}, fmt.Errorf("couldn't read card list - %v", err.Error())
	}

	parsedCards, err := parseCardList(cardStrings)
	if err != nil {
		return Collection{}, fmt.Errorf("couldn't parse cards - %v", err.Error())
	}

	return parsedCards, nil
}

func readCardList(loc string) (string, error) {
	dat, err := ioutil.ReadFile(loc)
	// log.Println("data string: ", string(dat), ". error: ", err.Error())
	if err != nil {
		return "", fmt.Errorf("couldn't read file from file location: %v", loc)
	}

	return string(dat), nil
}

func parseCardList(list string) (Collection, error) {
	var collection Collection
	json.Unmarshal([]byte(list), &collection)

	if len(collection.Cards) == 0 {
		return Collection{}, fmt.Errorf("There was an error parsing the collection - %v", list)
	}

	return standardizeCards(collection), nil
}

func standardizeCards(parsedCards Collection) Collection {
	standardizedCards := Collection{[]Card{}}

	var tempCard Card
	for _, card := range parsedCards.Cards {
		tempCard = MakeCard(card.Quantity, card.CardName, card.CardSet, card.Foil, card.Deck)
		standardizedCards.Cards = append(standardizedCards.Cards, tempCard)
	}

	return standardizedCards
}
