package piscine

func StrRev(s string) string {
	a := []rune(s)
	var b string
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	return b
}
