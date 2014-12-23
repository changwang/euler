package challenge04

import "testing"

func TestLargestPalindromicProduct(t *testing.T) {
	product := LargestPalindromicProduct()
	if product != 906609 {
		t.Error("Expect largest palindrome product to be 906609", product)
	}
}
