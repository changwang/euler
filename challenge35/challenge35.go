package challenge35

import (
	"euler/utils"
	"math"
)

func GetAllCircularPrimes() []int64 {
	circularsPrimes := make([]int64, 0)
	var i int64

	for i = 2; i < 1000000; i++ {
		if !utils.IsPrime(i) {
			continue
		}

		rotations := circularRotations(i)
		allPrime := false
		for _, r := range rotations {
			if !utils.IsPrime(int64(r)) {
				allPrime = false
				break
			}
			allPrime = true
		}

		if allPrime {
			circularsPrimes = append(circularsPrimes, i)
		}
	}
	return circularsPrimes
}

func circularRotations(num int64) []int64 {
	circulars := make([]int64, 0)
	circulars = append(circulars, num)

	if num < 10 {
		return circulars
	}

	numLen := utils.DigitLength(int64(num)) - 1

	factor := int64(math.Pow(10, float64(numLen)))

	for ; numLen > 0; numLen-- {
		rem := num % 10
		num = num/10 + rem*factor
		if !utils.Contains(circulars, num) {
			circulars = append(circulars, num)
		}
	}

	return circulars
}
