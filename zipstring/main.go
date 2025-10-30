package piscine

import "strconv"

func ZipString(s string) string {
	r, i := "", 0
	for i < len(s) {
		j := i + 1
		for j < len(s) && s[j] == s[i] {
			j++
		}
		r += strconv.Itoa(j-i) + string(s[i])
		i = j
	}
	return r
}
