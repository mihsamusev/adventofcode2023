package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)


func to_digit(r rune) int {
	return int(r - '0')
}

func line_value(str string) int {
	value := 0
	digits := make([]int, 0, 10)
	for _, rune := range str {
		if unicode.IsDigit(rune) {
			digits = append(digits, to_digit(rune))
		}
	}
	numDigits := len(digits)
	if numDigits != 0 {
		first := digits[0]
		last := digits[numDigits - 1]
		value = first*10 + last
	}

	return value
}

func parse(dataFile string) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		value := line_value(line)
		total += value
		fmt.Printf("%s -> %d\n", line, value)
	}
	fmt.Printf("total = %d\n", total)
	

	defer file.Close()
}

func main() {
	dataFile := "test.txt"
	args := os.Args
	if len(args) > 1 {
		dataFile = args[1]
	}
	parse(dataFile)
}
