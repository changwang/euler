package challenge16

import "testing"

func TestPowerDigitSum(t *testing.T) {
	sum := PowerDigitSum(1000)
	if sum != 1366 {
		t.Error("Expect sum to be 1366", sum)
	}
}
