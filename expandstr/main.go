package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		words := strings.Fields(os.Args[1])
		if len(words) > 0 {
			os.Stdout.WriteString(strings.Join(words, "   ") + "\n")
		}
	}
}
