package challenge07

import (
	"euler/utils"
)

func NthPrime(target int) int {
	start := 2
	for i := 1; i < target; {
		start++
		if utils.IsPrime(int64(start)) {
			i++
		}
	}
	return start
}
