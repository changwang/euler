package utils

import (
	"math"
)

func IsPrime(num int64) bool {
	if num < 2 {
		return false
	}

	var sqrt, i int64
	sqrt = int64(math.Sqrt(float64(num)))
	for i = 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Contains(divisors []int32, div int32) bool {
	for _, d := range divisors {
		if div == d {
			return true
		}
	}
	return false
}

func SliceSum(intSlice []int32) int32 {
	var sum int32

	for _, d := range intSlice {
		sum += d
	}
	return sum
}

func ProperDivisors(num int32) []int32 {
	if num == 1 {
		return []int32{0}
	}

	stopper := int32(math.Ceil(math.Sqrt(float64(num))))
	var i int32
	divisors := make([]int32, 0)
	divisors = append(divisors, 1)

	for i = 2; i <= stopper; i++ {
		if num%i == 0 {
			if !Contains(divisors, i) && i < num {
				divisors = append(divisors, i)
			}
			if !Contains(divisors, num/i) && num/i < num {
				divisors = append(divisors, num/i)
			}
		}
	}
	return divisors
}

func Factorial(num int64) float64 {
	if num == 0 || num == 1 {
		return 1
	}

	var factorial float64
	factorial = 1
	for num > 1 {
		factorial *= float64(num)
		num--
	}
	return factorial
}
