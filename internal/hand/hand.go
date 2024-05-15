package hand

import (
	"errors"
	"sort"

	"github.com/jqwez/pokergopher/internal/deck"
)

type Hand struct {
	Cards deck.Cards
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) AddCard(card *deck.Card) *Hand {
	h.Cards = append(h.Cards, card)
	return h
}

func (h *Hand) RemoveCard(card *deck.Card) (*Hand, error) {
	for i, c := range h.Cards {
		if c.IsSame(card) {
			h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
			return h, nil
		}
	}
	return h, errors.New("CARD NOT IN HAND")
}

func (h *Hand) HasCard(card *deck.Card) bool {
	for _, c := range h.Cards {
		if c.IsSame(card) {
			return true
		}
	}
	return false
}

func (h *Hand) FilterSuit(suit deck.CardSuit) *Hand {
	hand := NewHand()
	for _, card := range h.Cards {
		if card.Suit == suit {
			hand.AddCard(card)
		}
	}
	return hand
}

func (h *Hand) FilterRank(rank deck.CardRank) *Hand {
	hand := NewHand()
	for _, card := range h.Cards {
		if card.Rank == rank {
			hand.AddCard(card)
		}
	}
	return hand
}

func (h *Hand) SortByRank() *Hand {
	sort.Sort(h.Cards)
	return h
}
