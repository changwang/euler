package prime

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
