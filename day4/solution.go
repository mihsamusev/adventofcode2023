package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SolvePartOne(dataFile string, maxScans int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	i := 0
	total := 0
	for {
		if maxScans != -1 && i == maxScans {
			break
		}
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Printf("%d: %s\n", i+1, line)
		card, err := ParseCard(line)
		if err != nil {
			continue
		}
		matches := OwnedWinningNumbers(card)
		points := GetPoints(matches)
		cardIds := NextCardIds(card)
		fmt.Printf("    %v -> %d\n", matches, points)
		fmt.Printf("    %v\n", cardIds)
		total += points

		fmt.Println()
		i++

	}
	fmt.Printf("Total: %d\n", total)
}

type CardsLookup map[int][]int

func CountCardsRecursive(startId int, cardsMap CardsLookup) int {
	nextCards := cardsMap[startId]
	sum := len(nextCards)
	
	fmt.Printf("%d: %d ->  %v\n", startId, sum, nextCards)

	for _, nextId := range nextCards {
		sum += CountCardsRecursive(nextId, cardsMap)
	}

	return sum

}

func SolvePartTwo(dataFile string) {
	content, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	lines := strings.Split(string(content), "\n")
	cards := make([]Card, 0)
	cardsMap := make(CardsLookup)
	for _, line := range lines {
		card, err := ParseCard(line)
		if err != nil {
			continue
		}
		cards = append(cards, card)
		cardsMap[card.id] = NextCardIds(card)
	}

	//totalCards := CountTotalRecursive(cards, cardsMap)
	totalCards := CountTotalIterative(cards, cardsMap)
	fmt.Println(cardsMap)
	fmt.Println(totalCards)
}

func CountTotalRecursive(cards[]Card, cardsMap CardsLookup) int {
	totalCards := len(cards)
	for _, card := range cards {
		totalCards += CountCardsRecursive(card.id, cardsMap)
		fmt.Println()
	}
	return totalCards
}


func CountTotalIterative(cards[]Card, cardsMap CardsLookup) int {
	stack := make([]int, 0, len(cards))
	for _, card := range cards {
		stack = append(stack, card.id)
	}
	
	totalCards := 0
	for len(stack) != 0 {
		last := len(stack) - 1
		currentId := stack[last]
		stack = stack[:last]

		nextCards := cardsMap[currentId]
		stack = append(stack, nextCards...)
		totalCards ++

	}
	return totalCards
}