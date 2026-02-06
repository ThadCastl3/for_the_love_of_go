package creditcard

import "errors"

type card struct {
	number string
}

func NewCard(number string) (card, error) {
	if len(number) != 16 {
		return card{}, errors.New("invalid card number, must be 16 digits")
	}
	return card{number}, nil
}

func (c *card) Number() string {
	return c.number
}
