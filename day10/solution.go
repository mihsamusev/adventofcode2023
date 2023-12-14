package main

import (
	"common"
	"fmt"
	"os"
)

func SolvePartOne(args common.CliArgs) {
	bytes, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	m := ParseMap(string(bytes))
	dist := NewDist(len(m), len(m[0]))
	fmt.Printf("%v\n", m)



	start := m.StartPos('S')
	fmt.Printf("Start: %v\n", start)

	queue := make(PositionQueue, 0)
	queue.Enqueue(start)

	maxSteps := 0
	dirs := [4]int{North, East, South, West}
	
	for len(queue) > 0 {
		fromPos := queue.Dequeue()

		for _, dir := range dirs {
			fromPipe := m.At(fromPos)
			toPos := fromPos.Move(dir)
			toPipe := m.At(toPos)

			canMove := CanMove(fromPipe, toPipe, dir)
			notSeen := dist.At(toPos) == 0

			if canMove && notSeen {
				fmt.Printf("Moving: %v: %c -> %v: %c\n", fromPos, fromPipe, toPos, toPipe)
				queue.Enqueue(toPos)
				newValue := dist.At(fromPos) + 1
				if newValue > maxSteps {
					maxSteps = newValue
				}
				dist.Update(toPos, newValue)
			}
		}
	}
	fmt.Printf("\n%v\n", dist)
	fmt.Printf("Max steps: %v\n", maxSteps)
}

func SolvePartTwo(args common.CliArgs) {
	_, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	total := 0
	fmt.Printf("Total: %v\n", total)
}
