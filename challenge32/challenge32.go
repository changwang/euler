package challenge32

import (
	"euler/utils"
	"fmt"
	"strconv"
)

func FindAllPandigitalProducts() int64 {
	found := make(map[int64]bool, 0)
	var sum, i, j int64

	for i = 1; i < 10000; i++ {
		for j = 1; j < 1000; j++ {
			combination := fmt.Sprintf("%d%d%d", i, j, i*j)
			digit, _ := strconv.ParseInt(combination, 10, 64)
			if !found[i*j] && utils.IsPandigital(digit, 9) {
				found[i*j] = true
			}
		}
	}

	for k := range found {
		sum += k
	}
	return sum
}
