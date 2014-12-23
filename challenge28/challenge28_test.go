package challenge28

import "testing"

func TestSpiralDiagonalSum(t *testing.T) {
	sum := SpiralDiagonalSum(1001)
	if sum != 669171001 {
		t.Error("Expect sum to be 669171001", sum)
	}
}
