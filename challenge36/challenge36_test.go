package challenge36

import (
	"euler/utils"
	"testing"
)

func TestGetAllPalindromes(t *testing.T) {
	palindromes := GetAllPalindromes(1000000)
	if !utils.Contains(palindromes, 585) {
		t.Error("Expect 585 to be double-base palindromes")
	}

	sum := utils.SliceSum(palindromes)
	if sum != 872187 {
		t.Error("Expect sum of all double-base palindromes to be 872187", sum)
	}
}
