package challenge23

import "testing"

func TestSumOfNonAbundantNumbers(t *testing.T) {
	sum := SumOfNonAbundantNumbers()
	if sum != 4179871 {
		t.Error("Expect sum to be 4179871", sum)
	}
}
