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

func CountCards(startId int, cardsMap map[int][]int) int {
	nextCards := cardsMap[startId]
	sum := len(nextCards)
	
	fmt.Printf("%d: %d ->  %v\n", startId, sum, nextCards)

	for _, nextId := range nextCards {
		sum += CountCards(nextId, cardsMap)
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
	cardsMap := make(map[int][]int)
	for _, line := range lines {
		card, err := ParseCard(line)
		if err != nil {
			continue
		}
		cards = append(cards, card)
		cardsMap[card.id] = NextCardIds(card)
	}

	fmt.Println(cardsMap)
	totalCards := len(cards)
	for _, card := range cards {
		totalCards += CountCards(card.id, cardsMap)
		fmt.Println()
	}
	fmt.Println(totalCards)
}
