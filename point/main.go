package main

import (
	"piscine"

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

func main() {
	points := &point{}
	setPoint(points)
	piscine.PrintStr("x = ")
	piscine.PrintNbr(points.x)
	piscine.PrintStr(", y = ")
	piscine.PrintNbr(points.y)
	z01.PrintRune('\n')

	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
