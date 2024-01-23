package piscine

func IsPrime(nb int) bool {
	// return false if nb is 1
	if nb == 1 {
		return false
	}
	// return true if nb is prime, else false
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
