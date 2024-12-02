package utils

import "strconv"

func DigitToInt(digit string) int {
	number, err := strconv.Atoi(digit)
	if err != nil {
		panic("We have not a number, we can't proceed")
	}
	return number
}