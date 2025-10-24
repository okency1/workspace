package main

import "fmt"

func gcd(a, b uint) uint {
	// If either number is zero, return the other (standard GCD behavior)
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Euclidean algorithm for GCD
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
