package piscine

func Swap(a *int, b *int) {
	var c int
	var d int
	c = *a
	d = *b
	*a = d
	*b = c
}
