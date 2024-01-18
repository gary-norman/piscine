package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := []rune(os.Args[0])
	var lastDash int

	for index, char := range name {
		if char == '/' || char == '\\' {
			lastDash = index
		}
	}

	for i := lastDash + 1; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}
	z01.PrintRune('\n')
}
