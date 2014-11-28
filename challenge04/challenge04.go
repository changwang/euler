package challenge04

import (
	_ "fmt"
	"strconv"
)

func LargestPalindromicProduct() int32 {
	var factor1, factor2 int32
	var largestProduct int32

	for factor1 = 999; factor1 >= 100; factor1-- {
		for factor2 = factor1; factor2 >= 100; factor2-- {
			product := factor1 * factor2
			//fmt.Printf("processing %d - %d\n", factor1, factor2)
			if isPalindromic(product) {
				if product > largestProduct {
					largestProduct = product
				}
			}
		}
	}
	return largestProduct
}

func isPalindromic(num int32) bool {
	strRepresentation := strconv.Itoa(int(num))
	reversedStr := reverse(strRepresentation)
	reversedInt, _ := strconv.Atoi(reversedStr)
	//fmt.Printf("Original num is %d : reversed num is %d\n", num, reversedInt)
	return int32(reversedInt) == num
}

func reverse(s string) string {
	stringBytes := []byte(s)
	stringLength := len(stringBytes)
	for i := 0; i < stringLength/2; i++ {
		stringBytes[i], stringBytes[stringLength-i-1] = stringBytes[stringLength-i-1], stringBytes[i]
	}
	return string(stringBytes)
}
