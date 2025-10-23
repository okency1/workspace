package main

import (
	"fmt"
)

func CountCharacter(string, s c rune) int {
	if s == "" {
		return 0
	}
	count := 0
	for _, i := range s {
		if i == c {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountCharacter("", 'l'))
}
