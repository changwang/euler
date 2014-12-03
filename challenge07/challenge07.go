package challenge07

import (
	"euler/prime"
)

func NthPrime(target int) int {
	start := 2
	for i := 1; i < target; {
		start++
		if prime.IsPrime(int64(start)) {
			i++
		}
	}
	return start
}
