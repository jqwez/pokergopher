package hand

import "github.com/jqwez/pokergopher/internal/deck"

type Classer struct{}

func NewClasser() *Classer {
	return &Classer{}
}

func (c *Classer) IsFlush(_hand *Hand) (bool, *Hand) {
	suitCounter := make(map[deck.CardSuit]int)
	for _, card := range _hand.Cards {
		suitCounter[card.Suit] = suitCounter[card.Suit] + 1
		if suitCounter[card.Suit] >= 5 {
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
