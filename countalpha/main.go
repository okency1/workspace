package main

import "github.com/01-edu/z01"

func CountAlpha(s string) int {
	count := 0
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			count++
		}
	}
	return count
}
func main() {
	result := CountAlpha("christopheA 123!")
	z01.PrintRune('0' + rune(result/10))
	z01.PrintRune('0' + rune(result%10))
	z01.PrintRune('\n')
}
