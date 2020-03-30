package metrics

import (
	"github.com/barrebre/goGetMTGPrices/prices"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetCardTags(t *testing.T) {
	type testDefs struct {
		Name          string
		Price         prices.CardPrice
		ExpectPass    bool
		ExpectedTags  map[string]string
		ExpectedError string
	}

	tests := []testDefs{
		testDefs{
			Name:         "Valid",
			Price:        prices.MakeExampleCardPrice(),
			ExpectPass:   true,
			ExpectedTags: map[string]string{"cardName": "Mogis, God of Slaughter", "cardSet": "bng", "deck": "Mogis", "foil": "false", "quantity": "1"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tags := getCardTags(test.Price)

			assert.Equal(t, tags, test.ExpectedTags)
		})
	}
}

func TestGetCardFields(t *testing.T) {
	type testDefs struct {
		Name           string
		Price          prices.CardPrice
		ExpectPass     bool
		ExpectedFields map[string]interface{}
		ExpectedError  string
	}

	tests := []testDefs{
		testDefs{
			Name:           "Valid",
			Price:          prices.MakeExampleCardPrice(),
			ExpectPass:     true,
			ExpectedFields: map[string]interface{}{"quantity": 1, "totalValue": 7.21, "value": 7.21},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			fields, err := getCardFields(test.Price)

			if test.ExpectPass {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedFields, fields)
			} else {
				assert.EqualError(t, err, test.ExpectedError)
			}
		})
	}
}
