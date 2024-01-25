package piscine

func F(a, b int) int {
	return a - b
}

func IsSorted(F func(a, b int) int, slice []int) bool {
	counterFwd := 0
	counterBkwd := 0
	for i := 0; i < len(slice)-1; i++ {
		if F(slice[i], slice[i+1]) <= 0 {
			counterFwd++
			// fmt.Printf("f) i[%v]:%v - \n", i, counterFwd)
		}
		if F(slice[i], slice[i+1]) >= 0 {
			counterBkwd++
			// fmt.Printf("b) i[%v]:%v - \n", i, counterBkwd)
		}
	}
	// print("\nafter for loop\n")
	slices := len(slice) - 1
	if counterFwd == slices || counterBkwd == slices {
		return true
	} else {
		return false
	}
}
