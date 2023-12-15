package main




var okMoves = map[rune]map[int][3]rune{
	'|': {
		North: {'|', '7', 'F'},
		South: {'|', 'J', 'L'},
	},
	'-': {
		East: {'-', 'J', '7'},
		West: {'-', 'F', 'L'},
	},
	'7': {
		West: {'-', 'F', 'L'}, 
		South: {'|', 'J', 'L'},
	},
	'F': {
		East: {'-', 'J', '7'},
		South: {'|', 'J', 'L'},
	},
	'J': {
		West: {'-', 'F', 'L'}, 
		North: {'|', '7', 'F'},
	},
	'L': {
		East: {'-', 'J', '7'},
		North: {'|', '7', 'F'},
	},
	'S': {
		North: {'|', '7', 'F'},
		South: {'|', 'J', 'L'},
		East: {'-', 'J', '7'},
		West: {'-', 'F', 'L'},
	},
	'.': {},
}

func CanMove(from, to rune, dir int) bool {
	okOptions, exist:= okMoves[from][dir]
	if !exist {
		return false
	}
	for _, option := range okOptions {
		if option == to {
			return true
		}
	}
	return false
}

var testMovesCCW = map[int]map[rune][]int{
	North: {
		'|': {East},
		'F': {},
		'7': {East, North, NorthEast},
	},
	South: {
		'|': {West},
		'J': {},
		'L': {West, South, SouthWest},
	},
	West: {
		'-': {North},
		'L': {},
		'F': {North, West, NorthWest},
	},
	East: {
		'-': {South},
		'7': {},
		'J': {South, East, SouthEast},
	},
}


func PosToTest(thisPos Pos, nextPipe rune, flowDir int) []Pos {
	testDirs := testMovesCCW[flowDir][nextPipe]
	positions := make([]Pos, 0)
	for _, d := range testDirs {
		p := thisPos.Move(d)
		if (p != Pos{}) {
			positions = append(positions, p)
		}
	}
	return positions
}