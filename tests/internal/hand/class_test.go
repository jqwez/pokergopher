package hand_test

import (
	"testing"

	"github.com/jqwez/pokergopher/internal/deck"
	"github.com/jqwez/pokergopher/internal/hand"
	"github.com/jqwez/pokergopher/tests/internal/util"
)

func TestNewClasser(t *testing.T) {

}

func TestIsFlush(t *testing.T) {
	h := hand.NewHand()
	card1, err := deck.NewCard(deck.SPADE, deck.TEN)
	util.CheckFatal(err)
	card2, err := deck.NewCard(deck.SPADE, deck.TEN)
	util.CheckFatal(err)
	card3, err := deck.NewCard(deck.SPADE, deck.TEN)
	util.CheckFatal(err)
	card4, err := deck.NewCard(deck.SPADE, deck.TEN)
	util.CheckFatal(err)
	card5, err := deck.NewCard(deck.SPADE, deck.TEN)
	util.CheckFatal(err)
	cards := []*deck.Card{card1, card2, card3, card4, card5}
	for _, card := range cards {
		h.AddCard(card)
	}
	classer := hand.NewClasser()
	if res, _ := classer.IsFlush(h); !res {
		t.Fatal("Hand should be flush!", h.Cards)
	}
}

func TestIsStraight(t *testing.T) {
	hands := []deck.Cards{
		{
			{Suit: deck.SPADE, Rank: deck.ACE},
			{Suit: deck.SPADE, Rank: deck.FIVE},
			{Suit: deck.SPADE, Rank: deck.SEVEN},
			{Suit: deck.SPADE, Rank: deck.KING},
			{Suit: deck.CLUB, Rank: deck.TEN},
		},
		{
			{Suit: deck.SPADE, Rank: deck.TEN},
			{Suit: deck.SPADE, Rank: deck.JACK},
			{Suit: deck.SPADE, Rank: deck.QUEEN},
			{Suit: deck.SPADE, Rank: deck.KING},
			{Suit: deck.HEART, Rank: deck.ACE},
		},
		{
			{Suit: deck.SPADE, Rank: deck.ACE},
			{Suit: deck.SPADE, Rank: deck.TWO},
			{Suit: deck.SPADE, Rank: deck.THREE},
			{Suit: deck.SPADE, Rank: deck.FOUR},
			{Suit: deck.CLUB, Rank: deck.FIVE},
		},
		{
			{Suit: deck.SPADE, Rank: deck.ACE},
			{Suit: deck.HEART, Rank: deck.TWO},
			{Suit: deck.CLUB, Rank: deck.THREE},
			{Suit: deck.SPADE, Rank: deck.TWO},
			{Suit: deck.SPADE, Rank: deck.THREE},
			{Suit: deck.SPADE, Rank: deck.FOUR},
			{Suit: deck.CLUB, Rank: deck.FIVE},
		},
		{
			{Suit: deck.SPADE, Rank: deck.TEN},
			{Suit: deck.SPADE, Rank: deck.JACK},
			{Suit: deck.SPADE, Rank: deck.QUEEN},
			{Suit: deck.SPADE, Rank: deck.KING},
			{Suit: deck.SPADE, Rank: deck.THREE},
			{Suit: deck.HEART, Rank: deck.NINE},
		},
		{
			{Suit: deck.SPADE, Rank: deck.EIGHT},
			{Suit: deck.SPADE, Rank: deck.JACK},
			{Suit: deck.SPADE, Rank: deck.QUEEN},
			{Suit: deck.SPADE, Rank: deck.KING},
			{Suit: deck.SPADE, Rank: deck.THREE},
			{Suit: deck.HEART, Rank: deck.NINE},
		},
	}
	results := []bool{false, true, true, true, true, false}
	classer := hand.NewClasser()
	for i, h := range hands {
		_hand := hand.NewHand()
		_hand.Cards = h
		res, _ := classer.IsStraight(_hand)
		if res != results[i] {
			t.Fatal("Hand should not be", res, "for straight", h)
		}
	}
}
