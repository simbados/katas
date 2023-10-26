package main

import "strings"

func GetMiddle(s string) string {
	if len(s) == 1 {
		return s;
	}
	if len(s) % 2 == 0 {
		var builder strings.Builder
		builder.WriteByte(s[(len(s) / 2) - 1])
		builder.WriteByte(s[len(s) / 2])
		return builder.String()
	} else {
		return string(s[len(s) / 2])
	}
}
