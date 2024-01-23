package piscine

func Sqrt(nb int) int {
	root := 1
	for root*root <= nb { // loop is going as long as rootroot is less than nb
		if root*root == nb { // when reached, return root
			return root
		}
		root++ // otherwise keep looping
	}
	return 0 // indicates that the square root is not an int (float64)
}
