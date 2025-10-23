package main

func HashCode(dec string) string {
	size := len(dec)
	result := make([]byte, size)

	for i := 0; i < size; i++ {
		hashChar := (int(dec[i]) + size) % 127
		if hashChar < 33 {
			hashChar += 33
		}

		result[i] = byte(hashChar)
	}

	return string(result)
}
