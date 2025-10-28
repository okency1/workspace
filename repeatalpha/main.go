package main

import "fmt"

func RepeatAlpha(s string) string {
	if s == "" {
		return ""
	}
	var result []rune
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			var position int
			if i >= 'a' && i <= 'z' {
				position = int(i - 'a' + 1)
			} else {
				position = int(i - 'A' + 1)
			}
			for r := 0; r < position; r++ {
				result = append(result, i)
			}
		} else {
			result = append(result, i)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
