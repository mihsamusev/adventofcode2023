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
	dist.Update(start, 1)
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
	fmt.Printf("Max steps: %v\n", maxSteps - 1)
}

func SolvePartTwo(args common.CliArgs) {
	bytes, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	m := ParseMap(string(bytes))
	dist := NewDist(len(m), len(m[0]))
	fmt.Printf("%v\n", m)

	start := m.StartPos('S')
	dist.Update(start, 1)
	fmt.Printf("Start: %v\n", start)

	queue := make(PositionQueue, 0)
	queue.Enqueue(start)

	dirs := [4]int{North, East, South, West}
	path := []int{}
	
	for len(queue) > 0 {
		fromPos := queue.Dequeue()

		for _, dir := range dirs {
			fromPipe := m.At(fromPos)
			toPos := fromPos.Move(dir)
			toPipe := m.At(toPos)

			canMove := CanMove(fromPipe, toPipe, dir)
			notSeen := dist.At(toPos) <= 0

			if canMove && notSeen {
				fmt.Printf("Moving: %v: %c -> %v: %c\n", fromPos, fromPipe, toPos, toPipe)
				queue.Enqueue(toPos)
				path = append(path, dir)
				newValue := dist.At(fromPos) + 1
				dist.Update(toPos, newValue)
				break
			}
		}
	}

	m.UpdateMap(dist)
	MarkVoids(start, path, m, dist)
	markedBorder := dist.CountBorder(-1)
	count := dist.Count(-1)
	if markedBorder > 0 {
		count = dist.Count(0)
	}
	fmt.Printf("\n%v\n", dist)
	fmt.Printf("Marked border: %d\n", markedBorder)
	fmt.Printf("Void count: %d\n", count)
}


func (d Dist) UpdateWhileEmpty(m Map, start Pos, dir int) {
	p := start
	for d.CanUpdate(p) && m.CanUpdate(p) {
		fmt.Printf("    Updated: %v\n", p)
		d.Update(p, -1)
		p = p.Move(dir)
	}
}

func MarkVoids(start Pos, path []int, m Map, d Dist) {
	fromPos := start
	for i, dir := range path {
		toPos := fromPos.Move(dir)
		toPipe := m.At(toPos)

		testPos := PosToTest(toPos, toPipe, dir)
		fmt.Printf("Marks %d: %v -> %v\n", i + 1, toPos, testPos)
	
		for _, t := range testPos {
			d.UpdateWhileEmpty(m, t, North)
			d.UpdateWhileEmpty(m, t, South)
			d.UpdateWhileEmpty(m, t, East)
			d.UpdateWhileEmpty(m, t, West)
		}
		fromPos = toPos
	}
}
