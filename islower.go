package piscine

func IsLower(s string) bool {
	var result bool
	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			result = false
			return result
		} else {
			result = true
		}
	}
	return result
}
