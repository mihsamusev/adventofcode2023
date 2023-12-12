package main

import (
	"common"
	"fmt"
)

func main() {
	args := common.ReadCliArgs()

	fmt.Println("SOLUTION 1")
	SolvePartOne(args)

	fmt.Println("SOLUTION 2")
	SolvePartTwo(args)
}
