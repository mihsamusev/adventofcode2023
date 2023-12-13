package main

import (
	"common"
	"fmt"
	"os"
)

func SolvePartOne(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	numbersGroup, err := common.ParseSlices(string(content), "\n")
	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	total := 0
	for _, numbers := range numbersGroup {
		next := NextReading(numbers)
		total += next
		fmt.Printf("Numbers: %v prev %d\n", numbers, next)
	}

	fmt.Printf("Total: %v\n", total)
}

func SolvePartTwo(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	numbersGroup, err := common.ParseSlices(string(content), "\n")
	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	total := 0
	for _, numbers := range numbersGroup {
		prev := PrevReading(numbers)
		total += prev
		fmt.Printf("Numbers: %v prev %d\n", numbers, prev)
	}

	fmt.Printf("Total: %v\n", total)
}
