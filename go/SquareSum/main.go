package main

func SquareSum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number * number
	}
	return result
}
