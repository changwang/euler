package challenge37

import "testing"

func TestAllTruncatablePrimes(t *testing.T) {
	result := GetAllTruncatablePrimes()
	if result != 748317 {
		t.Error("Expect sum of all truncatable primes is 748317", result)
	}
}

func TestIsTruncatablePrime(t *testing.T) {
	var num int64

	num = 3797
	if !IsTruncatablePrime(num) {
		t.Errorf("expect %d to be truncatable prime", num)
	}

	num = 7
	if IsTruncatablePrime(num) {
		t.Errorf("expect %d not to be truncatable prime", num)
	}

	num = 2733
	if IsTruncatablePrime(num) {
		t.Errorf("expect %d not to be truncatable prime", num)
	}
}
