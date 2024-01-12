package piscine

func Atoi(s string) int {
	var b int
	isNeg := false
	for i := 0; i <= len(s)-1; i++ {
		// check for double pos or neg
		if (s[i] == '+' && s[i+1] == '+') || (s[i] == '-' && s[i+1] == '-') {
			b = 0
			break
		}
		// check for space
		if s[i] == ' ' {
			b = 0
			break
		}

		// check for digits
		if (s[i] >= '0' && s[i] <= '9') || s[i] == '+' || s[i] == '-' {
			if s[i] == '+' {
			} else if s[i] == '-' {
				isNeg = true
			} else {
				b += int(s[i] - 48)
				if i <= len(s)-2 {
					b *= 10
				}
			}
		} else {
			b = 0
			break
		}
	}
	if isNeg {
		return -(b)
	} else {
		return b
	}
}
