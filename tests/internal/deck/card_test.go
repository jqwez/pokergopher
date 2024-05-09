package deck_test

import (
	"testing"

	"github.com/jqwez/pokergopher/internal/deck"
)

func TestNewCard(t *testing.T) {
	card, err := deck.NewCard(deck.SPADE, deck.TEN)
	if err != nil {
		t.Error(err)
	}
	if card.Suit != deck.SPADE {
		t.Errorf("Incorrect suit")
	}
	if card.Rank != deck.TEN {
		t.Errorf("Incorrect rank")
	}
	_, err = deck.NewCard(10, -1)
	if err == nil {
		t.Errorf("Out of bounds card should throw")
	}
}

func TestGetRankString(t *testing.T) {
	card, err := deck.NewCard(deck.SPADE, deck.TEN)
	if err != nil {
		t.Error(err)
	}
	if card.Rank.String() != "TEN" {
		t.Errorf("Card Rank stringified should == TEN")
	}
}
