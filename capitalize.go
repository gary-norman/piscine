package piscine

func lower(s string) string {
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

func Capitalize(s string) string {
	s = lower(s)
	runes := []rune(s)
	newWord := false
	var result string
	for i := 0; i < len(s); i++ {
		if (i == 0 || newWord) && (runes[i] >= 97 && runes[i] <= 122) {
			runes[i] = rune(s[i] - 32)
			result += string(runes[i])
			newWord = false
		} else if !(s[i] >= 97 && s[i] <= 122 || s[i] >= 65 && s[i] <= 90 || s[i] >= 48 && s[i] <= 57) {
			result += string(runes[i])
			newWord = true
		} else {
			result += string(runes[i])
		}
	}
	return result
}
