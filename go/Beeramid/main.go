package main

import (
	"math"
)

func Beeramid(bonus int, price float64) int {
	if float64(bonus) < price {
		return 0
	}
	cans := int(math.Floor(float64(bonus) / price))
	level := 1
	currentCans := 1
	for currentCans+((level+1)*(level+1)) <= cans {
		level += 1
		currentCans = currentCans + (level * level)
	}
	return level
}
