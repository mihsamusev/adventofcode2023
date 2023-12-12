package main

import (
	"common"
	"errors"
	"math"
	"strings"
)

type Race struct {
	time int
	dist int
}

func ParseRaces(str string) ([]Race, error) {
	races := make([]Race, 0)

	parts := strings.Split(str, "\n")
	if len(parts) < 2 {
		return races, errors.New("not enough rows in parsed string")
	}

	times, err := common.ParseNamedSlice(parts[0], "Time:")
	if err != nil {
		return races, errors.New("could not parse times")
	}

	dists, err := common.ParseNamedSlice(parts[1], "Distance:")
	if err != nil {
		return races, errors.New("could not parse distances")
	}

	if len(times) != len(dists) {
		return races, errors.New("parsed times and distances size mismatch")
	}

	for i := range times {
		races = append(races, Race{times[i], dists[i]})
	}
	return races, nil
}

func ParseAsOneRace(str string) (Race, error) {
	parts := strings.Split(str, "\n")
	if len(parts) < 2 {
		return Race{}, errors.New("not enough rows in parsed string")
	}

	timeParts, err := common.ParseNamedSlice(parts[0], "Time:")
	if err != nil {
		return Race{}, errors.New("could not parse times")
	}

	distParts, err := common.ParseNamedSlice(parts[1], "Distance:")
	if err != nil {
		return Race{}, errors.New("could not parse distances")
	}

	time := ConcatSlice(timeParts)
	dist := ConcatSlice(distParts)
	return Race{time, dist}, nil
}

func ConcatSlice(slice []int) int {
	order := 0
	total := 0.0
	for i := range slice {
		pos := len(slice) - 1 - i
		number := float64(slice[pos])
		total += number * math.Pow10(order)
		numberOrder := int(math.Log10(number))
		order += 1 + numberOrder
	}
	return int(total)
}
