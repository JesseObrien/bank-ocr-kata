package bankmachine

import (
	"regexp"
)

// BankMachine go brrrrrr
type BankMachine struct {

	// @TODO we'll need something to parse digits

}

type Digits [9]*Digit

type Digit struct {
	Raw    string
	Parsed string
}

func (d *Digit) what() string {
	if d.Raw == Zero {
		return "0"
	}

	if d.Raw == One {
		return "1"
	}

	if d.Raw == Two {
		return "2"
	}

	if d.Raw == Three {
		return "3"
	}

	if d.Raw == Four {
		return "4"
	}

	if d.Raw == Five {
		return "5"
	}
	if d.Raw == Six {
		return "6"
	}

	if d.Raw == Seven {
		return "7"
	}

	if d.Raw == Eight {
		return "8"
	}

	if d.Raw == Nine {
		return "9"
	}

	return "ILL"
}

func (ds *Digits) out() string {

	var out string

	for _, d := range ds {
		out += d.what()
	}

	return out
}

func NewBankMachine() *BankMachine {
	return &BankMachine{}
}

func (bm *BankMachine) read(input string) string {
	ds := parseDigits(input)

	return ds.out()
}

func parseDigits(input string) Digits {

	var digits Digits

	// split the string into chunks of 9
	r := regexp.MustCompile(`[ _\|]{27}`)
	rawLines := r.FindAllStringSubmatch(input, -1)

	for _, line := range rawLines {
		l := regexp.MustCompile(`[ _\|]{3}`)
		lineChunks := l.FindAllStringSubmatch(line[0], -1)

		for pos, str := range lineChunks {
			if digits[pos] == nil {
				digits[pos] = &Digit{}
			}

			digits[pos].Raw += string(str[0])
		}
	}

	return digits
}
