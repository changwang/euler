package challenge35

import "testing"

func TestGetAllCircularPrimes(t *testing.T) {
	circlarsPrimes := GetAllCircularPrimes()
	if len(circlarsPrimes) != 55 {
		t.Error("Expect number of circular primes to be 0", len(circlarsPrimes))
	}
}
