package main

import (
	"common"
	"fmt"
	"os"
	"strings"
)

func SolvePartOne(args common.CliArgs) {
	bytes, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	rows := strings.Split(string(bytes), "\n")
	i := 0
	fmt.Printf("Row:\n%v\n", rows[i])
	springs, counts := ParseReading(rows[i])
	groups := Groups(springs)
	fmt.Printf("springs: %v  -> %v -> %v\n", springs, groups, counts)
	
	ok := CorrectPattern(springs, counts)
	fmt.Printf("correct?: %v\n", ok)
}

func SolvePartTwo(args common.CliArgs) {
	_, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}
}
