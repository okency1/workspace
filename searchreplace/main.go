package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	s := os.Args[1]
	oldi := os.Args[2]
	newi := os.Args[3]
	result := strings.ReplaceAll(s, oldi, newi)
	os.Stdout.WriteString(result + "\n")
}
