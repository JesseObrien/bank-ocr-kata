package bankmachine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NumberTestCase struct {
	Input          string
	ExpectedOutput string
}

var TestNumberCases = []NumberTestCase{
	{`
 _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|`,
		"000000000",
	},
	{`
                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |`,
		"111111111",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ `,
		"222222222",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
 _| _| _| _| _| _| _| _| _|`,
		"333333333",
	},
	{`
                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |`,
		"444444444",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
 _| _| _| _| _| _| _| _| _|`,
		"555555555",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
|_||_||_||_||_||_||_||_||_|`,
		"666666666",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |`,
		"777777777",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
|_||_||_||_||_||_||_||_||_|`,
		"888888888",
	},
	{`
 _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
 _| _| _| _| _| _| _| _| _|`,
		"999999999",
	},
	{`
    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|`,
		"123456789",
	},
}

func TestParseNumbers(t *testing.T) {
	m := NewBankMachine()

	for _, tc := range TestNumberCases {
		assert.Equal(t, tc.ExpectedOutput, m.read(tc.Input))
	}

}
