package main

import "fmt"

func gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
}

func main() {
	fmt.Println(gcd(12, 6))
}
