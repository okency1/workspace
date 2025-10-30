package main

func CAMELtoSnackCase(s string) string {
	if s == "" {
		return ""
	}
	if s[0] < 'a' || s[0] > 'z' {
		return s
	}

	hasUpper := false
	for i := 1; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return s
	}

	converted := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			converted += "_"
			ch = ch + 32 // convert to lowercase
		}
		converted += string(ch)
	}
	return converted
}
