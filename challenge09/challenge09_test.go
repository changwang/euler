package challenge09

import "testing"

func TestPythagoreanTripletProduct(t *testing.T) {
	product := PythagoreanTripletProduct()
	if product != 31875000 {
		t.Error("Expect product to be 31875000", product)
	}
}
