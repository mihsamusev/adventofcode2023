package main

type CubeSet map[string]int

var CubeColors = [3]string{"red", "green", "blue"}

func NewCubeSet(red, green, blue int) CubeSet {
	return CubeSet{
		CubeColors[0]: red,
		CubeColors[1]: green,
		CubeColors[2]: blue,
	}
}
func IsSubset(subset, superset CubeSet) bool {
	for _, color := range CubeColors {
		if subset[color] > superset[color] {
			return false
		}
	}
	return true
}

func CubeSetPower(cubeSet CubeSet) int {
	power := 1
	for _, color := range CubeColors {
		power *= cubeSet[color]
	}
	return power
}

func CubeSetsMax(cubeSets []CubeSet) CubeSet {
	maxCubeSet := NewCubeSet(0, 0, 0)
	for _, cubeSet := range cubeSets {
		for _, color := range CubeColors {
			if maxCubeSet[color] < cubeSet[color] {
				maxCubeSet[color] = cubeSet[color]
			}
		}
	}
	return maxCubeSet
}

func CubeSetsSum(cubeSets []CubeSet) CubeSet {
	sumCubeSet := NewCubeSet(0, 0, 0)
	for _, cubeSet := range cubeSets {
		for _, color := range CubeColors {
			sumCubeSet[color] += cubeSet[color]
		}
	}
	return sumCubeSet
}
