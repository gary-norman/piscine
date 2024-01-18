package piscine

func IsPrintable(s string) bool {
	var result bool
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 31 && s[i] <= 127) {
			result = false
			return result
		} else {
			result = true
		}
	}
	return result
}
