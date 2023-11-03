package main

// Kata: https://www.codewars.com/kata/55c45be3b2079eccff00010f

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a"))
}

func Order(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}
	words := strings.Split(sentence, " ")
	wordMap := map[int]string{}
	for _, word := range words {
		addToMap(word, wordMap)
	}
	var allKeys []int
	for key := range wordMap {
		allKeys = append(allKeys, key)
	}
	sort.Ints(allKeys)
	var sortedSentence = ""
	for _, value := range allKeys {
		sortedSentence += wordMap[value] + " "
	}
	return strings.Trim(sortedSentence, " ")
}

func addToMap(word string, wordMap map[int]string) map[int]string {
	for _, char := range word {
		if isADigit(char) {
			i, _ := strconv.Atoi(string(char))
			value, exists := wordMap[i]
			if exists {
				wordMap[i] += value + " " + word
			} else {
				wordMap[i] = word
			}
		}
	}
	return wordMap
}

func isADigit(toTest rune) bool {
	return '0' < toTest && toTest <= '9'
}
