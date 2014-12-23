package challenge01

import "testing"

func TestMulti3X5(t *testing.T) {
	sum := Multi3X5(1000)
	if sum != 233168 {
		t.Error("Expect sum of all multiples of 3 or 5 below 1000 to be 233168", sum)
	}
}
