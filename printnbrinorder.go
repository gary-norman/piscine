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
	for i := 0; i < sliceLen; i++ {
		for j := 0; j < sliceLen-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	for _, digit := range slice {
		z01.PrintRune(rune(digit + 48))
	}
}
