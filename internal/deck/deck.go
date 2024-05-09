package deck

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"log"
)

type DeckType int

const (
	Standard DeckType = iota
	TwoJokers
)

type Deck struct {
	Cards []*Card
}

func NewDeck(dt DeckType) *Deck {
	var cards = buildDeck(dt)
	return &Deck{
		Cards: cards,
	}
}

func (d *Deck) DrawCard() (*Card, error) {
	if len(d.Cards) < 1 {
		return nil, errors.New("NO MORE CARDS TO DRAW")
	}
	var randomNumber uint32
	err := binary.Read(rand.Reader, binary.BigEndian, &randomNumber)
	if err != nil {
		return nil, err
	}
	idx := int(randomNumber) % len(d.Cards)
	card := d.Cards[idx]
	d.Cards = append(d.Cards[:idx], d.Cards[idx+1:]...)
	return card, nil
}

var StandardDeck = func() []*Card {
	var cards []*Card
	for suit := range CardSuits {
		for rank := range CardRanks {
			if rank > KING {
				continue
			}
			card, err := NewCard(suit, rank)
			if err != nil {
				log.Fatal(err)
				return nil
			}
			cards = append(cards, card)
		}
	}
	return cards
}()

var OneJoker = func() []*Card {
	cards := StandardDeck
	card, err := NewCard(JOKERSUIT, JOKER1)
	if err != nil {
		log.Fatal(err)
	}
	cards = append(cards, card)
	return cards
}

func buildDeck(dt DeckType) []*Card {
	if dt == Standard {
		return StandardDeck
	}
	return StandardDeck
}
