package main

// kata: https://www.codewars.com/kata/517abf86da9663f1d2000003

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile("[-_]")
	runeStr := []rune(s)
	for _, value := range re.FindAllStringSubmatchIndex(s, -1) {
		runeStr[value[1]] = []rune(strings.ToUpper(string(runeStr[value[1]])))[0]
	}
	return re.ReplaceAllString(string(runeStr), "")
}
