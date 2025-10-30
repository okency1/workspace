package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	s, space := os.Args[1], false
	for _, c := range s {
		if c == ' ' || c == '\t' {
			space = len(s) > 0
			continue
		}
		if space {
			z01.PrintRune(' ')
			space = false
		}
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
