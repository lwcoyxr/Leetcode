func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	result := n & (n - 1) // took that beautiful solution from CharryLee0426

	if result == 0 {
		return true
	}
	return false
}