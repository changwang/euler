package challenge07

import (
	"math"
)

func NthPrime(target int) int {
	start := 2
	for i := 1; i < target; {
		start++
		if isPrime(start) {
			i++
		}
	}
	return start
}

func isPrime(num int) bool {
	var sqrt int
	sqrt = int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
