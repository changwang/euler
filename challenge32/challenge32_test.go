package challenge32

import "testing"

func TestFindAllPandigitalProducts(t *testing.T) {
	sum := FindAllPandigitalProducts()
	if sum != 45228 {
		t.Error("Expect sum to be 45228", sum)
	}
}
