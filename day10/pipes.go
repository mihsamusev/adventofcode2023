package main


const (
	North int = iota
	East
	South
	West
)


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
