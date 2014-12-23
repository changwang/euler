package challenge34

import (
	"euler/utils"
	"testing"
)

func TestGetAllDigitFactorials(t *testing.T) {
	digitFactorials := GetAllDigitFactorials()
	sum := utils.SliceSum(digitFactorials)
	if sum != 40730 {
		t.Error("Expect sum to be 40730", sum)
	}
}

func TestIsDigitalFactorial(t *testing.T) {
	var num int64

	num = 145
	_, is := IsDigitFactorial(num)
	if !is {
		t.Errorf("Expect %d to be digit factorials", num)
	}

	num = 232
	_, is = IsDigitFactorial(num)
	if is {
		t.Errorf("Expect %d not to be digit factorials", num)
	}
}
