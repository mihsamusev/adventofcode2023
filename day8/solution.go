package main

import (
	"common"
	"fmt"
	"os"
	"strings"
)

func SolvePartOne(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}

	parts := strings.Split(string(content), "\n\n")
	commands := ParseCommands(parts[0])
	graph, err := ParseGraph(parts[1])

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	traps := FindTraps(graph)
	fmt.Printf("Traps: %v\n", traps)

	RemoveTraps(graph, traps)
	fmt.Printf("Graph without traps:\n%v\n", graph)

	commandLoop := CommandLoop{commands: commands}
	start := Node("AAA")

	fmt.Printf("Start: %v\n", start)
	steps, target := SearchPath(start, graph, commandLoop, FoundAllZ)
	fmt.Printf("Target %v found\n", target)
	fmt.Printf("Steps: %d\n", steps)
}

func SolvePartTwo(args common.CliArgs) {
	content, err := os.ReadFile(args.FileName)
	if err != nil {
		fmt.Println("No puzzle input file")
	}
	parts := strings.Split(string(content), "\n\n")
	commands := ParseCommands(parts[0])
	graph, err := ParseGraph(parts[1])

	if err != nil {
		fmt.Printf("Could not parse puzzle input: %s\n", err)
	}

	traps := FindTraps(graph)
	fmt.Printf("Traps: %v\n", traps)

	RemoveTraps(graph, traps)

	//starts := FindStarts(graph)
	starts := []Node{"JHA", "DTA", "MMA", "AAA", "NCA", "TVA"}
	counts := make([]int, 0)
	for _, start := range starts {
		commandLoop := CommandLoop{commands: commands}
		count, _ := SearchPath(start, graph, commandLoop, FoundLastZ)
		counts = append(counts, count)
	}

	fmt.Printf("Individual counts: %v\n", counts)
	commonCount := LeastCommonMul(counts...)
	fmt.Printf("Least common multiple of counts: %v\n", commonCount)
}
