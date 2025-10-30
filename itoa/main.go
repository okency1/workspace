package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	if n < 0 {
		s, n = "-", -n
	}
	r := ""
	for n > 0 {
		r = string('0'+n%10) + r
		n /= 10
	}
	return s + r
}
