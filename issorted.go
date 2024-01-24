package piscine

func IsSorted(f func(a, b int) int, slice []int) bool {
	result := 0
	for i := len(slice) - 1; i > 0; i-- {
		// result = 0
		result = f(slice[i], slice[i-1])
		if result < 0 {
			return false
		} else {
			continue
		}
	}
	return true
}
