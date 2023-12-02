package main

import (
	"strconv"
	"strings"
)


func ParseGameId(s string) (int, error) {
	result, _ := strings.CutPrefix(s, "Game")
	result = strings.TrimSpace(result)
	return strconv.Atoi(result)
}

func ParseCubeSets(s string) ([]CubeSet, error) {
	cubeSets := make([]CubeSet, 0)
	cubeSetStr := strings.Split(s, ";")
	for _, str := range cubeSetStr {
		cubeSet, err := ParseCubeSet(str)
		if err != nil {
			return cubeSets, err
		}
		cubeSets = append(cubeSets, cubeSet)
	}
	return cubeSets, nil
}

func ParseCubeSet(s string) (CubeSet, error) {
	cubeSet := NewCubeSet(0, 0, 0)
	cubeSetStr := strings.Split(s, ",")

	for _, str := range cubeSetStr {
		for _, color := range CubeColors {
			value, err := ParseColor(str, color)
			if err != nil {
				return CubeSet{}, err
			}
			cubeSet[color] += value
		}
	}
	return cubeSet, nil
}

func ParseColor(s, color string) (int, error) {
	result, found := strings.CutSuffix(s, color)
	if !found {
		return 0, nil
	}
	result = strings.TrimSpace(result)
	return strconv.Atoi(result)
}