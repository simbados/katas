package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("123 456"))
}

func HighAndLow(in string) string {
	var newArr []int
	str := strings.Fields(in)
	for _, word := range str {
		asInt, _ := strconv.Atoi(string(word))
		newArr = append(newArr, asInt)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(newArr)))
	if len(newArr) == 1 {
		return fmt.Sprintf("%v %v", newArr[0], newArr[0])
	}
	return fmt.Sprintf("%v %v", newArr[0], newArr[len(newArr)-1])
}
