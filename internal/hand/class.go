package hand

import (
	"fmt"
	"github.com/jqwez/pokergopher/internal/deck"
)

type HandClass int

const (
	ROYAL HandClass = iota
	SFLUSH
	QUADS
	FHOUSE
	FLUSH
	STRAIGHT
	TRIPS
	TWOPAIR
	PAIR
	NOPAIR
)

type Classer struct {
	Hand  *Hand
	Class HandClass
}

func NewClasser() *Classer {
	return &Classer{}
}

func NewClasserH(h *Hand) *Classer {
	c := NewClasser()
	return c.SetHand(h)
}

func (c *Classer) SetHand(h *Hand) *Classer {
	c.Hand = h
	c.findBestHand()
	return c
}

func (c *Classer) findBestHand() *Classer {
	pairsType, pairs := c.PairsEvaluator(c.Hand)
	isStraightFlush, straightFlush := c.IsStraightFlush(c.Hand)
	if isStraightFlush {
		c.Hand = straightFlush
		c.Class = SFLUSH
		if c.Hand.Cards[4].Rank == deck.TEN {
			c.Class = ROYAL
		}
		return c
	}
	if pairsType == QUADS {
		c.Hand = pairs
		c.Class = QUADS
		return c
	}
	isFlush, flush := c.IsFlush(c.Hand)
	if isFlush {
		c.Hand = flush
		c.Class = FLUSH
		return c
	}
	isStraight, straight := c.IsStraight(c.Hand)
	if isStraight {
		c.Hand = straight
		c.Class = STRAIGHT
		return c
	}

	c.Hand.Cards = c.Hand.Cards[:5]
	c.Class = NOPAIR
	return c
}

func (c *Classer) PairsEvaluator(h *Hand) (HandClass, *Hand) {
	if len(h.Cards) < 5 {
		return NOPAIR, h.SortByRank()
	}
	counter := make(map[deck.CardRank]int)
	for _, card := range h.Cards {
		counter[card.Rank] += 1
	}
	quantities := make(map[int][]deck.CardRank)
	for rank, quantity := range counter {
		quantities[quantity] = append(quantities[quantity], rank)
	}
	for _, v := range []int{4, 3, 2, 1} {
		ranks := quantities[v]
		if len(ranks) < 1 {
			continue
		}
		tRank := deck.TWO
		for _, rank := range ranks {
			if rank == deck.ACE {
				tRank = rank
			} else if rank > tRank {
				tRank = rank
			}
		}
		_h := c.FillByHighCard(h, tRank, 1)
		return QUADS, _h
	}
	return NOPAIR, h
}

func (c *Classer) FillByHighCard(h *Hand, exclude deck.CardRank, addCount int) *Hand {
	h = h.SortByRank()
	_h := h.FilterRank(exclude)
	for _, card := range h.Cards {
		if card.Rank != exclude && addCount > 0 {
			_h.AddCard(card)
			addCount -= 1
		}
	}
	return _h
}

func (c *Classer) IsStraightFlush(h *Hand) (bool, *Hand) {
	res, straightFlush := c.IsStraight(h.FilterSuit(h.Cards[0].Suit))
	if res {
		return true, straightFlush
	}
	return false, h
}

func (c *Classer) IsFlush(_hand *Hand) (bool, *Hand) {
	suitCounter := make(map[deck.CardSuit]int)
	for _, card := range _hand.Cards {
		suitCounter[card.Suit] = suitCounter[card.Suit] + 1
		if suitCounter[card.Suit] >= 5 {
			_hand = _hand.FilterSuit(card.Suit).SortByRank()
			_hand.Cards = _hand.Cards[:5]
			return true, _hand
		}
	}
	return false, _hand
}

func (c *Classer) IsStraight(_hand *Hand) (bool, *Hand) {
	cards := _hand.SortByRank().Cards
	l := len(cards)
	inARow := 1
	for i := 1; i < l; i++ {
		prev := cards[i-1]
		card := cards[i]
		if prev.Rank-1 == card.Rank || (prev.Rank == deck.ACE && card.Rank == deck.KING) {
			inARow += 1
		} else if prev.Rank == card.Rank {
			continue
		} else {
			inARow = 1
		}
		if inARow == 5 {
			return true, _hand
		}
		if card.Rank == deck.TWO && inARow == 4 && cards[0].Rank == deck.ACE {
			return true, _hand
		}
	}
	return false, _hand
}

func (hc HandClass) String() string {
	s := []string{"Royal Flush", "Straight Flush", "Four of a Kind", "Full House", "Flush", "Straight", "Three of a Kind", "Two Pairs", "Pair", "No Pair"}
	return fmt.Sprintf("%s", s[hc])
}
