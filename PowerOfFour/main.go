func powerOfFour(n int) bool {
	return n > 0 && (n&(n-1) == 0) && (n&0x55555555 != 0) //same power check for pow of two but with additional mask check for making sure that every 1st, 5th, 9th one position of pow 4 is as expected
}
