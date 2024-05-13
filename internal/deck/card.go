package deck

import (
	"fmt"
)

type CardRank int
type CardSuit int

type Card struct {
	Suit CardSuit
	Rank CardRank
}

func NewCard(suit CardSuit, rank CardRank) (*Card, error) {
	if suit < CLUB || suit > SPADE || rank < ACE || rank > JOKER2 {
		return nil, fmt.Errorf("INVALID CARD : %s %s", suit, rank)
	}
	return &Card{
		Suit: suit,
		Rank: rank,
	}, nil
}

func (c *Card) String() string {
	return fmt.Sprintf("{Card %s : %s}", c.Suit, c.Rank)
}

func (c *Card) IsSame(other *Card) bool {
	return c.Rank == other.Rank && c.Suit == other.Suit
}

type Cards []*Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	a := c[i].Rank
	if a == ACE {
		a = 50
	}
	b := c[j].Rank
	if b == ACE {
		b = 50
	}
	if a == b {
		return c[i].Suit > c[j].Suit
	}
	return a > b
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

const (
	ACE CardRank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	JOKER1
	JOKER2
)

var CardRanks map[CardRank]string = func() map[CardRank]string {
	ranks := []CardRank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, JOKER1, JOKER2}
	rankStrings := []string{"ACE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE", "TEN", "JACK", "QUEEN", "KING", "JOKER1", "JOKER2"}
	rankMap := make(map[CardRank]string)
	for i, rank := range ranks {
		rankMap[rank] = rankStrings[i]
	}
	return rankMap
}()

func (cr CardRank) String() string {
	return CardRanks[cr]
}

const (
	CLUB CardSuit = iota
	DIAMOND
	HEART
	SPADE
	JOKERSUIT
)

var CardSuits map[CardSuit]string = func() map[CardSuit]string {
	suits := []CardSuit{CLUB, DIAMOND, HEART, SPADE}
	suitStrings := []string{"CLUB", "DIAMOND", "HEART", "SPADE"}
	suitMap := make(map[CardSuit]string)
	for i, suit := range suits {
		suitMap[suit] = suitStrings[i]
	}
	return suitMap
}()

func (cs CardSuit) String() string {
	return CardSuits[cs]
}
