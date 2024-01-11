package piscine

func StrLen(s string) int {
	a := []rune(s)
	var b int
	for i := 0; i < len(a); i++ {
		b++
	}
	return b
}
