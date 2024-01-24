package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 1)
	wordcount := 0
	newword := false
	for _, char := range s {
		if string(char) != " " {
			result[wordcount] += string(char)
			newword = false
		} else {
			if (string(char) == " " && !newword) || string(char) == "	" || string(char) == "\n" {
				result = append(result, "")
				wordcount++
				newword = true
			}
		}
	}

	if result[wordcount-1] == "" {
		result = result[:len(result)-1]
	}
	return result
}
