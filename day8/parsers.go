package main

import (
	"errors"
	"fmt"
	"strings"
)

func ParseCommands(s string) []Command {
	return []Command(s)
}

func ParseGraph(s string) (Graph, error) {
	graph := Graph{}

	transitions := strings.Split(s, "\n")

	for _, t := range transitions {
		n, f, err := ParseTransition(t)
		if err != nil {
			return graph, err
		}
		graph[n] = f
	}

	return graph, nil
}

func ParseTransition(s string) (Node, Fork, error) {
	node := Node("")
	fork := Fork{}
	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		return node, fork, errors.New("expected 2 parts")
	}

	node = Node(strings.TrimSpace(parts[0]))
	left, right, err := ParsePair(parts[1])
	if err != nil {
		return node, fork, err 
	}

	return node, Fork{Node(left), Node(right)}, nil
}

func ParsePair(s string) (string, string, error) {
	parts := strings.Split(s, ",")
	n := len(parts)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("'%s' -> expected 2 parts separated by , got %d", s, n)
	}
	first := strings.Trim(parts[0], "() ")
	second := strings.Trim(parts[1], "() ")
	return first, second, nil
}

