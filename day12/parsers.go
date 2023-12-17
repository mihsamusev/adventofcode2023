package main

import (
	"common"
	"fmt"
	"strings"
)


func ParseReading(str string) (string, []int) {
	parts := strings.Fields(str)
	slice, err := common.ParseDelimitedSlice(parts[1], ",")
	if err != nil {
		fmt.Errorf("dumbness happened")
	}
	return parts[0], slice
}