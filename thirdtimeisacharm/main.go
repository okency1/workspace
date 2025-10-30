package piscine

func ThirdTimeIsACharm(s string) string {
	if len(s) < 3 {
		return "\n"
	}
	r := ""
	for i := 2; i < len(s); i += 3 {
		r += string(s[i])
	}
	return r + "\n"
}
