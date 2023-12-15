package main

import (
	"fmt"
	"strings"
)

type Map [][]rune

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

func (m Map) RemoveUnreacheable(d Dist) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if d[i][j] <= 0 {
				m[i][j] = '.'
			}
		}
	}
}
