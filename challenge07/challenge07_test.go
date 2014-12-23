package challenge07

import "testing"

func TestNthPrime(t *testing.T) {
	nprime := NthPrime(10001)
	if nprime != 104743 {
		t.Error("Expect 10001th prime number to be 104743", nprime)
	}
}
