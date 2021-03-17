package bankmachine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DigitTestCase struct {
	Input          string
	ExpectedOutput string
}

var TestNumberCases = []DigitTestCase{
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

func TestParseNumbersParsesAllDigits(t *testing.T) {
	m := NewBankMachine()

	for _, tc := range TestNumberCases {
		ds := m.Parse(tc.Input)
		assert.Equal(t, tc.ExpectedOutput, ds.String())
	}
}

var ChecksumTestCases = []DigitTestCase{
	{`
 _  _     _  _        _  _ 
|_ |_ |_| _|  |  ||_||_||_ 
|_||_|  | _|  |  |  | _| _|`,
		"664371495 ERR",
	},
	{`
 _  _  _  _  _  _  _  _    
| || || || || || || ||_   |
|_||_||_||_||_||_||_| _|  |`,
		"000000051",
	},
	{`
    _  _  _  _  _  _     _ 
|_||_|| || ||_   |  |  | _ 
  | _||_||_||_|  |  |  | _|`,
		"49006771? ILL",
	},
	{`
    _  _     _  _  _  _  _ 
  | _| _||_| _ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _ `,
		"1234?678? ILL",
	},
}

func TestNumbersChecksumPasses(t *testing.T) {
	m := NewBankMachine()

	for _, ctc := range ChecksumTestCases {
		cs := m.Parse(ctc.Input)
		assert.Equal(t, ctc.ExpectedOutput, cs.Checksum())
	}
}
