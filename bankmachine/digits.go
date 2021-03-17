package bankmachine

import (
	"fmt"
	"strings"
)

type Digits [9]*Digit

type Digit struct {
	Raw string
}

func (d Digit) String() string {
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

	return "?"
}

func (ds Digits) String() string {
	var out string

	for _, d := range ds {
		out += d.String()
	}

	return out
}

func (ds *Digits) Checksum() string {

	number := ds.String()

	if strings.Contains(number, "?") {
		return fmt.Sprintf("%s ILL", number)
	}

	var ch int
	for i := len(number) - 1; i >= 0; i-- {
		ch += (9 * (i + 1))
	}

	if ch%11 != 0 {
		return fmt.Sprintf("%s ERR", number)
	}

	return number
}
