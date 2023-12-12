package main

import (
	"fmt"
	"math"
)

func SolveEq(a, b, c float64) (left, right float64) {
	det2 := b*b - 4*a*c
	if det2 < 0 {
		return math.NaN(), math.NaN()
	}

	det := math.Sqrt(det2)

	left = (-b + det) / (2 * a)
	right = (-b - det) / (2 * a)
	if left > right {
		right, left = left, right
	}
	return left, right
}

func FindValidChargeTimes(r Race) (minCharge, maxCharge int) {
	speed := 1
	left, right := SolveEq(float64(-speed), float64(r.time), float64(-r.dist))

	minCharge = int(math.Ceil(left))
	if speed * minCharge * (r.time - minCharge) == r.dist {
		minCharge++
	}

	maxCharge = int(math.Floor(right))
	if speed * maxCharge * (r.time - maxCharge) == r.dist {
		maxCharge--
	}
	fmt.Printf("Race %v: solutions: (%.2f, %.2f), (min, max) = (%d, %d)\n", r, left, right, minCharge, maxCharge)
	return minCharge, maxCharge
}
