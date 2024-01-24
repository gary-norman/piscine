package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 1)
	wordcount := 0
	for _, char := range s {
		if string(char) != " " {
			result[wordcount] += string(char)
		} else {
			if string(char) == " " && string(char-1) != " " || string(char) == "	" || string(char) == "\n" {
				result = append(result, "")
				wordcount++
			}
		}
	}

	if result[wordcount-1] == "" {
		result = result[:len(result)-1]
	}
	return result
}
