package deck

import (
	"cards/card"
	"cards/rank"
	"cards/suit"
	"math/rand"
)

// Repräsentiert einen Kartenstapel.
type Deck struct {
	cards []card.Card
}

// Empty erzeugt neues leeres Deck
func Empty() *Deck {
	return &Deck{}
}

// New32 erzeugt ein neues Skatblatt mit 32 Karten.
func New32() *Deck {
	d := Empty()
	for i := suit.Clubs; i < suit.Diamonds; i++ {
		for j := rank.Seven; j < rank.Ace; j++ {
			d.Add(card.Card{Rank: j, Suit: i})
		}
	}
	return d
}

// New52 erzeugt ein neues französisches Blatt mit 52 Karten.
func New52() *Deck {
	d := Empty()
	for i := suit.Clubs; i < suit.Diamonds; i++ {
		for j := rank.Two; j < rank.Ace; j++ {
			d.Add(card.Card{Rank: j, Suit: i})
		}
	}
	return d
}

// Add fügt eine Karte zum Deck hinzu.
func (d *Deck) Add(c card.Card) {
	d.cards = append(d.cards, c)
}

// Shuffle mischt das Deck.
func (d *Deck) Shuffle() {
	len := len(d.cards)

	for i := 0; i < len; i++ {
		r1 := rand.Intn(len) - 1
		r2 := rand.Intn(len) - 1

		d.cards[r1], d.cards[r2] = d.cards[r2], d.cards[r1]
	}

}

// Draw zieht die oberste Karte.
// Entfernt die Karte aus dem Deck und liefert sie zurück.
func (d *Deck) Draw() card.Card {
	drawn := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]

	return drawn
}

// Top liefert die oberste Karte.
// Entfernt sie aber nicht.
func (d Deck) Top() card.Card {
	return d.cards[len(d.cards)-1]
}
