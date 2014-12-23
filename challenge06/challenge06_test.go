package challenge06

import "testing"

func TestSumOfDifference(t *testing.T) {
	diff := SumOfDifference(100)
	if diff != 25164150 {
		t.Error("Expect the difference to be 25164150", diff)
	}
}
