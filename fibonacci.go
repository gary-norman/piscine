package piscine

func Fibonacci(index int) int {
	// return -1 if index is negative
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}
