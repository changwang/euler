package challenge38

import "testing"

func TestMaxPandigitalMultiples(t *testing.T) {
	result := FindMaxPandigitalMultiples()
	if result != 932718654 {
		t.Error("Expected maximum pandigit is 932718654", result)
	}
}

func TestConcatenatedProduct(t *testing.T) {
	concatProduct := concatenatedProduct(192, 9)
	if concatProduct != 192384576 {
		t.Error("Expected concatenated product 192384576", concatProduct)
	}

	concatProduct = concatenatedProduct(9, 9)
	if concatProduct != 918273645 {
		t.Error("Expected concatenated product 918273645", concatProduct)
	}
}
