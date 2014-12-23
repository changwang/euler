package challenge21

import (
	"euler/utils"
)

func SumOfAmicableNumbers(target int64) int64 {
	var sum int64

	dMap := divisorMap(target)
	for k, v := range dMap {
		pair, ok := dMap[v]
		if ok && k == pair && k != v {
			sum += k
		}
	}
	return sum
}

func SumOfAmicableNumbers2(target int64) int64 {
	var i int64
	amicableNumbers := make([]int64, 0)

	for i = 2; i <= target; i++ {
		divisorsLeft := utils.ProperDivisors(i)
		sumOfLeft := utils.SliceSum(divisorsLeft)

		if i == sumOfLeft {
			continue
		}

		divisorsRight := utils.ProperDivisors(sumOfLeft)
		sumOfRight := utils.SliceSum(divisorsRight)

		if i == sumOfRight {
			if !utils.Contains(amicableNumbers, i) {
				amicableNumbers = append(amicableNumbers, i)
			}
		}
	}
	return utils.SliceSum(amicableNumbers)
}

func divisorMap(target int64) map[int64]int64 {
	divisorMap := make(map[int64]int64, target)

	var i int64
	for i = 1; i <= target; i++ {
		divisors := utils.ProperDivisors(i)
		sum := utils.SliceSum(divisors)
		divisorMap[i] = sum
	}
	return divisorMap
}
