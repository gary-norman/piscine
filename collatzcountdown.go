package piscine

// If the number is even, divide it by two.
// If the number is odd, triple it and add one.

func CollatzCountdown(start int) int {
	counter := 0
	if start <= 0 {
		return -1
	}
	for start > 1 {
		if start%2 == 0 {
			start = start / 2
			counter++
		} else if start%2 == 1 {
			start = start*3 + 1
			counter++
		}
	}
	return counter
}
