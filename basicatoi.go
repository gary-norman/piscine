package piscine

func BasicAtoi(s string) int {
	var b int
	for i := 0; i <= len(s)-1; i++ {
		b += int(s[i] - 48)
		if i <= len(s)-2 {
			b *= 10
		}
	}
	return b
}
