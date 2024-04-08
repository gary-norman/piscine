package piscine

func factorial(a int) int {
	if a <= 1 {
		return 1
	}
	return a * factorial(a-1)
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
