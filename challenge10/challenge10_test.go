package challenge10

import "testing"

func TestPrimeSumUpto(t *testing.T) {
	sum := PrimeSumUpto(2000000)
	if sum != 142913828922 {
		t.Error("Expect sum to be 142913828922", sum)
	}
}
