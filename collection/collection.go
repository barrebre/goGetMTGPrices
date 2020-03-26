package collection

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Card contains the shorthand definition of a card
type Card struct {
	CardName string
	CardSet  string
}

// GetCards takes a pointer to a file to read a list of cards
func GetCards(loc string) ([]Card, error) {
	cardStrings, err := readCardList(loc)
	if err != nil {
		return nil, fmt.Errorf("couldn't read card list - %v", err)
	}

	parsedCards, err := parseCardList(cardStrings)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse cards - %v", err)
	}

	return parsedCards, nil
}

func parseCardList(list string) ([]Card, error) {
	collection := []Card{}

	cards := strings.Split(list, "\n")
	for _, card := range cards {
		cardInfo := strings.Split(card, " - ")
		if len(cardInfo) != 2 {
			return nil, fmt.Errorf("couldnt read card %v", card)
		}
		card := Card{
			CardName: cardInfo[0],
			CardSet:  cardInfo[1],
		}
		collection = append(collection, card)
	}

	return collection, nil
}

func readCardList(loc string) (string, error) {
	dat, err := ioutil.ReadFile(loc)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}
