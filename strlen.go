package piscine

func StrLen(s string) int {
	var b int
	for i := 0; i < len(s); i++ {
		b++
	}
	return b
}