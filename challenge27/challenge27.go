package challenge27

import (
	"euler/utils"
)

func QuadraticCoefficient(target int) float64 {
	var maxNOP int64
	var productOfCoeff float64
	primes := make([]int64, 0)
	primes = append(primes, 1)

	for i := 0; i < target; i++ {
		if utils.IsPrime(int64(i)) {
			primes = append(primes, int64(i))
		}
	}

	for _, p1 := range primes {
		for _, p2 := range primes {
			nop1 := numberOfPrimes(p1, p2)
			nop2 := numberOfPrimes(-p1, p2)
			nop3 := numberOfPrimes(p1, -p2)

			if maxNOP < nop1 {
				maxNOP = nop1
				productOfCoeff = float64(p1 * p2)
			}

			if maxNOP < nop2 {
				maxNOP = nop2
				productOfCoeff = float64((-p1) * p2)
			}
			if maxNOP < nop1 {
				maxNOP = nop3
				productOfCoeff = float64(p1 * (-p2))
			}
		}
	}

	return productOfCoeff
}

func numberOfPrimes(coefficient1, coefficient2 int64) int64 {
	var num, i int64

	for i = 0; ; i++ {
		quad := (i * i) + coefficient1*i + coefficient2
		if utils.IsPrime(quad) {
			num++
		} else {
			break
		}
	}

	return num
}
