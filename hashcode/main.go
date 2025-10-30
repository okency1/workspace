package piscine

func HashCode(dec string) string {
	result := ""
	for _, i := range dec {
		hash := (int(i) + len(dec)) % 127
		if hash < 33 {
			hash += 33
		}
		result += string(rune(hash))
	}
	return result
}
