package main

import "fmt"

func Digitlen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		n = n / base
		count++
	}
	return count
}

func main() {
	fmt.Println(Digitlen(100, 10))
	fmt.Println(Digitlen(100, 2))
	fmt.Println(Digitlen(-100, 16))
	fmt.Println(Digitlen(100, -1))
}
