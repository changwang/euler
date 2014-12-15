package challenge21

import (
	"euler/utils"
)

func SumOfAmicableNumbers(target int32) int32 {
	var sum int32

	dMap := divisorMap(target)
	for k, v := range dMap {
		pair, ok := dMap[v]
		if ok && k == pair && k != v {
			sum += k
		}
	}
	return sum
}

func SumOfAmicableNumbers2(target int32) int32 {
	var i int32
	amicableNumbers := make([]int32, 0)

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

func divisorMap(target int32) map[int32]int32 {
	divisorMap := make(map[int32]int32, target)

	var i int32
	for i = 1; i <= target; i++ {
		divisors := utils.ProperDivisors(i)
		sum := utils.SliceSum(divisors)
		divisorMap[i] = sum
	}
	return divisorMap
}
