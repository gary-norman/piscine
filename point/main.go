package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

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

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(", y = ")
	PrintNbr(points.y)
	z01.PrintRune('\n')
	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
