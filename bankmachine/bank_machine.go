package bankmachine

import (
	"regexp"
)

// BankMachine go brrrrrr
type BankMachine struct{}

func NewBankMachine() *BankMachine {
	return &BankMachine{}
}

func (bm *BankMachine) Parse(input string) Digits {
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
