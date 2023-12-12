package main


type Card rune

var cardValues = map[Card]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardValuesWithJoker = map[Card]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}
type ComboValue int

const (
	OneCard ComboValue = iota
	OnePair
	TwoPair
	Triplet
	FullHouse
	Four
	Five
)

type Hand struct {
	cards []Card
	bid   int
}

func Beats(h, other Hand) bool {
	if Combo(h) > Combo(other) {
		return true
	}

	if Combo(h) < Combo(other) {
		return false
	}
	
	return BeatsByCard(h, other)
}


func Combo(h Hand) ComboValue {
	uniqueCards := make(map[Card]int)
	for _, c := range h.cards {
		uniqueCards[c]++
	}

	switch len(uniqueCards) {
	case 1:
		return Five
	case 2:
		for _, count := range uniqueCards {
			if count == 4 {
				return Four
			}
		}
		return FullHouse
	case 3:
		for _, count := range uniqueCards {
			if count == 3 {
				return Triplet
			}
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return OneCard
	}
}

func BeatsByCard(h, other Hand) bool {

	for i := range h.cards {
		if h.cards[i] != other.cards[i] {
			return h.cards[i].Beats(other.cards[i])
		}

	}
	return false
}

func (c *Card) Beats(other Card) bool {
	return cardValues[*c] > cardValues[other]
}
