func smallestEvenMultiple(n int) int {
	if n%2 != 0 {
		n += n
	} else {
		return n
	}
	return n
}