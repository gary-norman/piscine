package piscine

func Abort(a, b, c, d, e int) int {
	seq := []int{a, b, c, d, e}
	for i := 0; i < len(seq)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(seq); j++ {
			if seq[j] < seq[minIndex] {
				minIndex = j
			}
		}
		seq[i], seq[minIndex] = seq[minIndex], seq[i]
	}
	middle := len(seq) / 2
	return seq[middle]
}
