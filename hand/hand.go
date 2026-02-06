package hand

import (
	"cards/card"
	"cards/deck"
)

// TODO: Hier Datentyp `Hand` definieren
// und Methoden hinzufügen.
// Tests in separater Datei nicht vergessen.

type Hand struct {
	cards []card.Card
}

func Empty() *Hand {
	return &Hand{}
}

func Start5(d *deck.Deck) *Hand {
	h := Empty()
	h.cards = append(h.cards, d.Draw())
	return h
}
