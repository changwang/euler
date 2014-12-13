package utils

import (
	"math"
)

func IsPrime(num int64) bool {
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
