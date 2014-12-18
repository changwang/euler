package challenge32

import (
	"fmt"
	"strings"
)

var occurred = map[string]bool{
	"1": false,
	"2": false,
	"3": false,
	"4": false,
	"5": false,
	"6": false,
	"7": false,
	"8": false,
	"9": false,
}

const (
	zero = "0"
)

func FindAllPandigitalProducts() int64 {
	found := make(map[int64]bool, 0)
	var sum, i, j int64

	for i = 1; i < 10000; i++ {
		for j = 1; j < 1000; j++ {
			if !found[i*j] && IsPandigitalProduct(i, j) {
				found[i*j] = true
			}
		}
	}

	for k, _ := range found {
		sum += k
	}
	return sum
}

func IsPandigitalProduct(multiplier, multiplicant int64) bool {
	combination := fmt.Sprintf("%d%d%d", multiplier, multiplicant, multiplier*multiplicant)
	if len(combination) != 9 || strings.Contains(combination, zero) {
		return false
	}

	defer resetCheckMap()
	for _, d := range combination {
		b, ok := occurred[string(d)]
		if !ok || b {
			return false
		} else {
			occurred[string(d)] = true
		}
	}

	for _, v := range occurred {
		if !v {
			return false
		}
	}
	return true
}

func resetCheckMap() {
	for k, _ := range occurred {
		occurred[k] = false
	}
}
