package piscine

func IterativePower(nb int, power int) int {
	// duplicate nb as multiplier
	multiplier := nb
	// return 0 if power is negative
	if power < 0 {
		return 0
	}
	//iterate through the powers
	for i := 1; i < power; i++ {
		// perform the calculation
		nb *= multiplier
	}
	// return the solution (nb)
	return nb
}
