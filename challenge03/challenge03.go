package challenge03

import (
	"math"
)

func LargestPrime(target int64) int64 {
	var lastPrime int64
	var iterator int64
	iterator = 2

	target = removeFactor2(target)

	ceil := math.Ceil(math.Sqrt(float64(target)))

	for ; iterator < int64(ceil); iterator++ {
		if target%iterator == 0 {
			lastPrime = iterator
			target /= iterator
		}
	}

	return lastPrime
}

func removeFactor2(target int64) int64 {
	for target%2 == 0 {
		target /= 2
	}
	return target
}
