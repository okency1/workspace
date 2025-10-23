package main

import "fmt"

func PrintIfNot(s string) string {
	if s == "" {
		return "G\n"
	}
	if len(s) < 3 {
		return "G\n"
	}
	return "Invalid input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
