package challenge10

import (
	"euler/prime"
)

func PrimeSumUpto(target int64) int64 {
	var sum, i int64

	for i = 2; i < target; i++ {
		if prime.IsPrime(i) {
			sum += i
		}
	}
	return sum
}
