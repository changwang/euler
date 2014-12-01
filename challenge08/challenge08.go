package challenge08

import (
	"strconv"
)

/*
  I don't like my solution, some calculation should be skipped, if
  newly coming number is smaller than poping out number.
*/
func AdjacentProduct(adjacents int, digits string) int {
	maxProduct := 0

	head := 0
	tail := adjacents
	for tail < len(digits) {
		head++
		tail++
		product := recalculateInterval(digits, head, tail)
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct
}

func recalculateInterval(digits string, start, end int) int {
	product := 1
	for i := start; i < end; i++ {
		digit, _ := strconv.Atoi(string(digits[i]))
		product *= digit
	}
	return product
}
