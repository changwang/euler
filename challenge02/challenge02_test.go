package challenge02

import "testing"

func TestFibonacciSum(t *testing.T) {
	sum := FibonacciSum(4000000)
	if sum != 4613732 {
		t.Error("Expect even-valued sum to be 4613732", sum)
	}
}
