package challenge15

import "testing"

func TestLatticePaths(t *testing.T) {
	sum := LatticePaths(20)
	if sum != 137846528820 {
		t.Error("Expect number of paths to be 137846528820", sum)
	}
}
