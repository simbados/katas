package main

// https://www.codewars.com/kata/556deca17c58da83c00002db

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return make([]float64, 0)
	} else if n <= 3 {
		lessThanThree := []float64{signature[0]}
		for j := 1; j < n; j++ {
			lessThanThree = append(lessThanThree, signature[j])
		}
		return lessThanThree
	}
	res := []float64{
		signature[0], signature[1], signature[2],
	}
	for i := n; i > 3; i-- {
		res = append(res, res[len(res) - 1] + res[len(res) - 2] + res[len(res) - 3])
	}
	return res
}

func main() {
	fmt.Println(Tribonacci([3]float64{15, 1, 15}, 3))
}

