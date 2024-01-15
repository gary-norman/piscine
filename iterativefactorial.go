package piscine

func IterativeFactorial(nb int) int {
	for i := nb - 1; i > 0; i-- {
		nb *= i
	}
	return nb
}
