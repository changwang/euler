package challenge27

import "testing"

func TestQuadraticCoefficient(t *testing.T) {
	product := QuadraticCoefficient(1000)
	if product != -59231 {
		t.Error("Expect product to be -59231", product)
	}
}
