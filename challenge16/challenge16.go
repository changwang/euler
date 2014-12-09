package challenge16

import (
	"math"
	"strconv"
)

func PowerDigitSum(power int64) int64 {
	var result int64
	pow := float64(math.Pow(2, float64(power)))
	strRep := strconv.FormatFloat(pow, 'f', 0, 64)

	for _, ch := range strRep {
		i, _ := strconv.Atoi(string(ch))
		result += int64(i)
	}
	return result
}
