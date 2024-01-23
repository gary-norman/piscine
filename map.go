package piscine

func Map(f func(int) bool, a []int) []bool {
	aLen := len(a)
	result := make([]bool, aLen)
	for i := 0; i < aLen; i++ {
		result[i] = f(a[i])
	}
	return result
}
