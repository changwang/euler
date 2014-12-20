package challenge36

import (
	"euler/utils"
	"strconv"
)

func GetAllPalindromes(target int64) []int64 {
	palindromes := make([]int64, 0)
	var i int64

	for i = 1; i < target; i++ {
		decmStr := strconv.FormatInt(i, 10)
		binStr := strconv.FormatInt(i, 2)
		if utils.IsPalindrome(decmStr) && utils.IsPalindrome(binStr) {
			palindromes = append(palindromes, i)
		}
	}
	return palindromes
}
