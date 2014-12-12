package challenge20

import (
	"math/big"
	"strconv"
)

func FactorialDigitSum(target int64) int64 {
	var num *big.Int
	num = big.NewInt(1)
	var sum, i int64

	for i = 1; i <= target; i++ {
		num = num.Mul(num, big.NewInt(int64(i)))
	}

	digitString := num.String()
	for _, ds := range digitString {
		d, _ := strconv.Atoi(string(ds))
		sum += int64(d)
	}
	return sum
}

/*
 think about, 2, 5, 10 only contribute 0 to the end of final result
*/
