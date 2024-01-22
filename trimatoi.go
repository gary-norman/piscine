package piscine

func TrimAtoi(s string) int {
	num := 0
	negative := false
	counter := 0
	multiplier := 0
	for _, char := range s {
		if !negative && counter == 0 {
			if char == 45 {
				negative = true
			}
		}
		if char > 48 && char <= 57 {
			if counter == 0 {
				num += (int(char - 48))
				multiplier = 10
			} else {
				num *= multiplier
				num += (int(char - 48))
			}
			counter++
		}
	}
	if negative {
		return num * -1
	} else {
		return num
	}
}
