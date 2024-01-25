package piscine

func IsSorted(f func(a, b int) int, slice []int) bool {
	counterFwd := 0
	counterBkwd := 0
	for i := 1; i < len(slice)-1; i++ {
		if f(slice[i], slice[i]-1) >= 0 {
			counterFwd++
		}
	}
	for i := len(slice) - 1; i > 0; i-- {
		if f(slice[i], slice[i-1]) >= 0 {
			counterBkwd++
		}
	}
	slices := len(slice) - 1
	if counterFwd == slices || counterBkwd == slices {
		return true
	} else {
		return false
	}
}
