package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parse(dataFile string, maxScans int, digitsFromLineFn func(string) []int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	scanner := bufio.NewScanner(file)

	total := 0
	i := 0
	for scanner.Scan() {
		if maxScans != -1 && i == maxScans {
			break
		}
		line := scanner.Text()
		digits := digitsFromLineFn(line)
		value := calibrationValue(digits)
		fmt.Printf("%s -> %v -> %d\n", line, digits, value)
		total += value
		i++
	}
	fmt.Printf("total = %d\n", total)

	defer file.Close()
}

func main() {
	dataFile := "test.txt"
	maxScans := -1
	args := os.Args
	if len(args) > 1 {
		dataFile = args[1]
	}

	if len(args) > 2 {
		maxScans, _ = strconv.Atoi(args[2])
	}

	fmt.Printf("Analyzing %d lines of %s\n", maxScans, dataFile)
	parse(dataFile, maxScans, lineDigitsExtended)
}
