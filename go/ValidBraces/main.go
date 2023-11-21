package main

import "fmt"

func main() {
	fmt.Println(ValidBraces("[(){}()]"))
}

var validOpeningBraces = [4]string{"{", "[", "[", "("}
var validClosingBraces = [4]string{"}", "]", "]", ")"}
var closingBracesMap = map[string]string{ "{": "}", "[": "]", "(": ")" }
func ValidBraces(str string) bool {
	var allClosing []string
	for _, value := range str {
		if contains := containsOpening(value); contains {
			allClosing =  append(allClosing, closingBracesMap[string(value)])
		} else if contains := containClosing(value); contains {
			if len(allClosing) == 0 || allClosing[len(allClosing) - 1] != string(value) {
				return false
			}
			allClosing = allClosing[:len(allClosing) - 1]
		}
	}
	return len(allClosing) == 0
}

func containClosing(char rune) bool {
	for _, value := range validClosingBraces {
		if value == string(char) {
			return true;
		}
	}
	return false
}

func containsOpening(char rune) bool {
	for _, value := range validOpeningBraces {
		if value == string(char) {
			return true;
		}
	}
	return false
}

