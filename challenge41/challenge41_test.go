package challenge41

import "testing"

func TestGetMaxPandigitalPrime(t *testing.T) {
	maxPandigitalPrime := GetMaxPandigitalPrime()
	if maxPandigitalPrime != 7652413 {
		t.Error("Expect maximum n-digit pandigital prime to be 7652413", maxPandigitalPrime)
	}
}
