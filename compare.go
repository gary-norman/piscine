package piscine

func Compare(a, b string) int {
	result := 0
	runea := []rune(a)
	runeb := []rune(b)
	for i := 0; i < min(len(a), len(b)); i++ {
		if runea[i] < 97 {
			runea[i] += 32
		}
		if runeb[i] < 97 {
			runeb[i] += 32
		}
		if runea[i] > runeb[i] {
			result = -1
			return result
		}
		if runea[i] < runeb[i] {
			result = 1
			return result
		}
	}
	if a > b {
		result = 1
	} else if b < a {
		result = -1
	}
	return result

}
