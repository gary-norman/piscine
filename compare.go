package piscine

func Compare(a, b string) int {
	result := 0
	smallest := len(a)
	if len(b) < len(a) {
		smallest = len(b)
	}
	// runea := []rune(a)
	// runeb := []rune(b)
	for i := 0; i < smallest; i++ {
		// if runea[i] < 97 {
		// 	runea[i] += 32
		// }
		// if runeb[i] < 97 {
		// 	runeb[i] += 32
		// }
		if a[i] > b[i] {
			result = 1
			return result
		}
		if a[i] < b[i] {
			result = -1
			return result
		}
	}
	if len(a) > len(b) {
		result = 1
	} else if len(a) < len(b) {
		result = -1
	}
	return result
}
