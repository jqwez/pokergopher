package deck_test

import (
	"testing"

	"github.com/jqwez/pokergopher/internal/deck"
)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck(deck.Standard)
	if len(d.Cards) != 52 {
		t.Errorf("Standard Deck is not 52 cards but showing %d", len(d.Cards))
	}
	for _, card := range d.Cards {
		if card.Suit == deck.SPADE && card.Rank == deck.JACK {
			return
		}
	}
	t.Errorf("Missing Card")
}

func TestDrawCard(t *testing.T) {
	r := deck.NewDeck(deck.Standard)
	_, err := r.DrawCard()
	if err != nil {
		t.Errorf("Failed to draw Card: %s", err)
	}
}

func TestDrawDeck(t *testing.T) {
	d := deck.NewDeck(deck.Standard)
	for i := 0; i < 52; i++ {
		_, err := d.DrawCard()
		if err != nil {
			t.Errorf("Error drawing card %s", err)
		}
	}
	if remaining := len(d.Cards); remaining != 0 {
		t.Errorf("Cards remaining: %d but expected 0", remaining)
	}
	_, err := d.DrawCard()
	if err == nil {
		t.Errorf("No error drew when drawing too many cards")
	}
}
