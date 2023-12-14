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
	case NorthWest:
		return Pos{p.x - 1, p.y - 1}
	case NorthEast:
		return Pos{p.x + 1, p.y - 1}
	case SouthWest:
		return Pos{p.x - 1, p.y + 1}
	case SouthEast:
		return Pos{p.x + 1, p.y + 1}
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

func (m Map) CanUpdate(p Pos) bool {
	yMax := len(m)
	xMax := len(m[0])
	return p.InBox(xMax, yMax) && m.At(p) == '.'
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

func (m Map) UpdateMap(d Dist) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if d[i][j] <= 0 {
				m[i][j] = '.'
			}
		}
	}
}

type Dist [][]int

func NewDist(rows, cols int) Dist {
	dist := make([][]int, rows)
		for i := range dist {
			dist[i] = make([]int, cols)
		}
	return dist
}

func (d Dist) Width() int {
	return len(d[0])
}

func (d Dist) Height() int {
	return len(d)
}

func (d Dist) At(p Pos) int {
	yMax := len(d)
	xMax := len(d[0])
	if !p.InBox(xMax, yMax) {
		return 0
	}
	return d[p.y][p.x]
}

func (d Dist) CanUpdate(p Pos) bool {
	yMax := len(d)
	xMax := len(d[0])
	return p.InBox(xMax, yMax) && d.At(p) <= 0
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
			fmt.Fprintf(&builder, "%4d", d[i][j])
		}
		fmt.Fprintf(&builder, "\n")
	}
	return builder.String()
}

func (d Dist) Count(v int) int {
	count := 0
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[0]); j++ {
			if d[i][j] == v {
				count++
			}
		}
	}
	return count
}

func (d Dist) CountBorder(v int) int {
	count := 0
	iLast := d.Height() - 1
	jLast := d.Width() - 1
	for i := 1; i < iLast; i++ {
		if d[i][0] == v || d[i][jLast] == v {
			count++
		}
	}

	for j := 0; j <= jLast; j++ {
		if d[0][j] == v || d[iLast][j] == v {
			count++
		}
	}
	return count	
}