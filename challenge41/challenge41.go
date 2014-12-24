package challenge41

import "euler/utils"

func GetMaxPandigitalPrime() int64 {
	/*
		sum of 9 pandigits is 45, sum of 8 pandigits 36 are both divisible by 3.
		then starting from 7 should be fine.
	*/
	var max int64
	digits := utils.CreatePandigitals(7)
	for _, d := range digits {
		if utils.IsPrime(d) {
			max = d
		}
	}
	return max
}
