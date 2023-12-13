package main

import (
	"common"
	"fmt"
	"os"
)

func SolvePartOne(args common.CliArgs) {
	_, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	total := 0
	fmt.Printf("Total: %v\n", total)
}

func SolvePartTwo(args common.CliArgs) {
	_, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	total := 0
	fmt.Printf("Total: %v\n", total)
}
