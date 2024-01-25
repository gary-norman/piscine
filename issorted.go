package piscine

func F(a, b int) int {
	return a - b
}

func IsSorted(F func(a, b int) int, slice []int) bool {
	counterFwd := 0
	counterBkwd := 0
	for i := 0; i < len(slice)-1; i++ {
		if F(slice[i+1], slice[i]) >= 0 {
			counterFwd++
		}
		if F(slice[i], slice[i+1]) >= 0 {
			counterBkwd++
		}
	}
	slices := len(slice) - 1
	if counterFwd == slices || counterBkwd == slices {
		return false
	} else {
		return true
	}
}
