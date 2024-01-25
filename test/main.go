package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{-3, -2, 0, 1}

	result1 := piscine.IsSorted(F, a1)
	result2 := piscine.IsSorted(F, a2)
	result3 := piscine.IsSorted(F, a3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func F(a, b int) int {
	return a - b
}
