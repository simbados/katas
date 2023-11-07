package main

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	start := numbers[0:3]
	middle := numbers[3:6]
	end := numbers[6:10]
	res := "("
	for _, number := range start {
		res += strconv.Itoa(int(number))
	}
	res += ") "
	for _, number := range middle {
		res += strconv.Itoa(int(number))
	}
	res += "-"
	for _, number := range end {
		res += strconv.Itoa(int(number))
	}
	return res
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0}))
}
