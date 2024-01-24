package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
		return
	}
	var slice []int
	for n > 0 {
		slice = append(slice, n%10)
		n /= 10
	}
	sliceLen := len(slice)
	for i := 0; i < sliceLen-1-1; i++ {
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
		}
	}
	for _, digit := range slice {
		z01.PrintRune(rune(digit + 48))
	}
}
