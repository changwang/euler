package challenge02

func FibonacciSum(target uint64) uint64 {
	var start1, start2 uint64
	start1, start2 = 1, 2
	var fibonacciSum uint64

	for start2 <= target {
		if (start2 % 2) == 0 {
			fibonacciSum += start2
		}
		start1, start2 = start2, (start1 + start2)
	}
	return fibonacciSum
}
