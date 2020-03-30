package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testingCardListLoc = "../config/exampleCardList.json"
	testingCardList    = `{
  "cards": [
    {
      "Quantity": 3,
      "CardName": "Mogis, God of Slaughter",
      "CardSet": "bng",
      "Deck": "Mogis"
    },
    {
      "CardName": "Alela, Artful Provocateur"
    },
    {
      "CardName": "Estrid, the Masked",
      "CardSet": "c18",
      "Foil": true
    }
  ]
}`
	testingExpectedCollection = Collection{
		[]Card{
			{
				Quantity: 3,
				CardName: "Mogis, God of Slaughter",
				CardSet:  "bng",
				Foil:     false,
				Deck:     "Mogis",
			},
			{
				Quantity: 1,
				CardName: "Alela, Artful Provocateur",
				CardSet:  "",
				Foil:     false,
			},
			{
				Quantity: 1,
				CardName: "Estrid, the Masked",
				CardSet:  "c18",
				Foil:     true,
			},
		},
	}
)

func TestGetCards(t *testing.T) {
	type testDefs struct {
		Name               string
		CardLocation       string
		ExpectPass         bool
		ExpectedError      string
		ExpectedCollection Collection
	}

	tests := []testDefs{
		testDefs{
			Name:               "Valid List",
			CardLocation:       testingCardListLoc,
			ExpectPass:         true,
			ExpectedCollection: testingExpectedCollection,
		},
		testDefs{
			Name:          "Invalid List",
			CardLocation:  "F@",
			ExpectPass:    false,
			ExpectedError: "couldn't read card list - couldn't read file from file location: F@",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cards, err := GetCards(test.CardLocation)

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
		ExpectedCollection Collection
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
			ExpectedError: "There was an error with parsing the collection - asdf",
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
			ExpectedError: "couldn't read file from file location: ",
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
