package piscine

func Rot14(s string) string {
	sNew := []rune(s)
	for index, n := range sNew {
		if n >= 65 && n <= 77 || n >= 97 && n <= 109 {
			sNew[index] = n + 14
		}
		if n >= 77 && n <= 90 || n >= 109 && n <= 122 {
			sNew[index] = n - 12
		}
	}
	return string(sNew)
}
