package challenge20

import "testing"

func TestFactorialDigitSum(t *testing.T) {
	sum := FactorialDigitSum(100)
	if sum != 648 {
		t.Error("Expect sum to be 648", sum)
	}
}
