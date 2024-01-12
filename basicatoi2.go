package piscine

func BasicAtoi2(s string) int {
	var b int

	for i := 0; i <= len(s)-1; i++ {
		if s[i] == ' ' {
			b = 0
			break
		}
		if s[i] >= '0' && s[i] <= '9' {
			b += int(s[i] - 48)
			if i <= len(s)-2 {
				b *= 10
			}
		} else {
			b = 0
			break
		}
	}
	return b
}
