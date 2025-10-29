package main

import (
	"fmt"
	"unicode"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	var result []rune
	for i, r := range runes {
		if !unicode.IsLetter(r) {
			return s
		}
		if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(runes[i-1]) && i > 1 {
			return s
		}
		if unicode.IsUpper(runes[len(runes)-1]) {
			return s
		}
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
