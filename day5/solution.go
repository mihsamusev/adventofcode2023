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

	min := math.MaxInt
    for i := 0; i < len(seeds) / 2; i++ {
        seedStart := seeds[2 * i]
        length := seeds[2 * i + 1]
        for j := 0; j < length; j++ {
			seed := seedStart + j
			location := Trace(seed, farmingMaps)
			if location < min {
				min = location
			}
		}
	}
	fmt.Println(min)
}
