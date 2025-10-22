package main

import "github.com/01-edu/z01"

func isnegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main() {
	isnegative(5)
	isnegative(-2)
	isnegative(0)
}
