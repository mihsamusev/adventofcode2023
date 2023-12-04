package main

import "math"


type Card struct {
    id int
    winning []int
    owned []int
}

func OwnedWinningNumbers(card Card) []int {
	result := make([]int, 0)
	for _, w := range card.winning {
		for _, o := range card.owned {
			if o == w {
				result = append(result, o)
				break
			}
		}
	}
	return result
}

func GetPoints(values []int) int {
	n := len(values)
	if n == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(n - 1)))
}

func NextCardIds(card Card) []int {
	matches := OwnedWinningNumbers(card)
	n := len(matches)
	if n == 0 {
		return make([]int, 0)
	}
	ids := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		ids = append(ids, card.id + i)	
	}
	return ids
}