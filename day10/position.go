package main

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

type Pos struct {
	x int
	y int
}

func (p Pos) InBox(xMax, yMax int) bool {
	return (p.x >= 0 && p.x < xMax) && (p.y >= 0 && p.y < yMax)
}


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