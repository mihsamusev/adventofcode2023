package main

import (
	"common"
	"fmt"
	"os"
	"sort"
)

func SolvePartOne(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	hands, err := ParseHands(string(content))

	sort.Slice(hands, func(i, j int) bool { return !Beats(hands[i], hands[j]) })

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %v\n", err)
	}

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
		fmt.Printf("Rank %d: cards: %s, bid: %d\n", i + 1, string(h.cards), h.bid)
	}
	fmt.Printf("Total: %d\n", total)
}

func SolvePartTwo(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	hands, err := ParseHands(string(content))

	sort.Slice(hands, func(i, j int) bool { return !BeatsWithJoker(hands[i], hands[j]) })

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %v\n", err)
	}

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
		fmt.Printf("Rank %d: cards: %s, bid: %d\n", i + 1, string(h.cards), h.bid)
	}
	fmt.Printf("Total: %d\n", total)
}
