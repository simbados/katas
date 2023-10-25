package main

import (
	"fmt"
)

func main() {
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	percentValue := percent / 100
	return calcYears(p0, percentValue, aug, p, 0)
}

func calcYears(p0 int, percentValue float64, aug int, p int, years int) int {
	if p0 >= p {
		return years
	} else {
		newVal := p0 + int(float64(p0) * percentValue) + aug
		return calcYears(newVal, percentValue, aug, p, years + 1)
	}
}


