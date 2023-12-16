package main

import (
	"common"
	"fmt"
)

func main() {
	args := common.ReadCliArgs()

	fmt.Println("SOLUTION 1")
	SolvePartOne(args)

	fmt.Println("\nSOLUTION 2")
	SolvePartTwo(args)
}
