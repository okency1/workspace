package main

import "fmt"

func PrintIf(s string) string {
	if s == "" {
		return "G\n"
	}
	if len(s) >= 3 {
		return "G\n"
	}
	return "Invalid input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
