package cmd

import (
	"fmt"

	"github.com/jqwez/pokergopher/internal/deck"
)

func Main() {
	d := deck.NewDeck(deck.Standard)
	for i := 0; i < 54; i++ {
		card, err := d.DrawCard()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(card)
	}
}
