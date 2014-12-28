package challenge44

import "testing"

func TestGetMinimumDiff(t *testing.T) {
	min := GetMinimumDiff()
	if min != 5482660 {
		t.Error("Expect min to be 5482660", min)
	}
}
