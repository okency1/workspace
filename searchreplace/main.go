package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]

	if len(old) != 1 || len(new) != 1 {
		return
	}

	oldRune := []rune(old)[0]
	newRune := []rune(new)[0]

	for _, ch := range str {
		if ch == oldRune {
			z01.PrintRune(newRune)
		} else {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n') // Always print newline at the end
}
