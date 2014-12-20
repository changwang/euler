package challenge35

import (
	"euler/utils"
	"fmt"
	"math"
)

func GetAllCircularPrimes() []int32 {
	circularsPrimes := make([]int32, 0)
	var i int32

	for i = 2; i < 1000000; i++ {
		if !utils.IsPrime(int64(i)) {
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

func circularRotations(num int32) []int32 {
	circulars := make([]int32, 0)
	circulars = append(circulars, num)

	if num < 10 {
		return circulars
	}

	strNum := fmt.Sprintf("%d", num)
	numLen := len(strNum) - 1

	factor := int32(math.Pow(10, float64(numLen)))

	for ; numLen > 0; numLen-- {
		rem := num % 10
		num = num/10 + rem*factor
		if !utils.Contains(circulars, num) {
			circulars = append(circulars, num)
		}
	}

	return circulars
}
