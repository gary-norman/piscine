package piscine

func IsPrime(nb int) bool {
	// return false if nb is 1, 0, or negative
	if nb < 2 {
		return false
	}
	// return false if nb is cleanly divisible by any number other than nb
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
