package challenge29

import "testing"

func TestDistinctPowers(t *testing.T) {
	distinctPowers := DistinctPowers(100)
	if distinctPowers != 9183 {
		t.Error("Expect distinct powers to be 9183", distinctPowers)
	}
}
