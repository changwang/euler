package utils

import (
	"math"
)

var occurred = map[int64]bool{
	1: false,
	2: false,
	3: false,
	4: false,
	5: false,
	6: false,
	7: false,
	8: false,
	9: false,
}

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

func Contains(intSlice []int32, num int32) bool {
	for _, d := range intSlice {
		if num == d {
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

func IsPalindrome(input string) bool {
	length := len(input)
	if length == 1 {
		return true
	}

	var start, end int
	if length%2 != 0 {
		start = length/2 - 1
		end = length/2 + 1
	} else {
		start, end = length/2-1, length/2
	}

	for start >= 0 && end < length {
		if input[start] != input[end] {
			return false
		}
		start--
		end++
	}

	return true
}

func DigitLength(num int64) int {
	var length int

	if num == 0 {
		return 1
	}

	if num < 0 {
		num -= num
	}

	for num > 0 {
		num /= 10
		length++
	}
	return length
}

func IsPandigital(num, span int64) bool {
	length := int64(DigitLength(num))
	defer resetCheckMap()
	if length != span {
		return false
	}

	for num > 0 {
		rem := num % 10
		if rem > span {
			return false
		}

		num /= 10
		found, ok := occurred[rem]
		if !ok || found {
			return false
		}

		occurred[rem] = true
	}
	return true
}

func resetCheckMap() {
	for k, _ := range occurred {
		occurred[k] = false
	}
}
