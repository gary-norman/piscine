package piscine

func IsUpper(s string) bool {
	var result bool
	for i := 0; i < len(s); i++ {
		if s[i] < 65 || s[i] > 90 {
			result = false
			return result
		} else {
			result = true
		}
	}
	return result
}
