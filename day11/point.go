package main

import (
	"fmt"
)

const (
	North int = iota
	East
	South
	West
	NorthWest
	NorthEast
	SouthWest
	SouthEast
)

var ClockwiseDirs = [4]int{North, East, South, West}

type Point struct {
	X int
	Y int
}

func (p Point) InBox(xMax, yMax int) bool {
	return (p.X >= 0 && p.X < xMax) && (p.Y >= 0 && p.Y < yMax)
}

func (p Point) Dist(other Point) int {
	return AbsDiff(p.X, other.X) + AbsDiff(p.Y, other.Y)
}


func (p Point) Move(dir int) Point {
	switch dir {
	case North:
		return Point{p.X, p.Y - 1}
	case South:
		return Point{p.X, p.Y + 1}
	case West:
		return Point{p.X - 1, p.Y}
	case East:
		return Point{p.X + 1, p.Y}
	case NorthWest:
		return Point{p.X - 1, p.Y - 1}
	case NorthEast:
		return Point{p.X + 1, p.Y - 1}
	case SouthWest:
		return Point{p.X - 1, p.Y + 1}
	case SouthEast:
		return Point{p.X + 1, p.Y + 1}
	default:
		return p
	}
}

func PairwiseDist(points []Point) []int {
	dists := make([]int, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := points[i].Dist(points[j])
			fmt.Printf("dist %v -> %v: %d\n", points[i], points[j], dist)
			dists = append(dists, dist)
		}
	}
	return dists
}

func PairwiseDistWithVoids(points []Point, rows, cols []int, voidSize int) []int {
	dists := make([]int, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := points[i].Dist(points[j])
			
			dx, dy := EmptyCountInPath(points[i], points[j], rows, cols)
			dx = dx * (voidSize - 1)
			dy = dy * (voidSize - 1)
			voidDist := dist + dx + dy	
			//fmt.Printf("dist %d -> %d : %v -> %v : %d\n", i, j, points[i], points[j], dist)
			//fmt.Printf("    + voids (%d, %d) -> %d\n", dx, dy, voidDist)
			dists = append(dists, voidDist)
		}
	}
	return dists
}

func AbsDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
