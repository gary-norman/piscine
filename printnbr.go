package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	SetNbr(n)
}

func SetNbr(n int) {
	var i int
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	i = (n % 10)
	if n/10 != 0 {
		SetNbr(n / 10)
	}
	if i < 0 {
		i = i * -1
	}
	z01.PrintRune(rune(i) + 48)
}
