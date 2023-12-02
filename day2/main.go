package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(dataFile string, maxScans int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	scanner := bufio.NewScanner(file)

	i := 0
	possibleGameIdsSum := 0
	powerSum := 0
	for scanner.Scan() {
		if maxScans != -1 && i == maxScans {
			break
		}
		line := scanner.Text()
		gameStr, cubeSetStr, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		gameId, err := ParseGameId(gameStr)
		if err != nil {
			continue
		}
		cubeSets, err := ParseCubeSets(cubeSetStr)
		if err != nil {
			continue
		}
		maxCubeSet := CubeSetsMax(cubeSets)
		minPower := CubeSetPower(maxCubeSet)
		powerSum += minPower
		fmt.Printf("%s\n", line)
		fmt.Printf("Max cubes: %v\n", maxCubeSet)
		fmt.Printf("Set power %v\n", minPower)

		if IsSubset(maxCubeSet, NewCubeSet(12, 13, 14)) {
			possibleGameIdsSum += gameId
		}
		i++
	}

	fmt.Printf("Possible game id's sum: %d\n", possibleGameIdsSum)
	fmt.Printf("Power of all game cube sets: %d\n", powerSum)
	defer file.Close()
}

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
	parse(dataFile, maxScans)
}
