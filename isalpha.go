package piscine

func IsAlpha(s string) bool {
	var result bool
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 48 && s[i] <= 57) &&
			!(s[i] >= 65 && s[i] <= 90) &&
			!(s[i] >= 97 && s[i] <= 122) {
			result = false
			return result
		} else {
			result = true
		}
	}
	return result
}
