package piscine

func MakeRange(min, max int) []int {
	// var result []int
	if min >= max {
		return nil
	}
	// loop through and declare values
	result := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		result[i] = min + i
	}
	return result
}
