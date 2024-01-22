package piscine

func lower(s string) string {
	runes := []rune(s)
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			runes[i] = rune(s[i] + 32)
			result += string(runes[i])
		} else {
			result += string(runes[i])
		}
	}
	return result
}

func Capitalize(s string) string {
	s = lower(s)
	runes := []rune(s)
	newWord := true
	var result string
	for i := 0; i < len(s); i++ {
		if newWord {
			if runes[i] >= 97 && runes[i] <= 122 {
				runes[i] = rune(s[i] - 32)
				result += string(runes[i])
				newWord = false
			} else {
				result += string(runes[i])
			}
			if runes[i] >= 48 && runes[i] <= 57 {
				newWord = false
			}
		} else if !((s[i] >= 97 && s[i] <= 122) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 48 && s[i] <= 57)) {
			result += string(runes[i])
			newWord = true
		} else {
			result += string(runes[i])
		}
	}
	return result
}

// func Capitalize(s string) string {
// 	runes := []rune(s)
// 	for i, chat := range runes {
// 		if chat < 97 || chat > 122 {
// 			for j := 0; j < 26; j++ {
// 				if i < len(runes)-1 {
// 					if runes[i+1] == rune(j+97) {
// 						runes[i+1] = rune(j + 65)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return string(runes)
// }
