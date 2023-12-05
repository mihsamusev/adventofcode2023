package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dataFile := "test_1.txt"
	maxScans := -1
	args := os.Args
	if len(args) > 1 {
		dataFile = args[1]
	}

	if len(args) > 2 {
		maxScans, _ = strconv.Atoi(args[2])
	}

	fmt.Printf("Analyzing %d lines of %s\n", maxScans, dataFile)
	//SolvePartOne(dataFile)
	SolvePartTwo(dataFile)
}
