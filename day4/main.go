package main

import (
	"common"
	"fmt"
)

func main() {
	args := common.ReadCliArgs()

	fmt.Println("SOLUTION PART 1")
	SolvePartOne(args)

	fmt.Println("SOLUTION PART 2")
	SolvePartTwo(args)
}
