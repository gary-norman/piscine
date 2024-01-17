package piscine

func AlphaCount(s string) int {
	counter := 0
	for index, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			counter++
			index += index
		}
	}
	return counter
}
