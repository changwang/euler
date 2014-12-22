package utils

import "testing"

func TestIsPrime(t *testing.T) {
	var num int64
	num = 2
	result := IsPrime(num)
	if !result {
		t.Error("Expect 2 to be prime", result)
	}

	num = 1
	result = IsPrime(num)
	if result {
		t.Error("Expect 1 not to be prime", result)
	}

	num = 9
	result = IsPrime(num)
	if result {
		t.Error("Expect 9 not to be prime", result)
	}

	num = 15485863
	result = IsPrime(num)
	if !result {
		t.Error("Expect 15485863 not to be prime", result)
	}
}

func TestContains(t *testing.T) {
	container := make([]int32, 1)
	container[0] = 11

	if !Contains(container, 11) {
		t.Error("Expect container contains 11")
	}

	if Contains(container, 10) {
		t.Error("Expect container does not contain 10")
	}

	var i int32
	for i = 1; i < 100; i++ {
		container = append(container, i)
		if !Contains(container, i) {
			t.Errorf("Expect contains cotnains %d\n", i)
		}
	}
}

func TestSliceSum(t *testing.T) {
	var i int32
	intSlice := make([]int32, 0)
	for i = 1; i < 11; i++ {
		intSlice = append(intSlice, i)
	}

	result := SliceSum(intSlice)
	if result != 55 {
		t.Error("expected sum to be 55")
	}
}

func TestFactorial(t *testing.T) {
	var num int64

	fact := Factorial(num)
	if fact != 1 {
		t.Errorf("Expect %d's factorial to be 1\n", num)
	}

	num = 1
	fact = Factorial(num)
	if fact != 1 {
		t.Errorf("Expect %d's factorial to be 1\n", num)
	}

	num = 10
	fact = Factorial(num)
	if fact != 3628800 {
		t.Errorf("Expect %d's factorial to be 3628800\n", num)
	}
}

func TestIsPalindrome(t *testing.T) {
	input := "abcdcba"
	if !IsPalindrome(input) {
		t.Errorf("expect %s to be palindrome\n", input)
	}

	input = "100001"
	if !IsPalindrome(input) {
		t.Errorf("expect %s to be palindrome\n", input)
	}

	input = "ab"
	if IsPalindrome(input) {
		t.Errorf("expect %s not to be palindrome\n", input)
	}

	input = "100101"
	if IsPalindrome(input) {
		t.Errorf("expect %s not to be palindrome\n", input)
	}
}

func TestDigitLength(t *testing.T) {
	var num int64
	nLength := DigitLength(num)
	if nLength != 1 {
		t.Error("Expect num length is 1", nLength)
	}

	num = 123456789
	nLength = DigitLength(num)
	if nLength != 9 {
		t.Error("Expect num length is 9", nLength)
	}

	num = 1234567890
	nLength = DigitLength(num)
	if nLength != 10 {
		t.Error("Expect num length is 10", nLength)
	}
}

func TestPandigital(t *testing.T) {
	var num int64
	num = 12345
	numLen := int64(DigitLength(num))
	if !IsPandigital(num, numLen) {
		t.Errorf("Expect %d is pandigit\n", num)
	}

	num = 123456789
	numLen = int64(DigitLength(num))
	if !IsPandigital(num, numLen) {
		t.Errorf("Expect %d is pandigit\n", num)
	}

	num = 12346
	if IsPandigital(num, numLen) {
		t.Errorf("Expect %d is not pandigit\n", num)
	}

	num = 13579
	if IsPandigital(num, numLen) {
		t.Errorf("Expect %d is not pandigit\n", num)
	}

	num = 13579
	if IsPandigital(num, 9) {
		t.Errorf("Expect %d is not pandigit\n", num)
	}

	num = 1234567890
	numLen = int64(DigitLength(num))
	if IsPandigital(num, numLen) {
		t.Errorf("Expect %d is not pandigit\n", num)
	}
}
