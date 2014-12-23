package challenge13

func LargeSum(largeDigits []float64) float64 {
	var sum float64
	for _, digit := range largeDigits {
		sum += digit
	}
	return sum
}
