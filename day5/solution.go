package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func SolvePartOne(dataFile string) {
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
	fmt.Println(farmingMaps)

	ranges := InterpretRanges(seeds)
	min := math.MaxInt
	for _, r := range ranges {
		locationRanges := TraceRange(r, farmingMaps)
		for _, loc := range locationRanges {
			if loc.start < min {
				min = loc.start
			}
		}
	}

	fmt.Println(min)
}
