package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Three, Suit: Club})
	fmt.Println(Card{Rank: Four, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Three of Clubs
	// Four of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{
		Suit: Spade,
		Rank: Ace,
	}
	if cards[0] != exp {
		t.Error("Expected ace of spades as first card. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{
		Suit: Spade,
		Rank: Ace,
	}
	if cards[0] != exp {
		t.Error("Expected ace of spades as first card. Received: ", cards[0])
	}
}

func TestJoker(t *testing.T){
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3{
		t.Error("Expected 3 Jokers. Received: ", count)
	}
}

func TestFilter(t *testing.T){
	filter := func(card Card) bool {
		return card.Rank == 2 || card.Rank == 3
	}

	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == 2 || c.Rank == 3 {
			t.Error("Did not expect Two's or Three's")
		}
	}
}

func TestDeck(t *testing.T){
	cards := New(Deck(3))
	if len(cards) != 3 * 52 {
		t.Error("Expected 156 cards. Found: ", len(cards))
	}
}
