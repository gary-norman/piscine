package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{902620, -166755, 234840, 721470, 817932, -731158, -186943, 825770}

	result1 := piscine.IsSorted(F, a1)
	result2 := piscine.IsSorted(F, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func F(a, b int) int {
	return a - b
}
