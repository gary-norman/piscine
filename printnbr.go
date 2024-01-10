package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var s []int
	for n > 0 {
		i := (n % 10) + 48
		s = append(s, i)
		n = (n / 10) - ((n % 10) / 10)
	}
	for j := len(s) - 1; j >= 0; j-- {
		z01.PrintRune(rune(s[j]))
	}
}
