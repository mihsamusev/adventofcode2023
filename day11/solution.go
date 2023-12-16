package main

import (
	"common"
	"fmt"
	"os"
)

func SolvePartOne(args common.CliArgs) {
	bytes, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	grid := ParseGrid(string(bytes))
	fmt.Printf("Grid:\n%v\n", grid)

	stars := grid.Find('#')
	fmt.Printf("Stars:\n%v\n", stars)

	rows := grid.EmptyRows()
	cols := grid.EmptyColums()
	fmt.Printf("Void rows:\n%v\n", rows)
	fmt.Printf("Void cols:\n%v\n", cols)

	voidSize := 1000000
	dists := PairwiseDistWithVoids(stars, rows, cols, voidSize)
	total := common.Sum(dists)
	fmt.Printf("Sum of dists: %v\n", total)

}

func SolvePartTwo(args common.CliArgs) {
	_, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

}
