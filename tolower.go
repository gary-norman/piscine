package piscine

func ToLower(s string) string {
	runes := []rune(s)
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			runes[i] = rune(s[i] + 32)
			result += string(runes[i])
		} else {
			result += string(runes[i])
		}
	}
	return result
}
