package challenge06

func SumOfDifference(target int64) int64 {
	return squareOfSum(target) - sumOfSquares(target)
}

func sumOfSquares(target int64) int64 {
	var i int64
	i = 1

	var sum int64

	for ; i <= target; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(target int64) int64 {
	var i int64
	i = 1

	var sum int64

	for ; i <= target; i++ {
		sum += i
	}
	return sum * sum
}
