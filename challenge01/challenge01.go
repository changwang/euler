package challenge01

func Multi3X5(max int32) int64 {
	num := int32(1)
	result := int64(0)

	for num < max {
		if hasFactor3(num) || hasFactor5(num) {
			result += int64(num)
		}
		num += 1
	}
	return result
}

func hasFactor3(num int32) bool {
	return (num % 3) == 0
}

func hasFactor5(num int32) bool {
	return (num % 5) == 0
}
