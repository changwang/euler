package challenge25

import (
	"math/big"
)

func NLengthFibonacci(length int64) int64 {
	var f1, f2 *big.Int
	var term int64
	f1, f2 = big.NewInt(1), big.NewInt(1)
	term = 2

	for {
		f1, f2 = f2, f1.Add(f1, f2)
		term++

		str := f2.String()
		if int64(len(str)) == length {
			return term
		}
	}
}
