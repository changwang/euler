package challenge05

import "testing"

func TestSmallestMultiple(t *testing.T) {
	product := SmallestMultiple(20)
	if product != 232792560 {
		t.Error("Expect smallest positive number to be 232792560", product)
	}
}
