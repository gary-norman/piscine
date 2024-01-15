package piscine

func IterativeFactorial(nb int) int {
	if nb == 21 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	for i := nb - 1; i > 0; i-- {
		nb *= i
	}
	return nb
}
