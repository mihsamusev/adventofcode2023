package main

import "fmt"


type Node string

type Fork struct {
	left  Node
	right Node
}

type Graph map[Node]Fork

func FindTraps(g Graph) []Node {
	traps := make([]Node, 0)
	for key, value := range g {
		if (key == value.left) && (key == value.right) {
			traps = append(traps, key)
		}
	}
	return traps
}

func RemoveTraps(g Graph, traps []Node) {
	for _, t := range traps {
		delete(g, t)
	}
}

func SearchPath(start Node, graph Graph, commandLoop CommandLoop, StopFn func(Node) bool) (int, Node) {
	steps := 0
	next := Node("XXX")
	for {
		command := commandLoop.Next()
		switch command {
		case Left:
			next = graph[start].left
		case Right:
			next = graph[start].right
		}
		fmt.Printf("Command: %v from %v -> to %v\n", command, start, next)
		start = next
		steps++
		if StopFn(next) {
			break
		}
	}
	return steps, next
}

func FindStarts(g Graph) []Node {
	starts := make([]Node, 0)
	for key := range g {
		if key[2] == 'A' {
			starts = append(starts, key)
		}
	}
	return starts
}

func FoundAllZ(n Node) bool {
	return n == "ZZZ"
}

func FoundLastZ(n Node) bool {
	return n[2] == 'Z'
}
