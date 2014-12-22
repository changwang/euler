package challenge37

import (
	"euler/utils"
	"math"
)

func GetAllTruncatablePrimes() int64 {
	var sum int64
	var i int64
	found := 0
	for i = 10; ; i++ {
		if IsTruncatablePrime(i) {
			found++
			sum += i
		}
		if found == 11 {
			break
		}
	}

	return sum
}

func IsTruncatablePrime(num int64) bool {
	if num < 10 {
		return false
	}

	if !utils.IsPrime(num) {
		return false
	}

	shiftLeft := num
	shiftRight := num
	for shiftLeft > 10 && shiftRight > 10 {
		shiftLeft = shiftToLeft(shiftLeft)
		shiftRight = shiftToRight(shiftRight)
		if !utils.IsPrime(shiftLeft) || !utils.IsPrime(shiftRight) {
			return false
		}
	}

	return true
}

func shiftToRight(num int64) int64 {
	return num / 10
}

func shiftToLeft(num int64) int64 {
	length := utils.DigitLength(num)
	base := int64(math.Pow(10, float64(length-1)))
	return num - (num/base)*base
}
