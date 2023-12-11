package main

import (
	"common"
)

func main() {
	args := common.ReadCliArgs()
	SolvePartOne(args)
	SolvePartTwo(args)
}
