//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

//Suit Type
type Suit uint8

//Suit Constants
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

//Rank Type
type Rank uint8

//Rank Constants
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//Card Struct
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
