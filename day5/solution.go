package main

import (
	"common"
	"fmt"
	"math"
	"os"
	"strings"
)

func SolvePartOne(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("im dead")
	}
	blocks := strings.Split(string(content), "\n\n")

	seeds, err := common.ParseNamedSlice(blocks[0], "seeds:")
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

func SolvePartTwo(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("im dead")
	}
	blocks := strings.Split(string(content), "\n\n")

	seeds, err := common.ParseNamedSlice(blocks[0], "seeds:")
	if err != nil {
		panic(err)
	}
	seedRanges := InterpretSlice(seeds)
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