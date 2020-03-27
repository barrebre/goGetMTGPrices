package prices

import (
	"testing"

	"github.com/barrebre/goGetMTGPrices/collection"
)

func TestGetCardPrices(t *testing.T) {
	type inputs struct {
		Cards        []collection.Card
		priceChannel chan CardPrice
	}

	type testDefs struct {
		Name          string
		Inputs        inputs
		ExpectPass    bool
		ExpectedError string
	}

	tests := []testDefs{
		testDefs{
			Name: "Valid",
			Inputs: inputs{
				Cards:        collection.GetExampleCollection(),
				priceChannel: make(chan CardPrice),
			},
			ExpectPass: true,
		},
		testDefs{
			Name: "Invalid Card",
			Inputs: inputs{
				Cards:        collection.GetInvalidExampleCollection(),
				priceChannel: make(chan CardPrice),
			},
			ExpectPass: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			GetCardPrices(test.Inputs.Cards, test.Inputs.priceChannel)
		})
	}
}
