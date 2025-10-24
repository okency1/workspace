package main

import "fmt"

func HalfSlice(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	half := (len(slice) + 1) / 2
	return slice[:half]
}

func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
}
