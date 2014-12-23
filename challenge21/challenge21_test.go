package challenge21

import "testing"

func TestSumOfAmicableNumbers(t *testing.T) {
	sum := SumOfAmicableNumbers(10000)
	if sum != 31626 {
		t.Error("Expect sum to be 31626", sum)
	}
}

func TestSumOfAmicableNumbers2(t *testing.T) {
	sum := SumOfAmicableNumbers2(10000)
	if sum != 31626 {
		t.Error("Expect sum to be 31626", sum)
	}
}
