package challenge03

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	largestPrimeFactor := LargestPrime(600851475143)
	if largestPrimeFactor != 6857 {
		t.Error("Expect largest prime factor to be 6857", largestPrimeFactor)
	}
}
