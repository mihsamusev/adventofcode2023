package main

import (
	"fmt"
	"strings"
)

type Pos struct {
	x int
	y int
}

func (p Pos) InBox(xMax, yMax int) bool {
	return (p.x >= 0 && p.x < xMax) && (p.y >= 0 && p.y < yMax)
}

type Map [][]rune

func (p Pos) Move(dir int) Pos {
	switch dir {
	case North:
		return Pos{p.x, p.y - 1}
	case South:
		return Pos{p.x, p.y + 1}
	case West:
		return Pos{p.x - 1, p.y}
	case East:
		return Pos{p.x + 1, p.y}
	default:
		return p
	}
}

func (m Map) At(p Pos) rune {
	yMax := len(m)
	xMax := len((m)[0])
	if !p.InBox(xMax, yMax) {
		return '.'
	}
	return m[p.y][p.x]
}

func (m Map) Neigbour(p Pos, dir int) rune {
	p = p.Move(dir)
	return m.At(p)
}

func (m Map) StartPos(symbol rune) Pos {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == symbol {
				return Pos{j, i}
			}
		}
	}
	return Pos{-1, -1}
}

func (m Map) String() string {
	var builder strings.Builder
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Fprintf(&builder, "%c", m[i][j])
		}
		fmt.Fprintf(&builder, "\n")
	}
	return builder.String()
}

type Dist [][]int

func NewDist(rows, cols int) Dist {
	dist := make([][]int, rows)
		for i := range dist {
			dist[i] = make([]int, cols)
		}
	return dist
}

func (d Dist) At(p Pos) int {
	yMax := len(d)
	xMax := len(d[0])
	if !p.InBox(xMax, yMax) {
		return 0
	}
	return d[p.y][p.x]
}

func (d Dist) Update(p Pos, value int) {
	d[p.y][p.x] = value
}

func (d Dist) Neigbour(p Pos, dir int) int {
	p = p.Move(dir)
	return d.At(p)
}

func (d Dist) String() string {
	var builder strings.Builder
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[0]); j++ {
			fmt.Fprintf(&builder, "%d", d[i][j])
		}
		fmt.Fprintf(&builder, "\n")
	}
	return builder.String()
}