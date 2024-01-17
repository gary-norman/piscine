package piscine

func Index(s string, toFind string) int {
	var result int
	for index, char := range s {
		for i := 0; i < len(toFind); i++ {
			if rune(toFind[i]) == char {
				if i > 0 {
					result = index - i
				} else {
					result = index
					return result
				}
			}
		}
	}
	return result
}
