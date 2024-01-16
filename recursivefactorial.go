package piscine

func factorial(a int) int {
	b := a
	if a > 1 {
		b = b * factorial(a-1)
	}
	return b
}

func RecursiveFactorial(nb int) int {
	if nb >= 21 || nb <= -21 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	result := factorial(nb)
	return result
}
