package piscine

import "github.com/01-edu/z01"

func PrintMemory(a [10]byte) {
	for i, v := range a {
		printHex(v/16)
		printHex(v%16)
		if (i+1)%4 == 0 || i == len(a)-1 {
			z01.PrintRune('$')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, v := range a {
		if v >= 32 && v <= 126 {
			z01.PrintRune(rune(v))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('$')
	z01.PrintRune('\n')
}

func printHex(n byte) {
	if n < 10 {
		z01.PrintRune(rune('0' + n))
	} else {
		z01.PrintRune(rune('a' + n - 10))
	}
}
