package piscine

func Capitalize(s string) string {
	s = ToLower(s)
	runes := []rune(s)
	println(runes)
	newWord := false
	var result string
	for i := 0; i < len(s); i++ {
		if i == 0 || newWord {
			runes[i] = rune(s[i] - 32)
			result += string(runes[i])
			newWord = false
		} else if !(s[i] >= 33 && s[i] <= 127) {
			result += string(runes[i])
			newWord = true
		} else {
			result += string(runes[i])
		}
	}
	return result
}
