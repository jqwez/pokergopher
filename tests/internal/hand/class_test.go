package hand_test

import (
	"testing"

	"github.com/jqwez/pokergopher/internal/deck"
	"github.com/jqwez/pokergopher/internal/hand"
)

func TestNewClasser(t *testing.T) {

}

func TestIsFlush(t *testing.T) {
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
			{Suit: deck.SPADE, Rank: deck.KING},
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
	results := []bool{false, false, false, true, true, true}
	classer := hand.NewClasser()
	for i, cards := range hands {
		_h := hand.NewHand()
		_h.Cards = cards
		res, h := classer.IsFlush(_h)
		if res != results[i] {
			t.Fatal("Hand should be flush!", h.Cards)
		}
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

func TestFindBestHand(t *testing.T) {

	hands := []struct {
		Cards deck.Cards
		Class hand.HandClass
	}{
		{
			Cards: deck.Cards{
				{Suit: deck.SPADE, Rank: deck.ACE},
				{Suit: deck.SPADE, Rank: deck.FIVE},
				{Suit: deck.SPADE, Rank: deck.SEVEN},
				{Suit: deck.SPADE, Rank: deck.KING},
				{Suit: deck.CLUB, Rank: deck.TEN},
			},
			Class: hand.NOPAIR,
		},
		{
			Cards: deck.Cards{
				{Suit: deck.SPADE, Rank: deck.TEN},
				{Suit: deck.SPADE, Rank: deck.JACK},
				{Suit: deck.SPADE, Rank: deck.QUEEN},
				{Suit: deck.SPADE, Rank: deck.KING},
				{Suit: deck.SPADE, Rank: deck.ACE},
			},
			Class: hand.ROYAL,
		},
		{
			Cards: deck.Cards{
				{Suit: deck.SPADE, Rank: deck.ACE},
				{Suit: deck.SPADE, Rank: deck.TWO},
				{Suit: deck.SPADE, Rank: deck.THREE},
				{Suit: deck.SPADE, Rank: deck.FOUR},
				{Suit: deck.CLUB, Rank: deck.FIVE},
			},
			Class: hand.STRAIGHT,
		},
		{
			Cards: deck.Cards{
				{Suit: deck.SPADE, Rank: deck.ACE},
				{Suit: deck.HEART, Rank: deck.TWO},
				{Suit: deck.CLUB, Rank: deck.THREE},
				{Suit: deck.SPADE, Rank: deck.TWO},
				{Suit: deck.CLUB, Rank: deck.THREE},
				{Suit: deck.HEART, Rank: deck.FOUR},
				{Suit: deck.CLUB, Rank: deck.FIVE},
			},
			Class: hand.STRAIGHT,
		},
		{
			Cards: deck.Cards{
				{Suit: deck.SPADE, Rank: deck.TEN},
				{Suit: deck.SPADE, Rank: deck.JACK},
				{Suit: deck.SPADE, Rank: deck.QUEEN},
				{Suit: deck.SPADE, Rank: deck.KING},
				{Suit: deck.SPADE, Rank: deck.THREE},
				{Suit: deck.HEART, Rank: deck.NINE},
			},
			Class: hand.FLUSH,
		},
		{
			Cards: deck.Cards{
				{Suit: deck.SPADE, Rank: deck.EIGHT},
				{Suit: deck.SPADE, Rank: deck.JACK},
				{Suit: deck.SPADE, Rank: deck.QUEEN},
				{Suit: deck.SPADE, Rank: deck.KING},
				{Suit: deck.SPADE, Rank: deck.THREE},
				{Suit: deck.HEART, Rank: deck.NINE},
			},
			Class: hand.FLUSH,
		},
	}
	for i, testCase := range hands {
		h := hand.NewHand()
		h.Cards = testCase.Cards
		classer := hand.NewClasser().SetHand(h)
		if testCase.Class != classer.Class {
			t.Fatal(i, " : Evaluated to ", classer.Class, " but should be ", testCase.Class)
		}
	}
}
