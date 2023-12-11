package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func SolvePartOne(dataFile string) {
	fmt.Println("SOLUTION 1")

	content, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	blocks := strings.Split(string(content), "\n\n")

	seeds, err := ParseNamedSlice(blocks[0], "seeds:")
	if err != nil {
		panic(err)
	}
	fmt.Println(seeds)

	farmingMaps, err := ParseFarmingMaps(blocks[1:])
	if err != nil {
		panic(err)
	}

	min := math.MaxInt
	for _, seed := range seeds {
		location := Trace(seed, farmingMaps)
		if location < min {
			min = location
		}
	}

	fmt.Println(min)
}

func SolvePartTwo(dataFile string) {
	fmt.Println("SOLUTION 2")

	content, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	blocks := strings.Split(string(content), "\n\n")

	seeds, err := ParseNamedSlice(blocks[0], "seeds:")
	if err != nil {
		panic(err)
	}
	seedRanges := SliceAsRanges(seeds)
	fmt.Println(seedRanges)

	farmingMaps, err := ParseFarmingMaps(blocks[1:])
	if err != nil {
		panic(err)
	}

	minRanges := TraceRanges(seedRanges, farmingMaps)

	min := math.MaxInt
	for _, r := range minRanges {
		if r.start < min {
			min = r.start
		}
	}

	fmt.Printf("Min seed location %d\n", min)
}
