package piscine

func RecursivePower(nb int, power int) int {
	// duplicate nb as multiplier
	var result int = 1
	// return 0 if power is negative
	if power < 0 {
		return 0
	}
	// return 1 if exponent is 0
	// if power == 0 {
	// 	return 1
	// }
	if power != 0 {
		result = (nb * RecursivePower(nb, power-1))
	}
	return result
}
