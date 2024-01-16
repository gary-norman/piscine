package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if n < 0 || n > len(s) {
		return 0
	}
	return a[n-1]
}
