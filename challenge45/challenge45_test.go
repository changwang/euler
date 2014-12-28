package challenge45

import "testing"

func TestNextSatisfiesAll(t *testing.T) {
	next := NextNumberSatisfiesAll(40755)
	if next != 1533776805 {
		t.Error("Expect next number to be 1533776805", next)
	}
}
