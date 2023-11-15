package main

import (
	"fmt"
	"math"
)

func RGB(r, g, b int) string {
	red := expandHex(fmt.Sprintf("%X", int(math.Max(math.Min(float64(r), 255), 0))))
	green := expandHex(fmt.Sprintf("%X", int(math.Max(math.Min(float64(g), 255), 0))))
	blue := expandHex(fmt.Sprintf("%X", int(math.Max(math.Min(float64(b), 255), 0))))
	return red + green + blue
}

func expandHex(hex string) string {
	if len(hex) < 2 {
		hex = "0" + hex
	}
	return hex
}

func main() {
	fmt.Println(RGB(-20,275,125))
}
