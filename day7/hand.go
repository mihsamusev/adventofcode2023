package main

type Card rune
type CardValues map[Card]int

var cardValues = CardValues{
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

var cardValuesWithJoker = CardValues{
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


func Beats(first, second Hand) bool {
	if Combo(first) > Combo(second) {
		return true
	}

	if Combo(first) < Combo(second) {
		return false
	}

	return BeatsByCard(first, second, cardValues)
}

func BeatsWithJoker(first, second Hand) bool {
	if ComboWithJoker(first) > ComboWithJoker(second) {
		return true
	}

	if ComboWithJoker(first) < ComboWithJoker(second) {
		return false
	}

	return BeatsByCard(first, second, cardValuesWithJoker)
}

func ComboWithJoker(h Hand) ComboValue {
	counts := make(map[Card]int)
	for _, c := range h.cards {
		counts[c]++
	}

	jokerCount, exist := counts['J']
	if !exist {
		return ComboFromCounts(counts)
	}

	if jokerCount > 3 {
		return Five
	}

	delete(counts, 'J')
	maxCard := Card('J')
	maxCount := 0
	for card, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCard = card
		}
	}
	counts[maxCard] += jokerCount

	return ComboFromCounts(counts)
}

func Combo(h Hand) ComboValue {
	counts := make(map[Card]int)
	for _, c := range h.cards {
		counts[c]++
	}
	return ComboFromCounts(counts)

}

func ComboFromCounts(counts map[Card]int) ComboValue {
	switch len(counts) {
	case 1:
		return Five
	case 2:
		for _, count := range counts {
			if count == 4 {
				return Four
			}
		}
		return FullHouse
	case 3:
		for _, count := range counts {
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

func BeatsByCard(first, second Hand, cardValues CardValues) bool {
	for i := range first.cards {
		if first.cards[i] != second.cards[i] {
			firstScore := cardValues[first.cards[i]]
			secondScore := cardValues[second.cards[i]]
			return firstScore > secondScore
		}
	}
	return false
}
