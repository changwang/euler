package challenge25

import "testing"

func TestNLengthFibonacci(t *testing.T) {
	term := NLengthFibonacci(1000)
	if term != 4782 {
		t.Error("Expect first term to be 4782", term)
	}
}
