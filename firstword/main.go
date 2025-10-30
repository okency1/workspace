package piscine

func FirstWord(s string) string {
	for i := range s {
		if s[i] == ' ' || s[i] == '\t' {
			return s[:i] + "\n"
		}
	}
	return s + "\n"
}
