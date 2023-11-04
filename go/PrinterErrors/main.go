package main

// kata: https://www.codewars.com/kata/56541980fa08ab47a0000040

import (
	"fmt"
	"regexp"
	"strconv"
)

func PrinterError(s string) string {
	re := regexp.MustCompile("[a-m]")
	correctMatches := re.FindAllStringSubmatch(s, -1)
	return strconv.Itoa(len(s) - len(correctMatches)) + "/" + strconv.Itoa(len(s))
}

func main() {
	fmt.Println(PrinterError("kkkaaaaabbbbbmmmmxxxxxyyy"))
}
