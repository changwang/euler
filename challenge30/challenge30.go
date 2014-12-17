package challenge30

import (
	"math"
)

func SumOfNthPowers(exp int64) float64 {
	var sum, i float64
	start := float64(2)
	end := math.Pow(10, float64(exp+2))
	for i = start; i < end; i++ {
		result := powerOfNth(int64(i), exp)
		if result == i {
			sum += result
		}
	}
	return sum
}

func powerOfNth(num, exp int64) float64 {
	var pow float64

	for num > 0 {
		d := num % 10
		pow += math.Pow(float64(d), float64(exp))
		num /= 10
	}
	return pow
}
