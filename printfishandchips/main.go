package main

import "github.com/01-edu/z01"

func PrintFish(s string, n int) {
	if n < 0 {
		print("error")
	}
	if n%2 == 0 && n%3 == 0 {
		print("fish and chips")
	}
	if n%2 == 0 {
		print("fish")
	}
	if n%3 == 0 {
		print("chips")
	}
}

func main() {
	PrintFish("hello", 6)
	z01.PrintRune('\n')
	PrintFish("hello", 4)
	z01.PrintRune('\n')
	PrintFish("hello", 9)
	z01.PrintRune('\n')
	PrintFish("hello", 5)
	z01.PrintRune('\n')
	PrintFish("test", -1)
	z01.PrintRune('\n')
}
