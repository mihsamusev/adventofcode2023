package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parse(dataFile string, maxScans int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	thisLine := ""
	nextLine := ""
	prevLine := ""

	i := 0
	total := 0
	for {
		if maxScans != -1 && i == maxScans {
			break
		}
		if i == 0 && scanner.Scan() {
			thisLine = scanner.Text()
		}

		if scanner.Scan() {
			nextLine = scanner.Text()
		} else {
			nextLine = ""
		}

		numberRefs := ParseNumberRefs(thisLine)
		fmt.Printf("LOOKING FOR PARTS IN ROW %d\n", i)
		fmt.Printf("%d: %s\n", i-1, prevLine)
		fmt.Printf("%d: %s\n", i, thisLine)
		fmt.Printf("%d: %s\n", i+1, nextLine)

		fmt.Printf("PARTS FOUND: %d\n", len(numberRefs))
		filtered := FilterNumberRefs(numberRefs, prevLine, thisLine, nextLine)
		for _, f := range filtered {
			total += f.number
			fmt.Printf("    %v\n", f)
		}
		fmt.Println()

		prevLine = string([]byte(thisLine))
		thisLine = string([]byte(nextLine))
		i++

		if nextLine == "" {
			break
		}
	}
	fmt.Printf("Total engine number ref: %d\n", total)
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
