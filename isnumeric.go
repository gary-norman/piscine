package piscine

func IsNumeric(s string) bool {
	var result bool
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 48 && s[i] <= 57) {
			result = false
			return result
		} else {
			result = true
		}
	}
	return result
}
