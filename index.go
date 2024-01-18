package piscine

// func Index(s string, toFind string) int {
// 	result := -1
// 	contains := 0
// 	for index, char := range s {
// 		for i := 0; i < len(toFind); i++ {
// 			if rune(toFind[i]) == char {
// 				result = index
// 				contains++

// 			}
// 		}
// 	}
// 	return result
// }

func Index(s string, toFind string) int {
	// if toFind is an empty string
	if toFind == "" || s == "" {
		return 0
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind { // slices the string to check for toFind inside it
			return i
		}
	}
	return -1
}
