package main

import (
	"common"
	"fmt"
	"strings"
)

const DefaultEmpty rune = '.'

type Grid [][]rune

func ParseGrid(str string) Grid {
	rows := strings.Split(str, "\n")

	m := make(Grid, 0)
	for i := range rows {
		m = append(m, []rune(rows[i]))
	}
	return m
}

func (g Grid) At(p Point) rune {
	if !g.Contains(p) {
		return DefaultEmpty
	}
	return g[p.Y][p.X]
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	return len(g[0])
}

func (m Grid) Contains(p Point) bool {
	xIn := (p.X >= 0) && (p.X < m.Width())
	yIn := (p.Y >= 0) && (p.Y < m.Height())
	return xIn && yIn
}

func (m Grid) Update(p Point, symbol rune) {
	if m.Contains(p) {
		m[p.Y][p.X] = symbol
	}
}

func (m Grid) Find(symbol rune) []Point {
	positions := make([]Point, 0)
	for i := 0; i < m.Height(); i++ {
		for j := 0; j < m.Width(); j++ {
			if m[i][j] == symbol {
				positions = append(positions, Point{j, i})
			}
		}
	}
	return positions
}

func (m Grid) SymbolInRow(row int, symbol rune) bool {
	for j := 0; j < m.Width(); j++ {
		if m[row][j] == symbol {
			return true
		}
	}
	return false
}

func (m Grid) SymbolInColumn(col int, symbol rune) bool {
	for i := 0; i < m.Height(); i++ {
		if m[i][col] == symbol {
			return true
		}
	}
	return false
}

func (m Grid) EmptyRows() []int {
	rows := make([]int, 0)
	for i := 0; i < m.Height(); i++ {
		if !m.SymbolInRow(i, '#') {
			rows = append(rows, i)
		}
	}
	return rows
}

func (m Grid) EmptyColums() []int {
	cols := make([]int, 0)
	for i := 0; i < m.Width(); i++ {
		if !m.SymbolInColumn(i, '#') {
			cols = append(cols, i)
		}
	}
	return cols
}

func (m Grid) String() string {
	var builder strings.Builder
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Fprintf(&builder, "%c", m[i][j])
		}
		fmt.Fprintf(&builder, "\n")
	}
	return builder.String()
}

func EmptyCountInPath(p1, p2 Point, rows, cols []int) (int, int) {
	colCount := common.CountInRange(rows, p1.Y, p2.Y)
	rowCount := common.CountInRange(cols, p1.X, p2.X)
	return rowCount, colCount
}