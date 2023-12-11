package main

import (
	"common"
	"errors"
	"strings"
)


func ParseFarmingMap(str string, withHeader bool) (FarmingMap, error) {
	lines := strings.Split(str, "\n")
	if len(lines) == 0 {
		return FarmingMap{}, errors.New("rows not found")
	}

	if withHeader {
		lines = lines[1:]
	}

	lookups := make([]Lookup, 0)
	for _, line := range lines {
		n, err := common.ParseSlice(line)
		if err != nil {
			return FarmingMap{}, errors.New("wrong row")
		}
		if len(n) != 3 {
			return FarmingMap{}, errors.New("expected 3 numbers")
		}
		dst := n[0]
		srcRange := Range{n[1], n[1] + n[2] - 1}
		lookups = append(lookups, Lookup{dst, srcRange})
	}

	return FarmingMap{lookups}, nil
}

func ParseFarmingMaps(lines []string) ([]FarmingMap, error) {
	farmingMaps := make([]FarmingMap, 0)
	for _, b := range lines {
		farmingMap, err := ParseFarmingMap(b, true)
		if err != nil {
			panic(err)
		}
		farmingMaps = append(farmingMaps, farmingMap)
	}
	return farmingMaps, nil
}



