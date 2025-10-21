package main

import (
	"fmt"
)

func CheckNumber(s string) bool {
	for _, i := range s {
		if i >= '0' && i <= '9' {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(CheckNumber("2dfghj"))
	fmt.Println(CheckNumber("ghj"))
	fmt.Println(CheckNumber("2344"))
}
