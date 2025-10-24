package main

import "github.com/01-edu/z01"

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w + h)
}

func main() {
	z01.PrintRune(rune(RectPerimeter(1, 6)))
}
