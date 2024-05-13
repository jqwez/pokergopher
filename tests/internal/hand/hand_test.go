package hand_test

import (
	"testing"

	"github.com/jqwez/pokergopher/internal/deck"
	"github.com/jqwez/pokergopher/internal/hand"
	"github.com/jqwez/pokergopher/tests/internal/util"
)

func TestNewHand(t *testing.T) {
	_ = hand.NewHand()
}

func TestHasCard(t *testing.T) {
	hand := hand.NewHand()
	carda, err := deck.NewCard(deck.HEART, deck.TEN)
	util.CheckFatal(err)
	cardb, _ := deck.NewCard(deck.SPADE, deck.TEN)
	util.CheckFatal(err)
	if hand.HasCard(carda) {
		t.Fatal(carda, "Should not be in hand")
	}
	hand.Cards = []*deck.Card{carda, cardb}
	cardc, err := deck.NewCard(deck.CLUB, deck.TEN)
	util.CheckFatal(err)
	if hand.HasCard(cardc) {
		t.Fatal(cardc, "Should not be in hand")
	}
	if !hand.HasCard(carda) {
		t.Fatal(carda, "Should be in hand")
	}
}

func TestAddCard(t *testing.T) {
	hand := hand.NewHand()
	carda, err := deck.NewCard(deck.HEART, deck.ACE)
	util.CheckFatal(err)
	hand = hand.AddCard(carda)
	util.CheckFatal(err)
	cardb, err := deck.NewCard(deck.HEART, deck.TEN)
	util.CheckFatal(err)
	if !hand.HasCard(carda) || hand.HasCard(cardb) {
		t.Fatal("Should contain only HA but contains :", hand.Cards)
	}
}

func TestRemoveCard(t *testing.T) {
	hand := hand.NewHand()
	carda, err := deck.NewCard(deck.HEART, deck.ACE)
	hand.AddCard(carda)
	util.CheckFatal(err)
	cardb, err := deck.NewCard(deck.HEART, deck.TEN)
	util.CheckFatal(err)
	hand.AddCard(cardb)
	cardc, err := deck.NewCard(deck.HEART, deck.KING)
	util.CheckFatal(err)
	hand.AddCard(cardc)
	if !hand.HasCard(cardc) {
		t.Fatal(cardc, "Should be in the hand")
	}
	hand.RemoveCard(cardc)
	if hand.HasCard(cardc) {
		t.Fatal(cardc, "Should NOT be in the hand")
	}
}

func TestFilterSuit(t *testing.T) {
	hand := buildTestHand()
	hand = hand.FilterSuit(deck.HEART)
	if hand.HasCard(&deck.Card{Suit: deck.SPADE, Rank: deck.KING}) {
		t.Fatal(&deck.Card{Suit: deck.SPADE, Rank: deck.KING}, "Should NOT be in the hand")
	}
	if !hand.HasCard(&deck.Card{Suit: deck.HEART, Rank: deck.TEN}) {
		t.Fatal(&deck.Card{Suit: deck.HEART, Rank: deck.TEN}, "Should be in the hand")
	}
}

func TestSortByRank(t *testing.T) {
	h := buildTestHand()
	h.SortByRank()
	if h.Cards[4].Rank != deck.THREE {
		t.Fatal(h.Cards[4], "Card should be Three")
	}
	if h.Cards[3].Rank != deck.TEN {
		t.Fatal(h.Cards[3], "Card should be TEN")
	}
}

func buildTestHand() *hand.Hand {
	h := hand.NewHand()
	carda, err := deck.NewCard(deck.HEART, deck.ACE)
	h.AddCard(carda)
	util.CheckFatal(err)
	cardb, err := deck.NewCard(deck.HEART, deck.TEN)
	util.CheckFatal(err)
	h.AddCard(cardb)
	cardc, err := deck.NewCard(deck.HEART, deck.KING)
	util.CheckFatal(err)
	h.AddCard(cardc)
	card, err := deck.NewCard(deck.SPADE, deck.KING)
	util.CheckFatal(err)
	h.AddCard(card)
	carde, err := deck.NewCard(deck.CLUB, deck.THREE)
	util.CheckFatal(err)
	h.AddCard(carde)
	return h
}
