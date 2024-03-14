package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]]

		if value < prevValue {
			sum -= value
		} else {
			sum += value
		}
		prevValue = value
	}
	return sum
}

func main() {
	var input string
	fmt.Print("Enter a Roman numeral: ")
	fmt.Scanln(&input)

	result := romanToInt(input)
	fmt.Printf("The integer value is: %d\n", result)
}
