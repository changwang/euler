package challenge30

import "testing"

func TestSumOfNthPowers(t *testing.T) {
	sum := SumOfNthPowers(5)
	if sum != 443839 {
		t.Error("Expect sum to be 443839", sum)
	}
}
