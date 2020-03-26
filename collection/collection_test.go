package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testingCardListLoc        = "../example/cardList.txt"
	testingCardList           = "Mogis, God of Slaughter - bng\nAlela, Artful Provocateur - eld\nEstrid, the Masked - c18"
	testingExpectedCollection = []Card{
		{
			CardName: "Mogis, God of Slaughter",
			CardSet:  "bng",
		},
		{
			CardName: "Alela, Artful Provocateur",
			CardSet:  "eld",
		},
		{
			CardName: "Estrid, the Masked",
			CardSet:  "c18",
		},
	}
)

func TestGetCards(t *testing.T) {
	type testDefs struct {
		Name               string
		CardLocation       string
		ExpectPass         bool
		ExpectedError      string
		ExpectedCollection []Card
	}

	tests := []testDefs{
		testDefs{
			Name:               "Valid List",
			CardLocation:       testingCardListLoc,
			ExpectPass:         true,
			ExpectedCollection: testingExpectedCollection,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cards, err := GetCards(testingCardListLoc)

			if test.ExpectPass {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedCollection, cards)
			} else {
				assert.EqualError(t, err, test.ExpectedError)
			}
		})
	}
}

func TestParseCardList(t *testing.T) {
	type testDefs struct {
		Name               string
		CardList           string
		ExpectPass         bool
		ExpectedError      string
		ExpectedCollection []Card
	}

	tests := []testDefs{
		testDefs{
			Name:               "Valid List",
			CardList:           testingCardList,
			ExpectPass:         true,
			ExpectedCollection: testingExpectedCollection,
		},
		testDefs{
			Name:          "Invalid List",
			CardList:      "asdf",
			ExpectPass:    false,
			ExpectedError: "couldnt read card asdf",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			collection, err := parseCardList(test.CardList)

			if test.ExpectPass {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedCollection, collection)
			} else {
				assert.EqualError(t, err, test.ExpectedError)
			}
		})
	}
}

func TestReadCardList(t *testing.T) {
	type testDefs struct {
		Name          string
		ExpectedList  string
		ExpectPass    bool
		ExpectedError string
		FileLocation  string
	}

	tests := []testDefs{
		testDefs{
			Name:         "Expected Pass",
			ExpectedList: testingCardList,
			ExpectPass:   true,
			FileLocation: testingCardListLoc,
		},
		testDefs{
			Name:          "Missing File",
			ExpectPass:    false,
			ExpectedError: "open : no such file or directory",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cardList, err := readCardList(test.FileLocation)
			if test.ExpectPass {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedList, cardList)
			} else {
				assert.EqualError(t, err, test.ExpectedError)
			}
		})
	}
}
