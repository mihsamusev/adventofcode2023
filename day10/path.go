package main

import (
	"fmt"
	"strings"
)

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

func (d Dist) UpdateLine(start Pos, dir int) {
	p := start
	for d.CanUpdate(p) {
		fmt.Printf("    Updated: %v\n", p)
		d.Update(p, -1)
		p = p.Move(dir)
	}
}

func (d Dist) MarkVoids(start Pos, path []int, m Map) {
	fromPos := start
	for i, dir := range path {
		toPos := fromPos.Move(dir)
		toPipe := m.At(toPos)

		testPos := PosToTest(toPos, toPipe, dir)
		fmt.Printf("Marks %d: %v -> %v\n", i+1, toPos, testPos)

		for _, t := range testPos {
			d.UpdateLine(t, North)
			d.UpdateLine(t, South)
			d.UpdateLine(t, East)
			d.UpdateLine(t, West)
		}
		fromPos = toPos
	}
}
