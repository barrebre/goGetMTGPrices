package collection

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GetCards takes a pointer to a file to read a list of cards
func GetCards(loc string) ([]Card, error) {
	cardStrings, err := readCardList(loc)
	if err != nil {
		return nil, fmt.Errorf("couldn't read card list - %v", err.Error())
	}

	parsedCards, err := parseCardList(cardStrings)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse cards - %v", err.Error())
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
	// log.Println("data string: ", string(dat), ". error: ", err.Error())
	if err != nil {
		return "", fmt.Errorf("couldn't read file from file location: %v", loc)
	}

	return string(dat), nil
}
