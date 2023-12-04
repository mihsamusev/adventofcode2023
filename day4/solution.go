package main

import (
    "os"
    "bufio"
    "fmt"
)

func Parse(dataFile string, maxScans int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()


	i := 0
	total := 0
	for {
		if maxScans != -1 && i == maxScans {
			break
		}
		if !scanner.Scan() {
            break
		}
        line := scanner.Text()
		fmt.Printf("%d: %s\n", i+1, line)


		fmt.Println()
		i++

	}
    fmt.Printf("Total: %d\n", total)
}
