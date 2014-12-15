package challenge23

import (
	"euler/utils"
)

func SumOfNonAbundantNumbers() float64 {
	var sum float64
	var i int32
	abundant := make([]int32, 0)

	for i = 1; i < 28125; i++ {
		divisors := utils.ProperDivisors(i)
		result := utils.SliceSum(divisors)
		if i < result && !utils.Contains(abundant, i) {
			abundant = append(abundant, i)
		}

		if !isSumOfAbundantNumbers(i, abundant) {
			sum += float64(i)
		}
	}
	return sum
}

func isSumOfAbundantNumbers(number int32, abundants []int32) bool {
	var i int32

	for i = 1; i <= number/2; i++ {
		if utils.Contains(abundants, i) && utils.Contains(abundants, (number-i)) {
			return true
		}
	}
	return false
}
