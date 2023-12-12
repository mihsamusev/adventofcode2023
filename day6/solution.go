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

	races, err := ParseRaces(string(content))
	if err != nil {
		fmt.Printf("Could not parse puzzle input: %v\n", err)
	}

	total := 1
	for _, r := range races {
		min, max := FindValidChargeTimes(r)
		okRaceCount := max - min + 1
		total *= okRaceCount
		fmt.Printf("Race %v: ok solutions: %d\n", r, okRaceCount)
	}
	fmt.Printf("Total ok races: %d\n", total)
}

func SolvePartTwo(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	race, err := ParseAsOneRace(string(content))
	if err != nil {
		fmt.Printf("Could not parse puzzle input: %v\n", err)
	}

	fmt.Printf("Race %v\n", race)
	min, max := FindValidChargeTimes(race)
	total := max - min + 1

	fmt.Printf("Total ok races: %d\n", total)
}
