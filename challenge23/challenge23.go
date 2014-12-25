package challenge23

import (
	"euler/utils"
)

func SumOfNonAbundantNumbers() float64 {
	var sum float64
	var i int64
	var abundants []int64

	for i = 1; i < 28125; i++ {
		divisors := utils.ProperDivisors(i)
		result := utils.SliceSum(divisors)
		if i < result && !utils.Contains(abundants, i) {
			abundants = append(abundants, i)
		}

		if !isSumOfAbundantNumbers(i, abundants) {
			sum += float64(i)
		}
	}
	return sum
}

func isSumOfAbundantNumbers(number int64, abundants []int64) bool {
	var i int64

	for i = 1; i <= number/2; i++ {
		if utils.Contains(abundants, i) && utils.Contains(abundants, (number-i)) {
			return true
		}
	}
	return false
}
