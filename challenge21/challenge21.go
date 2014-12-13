package challenge21

import (
	"euler/utils"
	"math"
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
		divisorsLeft := properDivisors(i)
		sumOfLeft := utils.SliceSum(divisorsLeft)

		if i == sumOfLeft {
			continue
		}

		divisorsRight := properDivisors(sumOfLeft)
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
		divisors := properDivisors(i)
		sum := utils.SliceSum(divisors)
		divisorMap[i] = sum
	}
	return divisorMap
}

func properDivisors(num int32) []int32 {
	if num == 1 {
		return []int32{0}
	}

	stopper := int32(math.Ceil(math.Sqrt(float64(num))))
	var i int32
	divisors := make([]int32, 0)
	divisors = append(divisors, 1)

	for i = 2; i <= stopper; i++ {
		if num%i == 0 {
			if !utils.Contains(divisors, i) && i < num {
				divisors = append(divisors, i)
			}
			if !utils.Contains(divisors, num/i) && num/i < num {
				divisors = append(divisors, num/i)
			}
		}
	}
	return divisors
}
