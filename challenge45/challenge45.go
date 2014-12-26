package challenge45

import "math"

func NextNumberSatisfiesAll(start int64) int64 {
	var i int64
	for i = start + 1; ; i++ {
		if isTriangleNumber(i) && isPentagonalNumber(i) && isHexagonalNumber(i) {
			return i
		}
	}
}

func isTriangleNumber(num int64) bool {
	num *= 2
	left := int64(math.Floor(math.Sqrt(float64(num))))
	right := left + 1
	//fmt.Printf("%d %d\n", left, right)
	return num == (left * right)
}

func isPentagonalNumber(num int64) bool {
	start := int64(math.Floor(math.Sqrt(float64(num * 2 / 3))))
	end := int64(math.Floor(math.Sqrt(float64(num))))

	var i int64
	for i = start; i <= end; i++ {
		left := i
		right := 3*left - 1
		if (left * right) == (num * 2) {
			//fmt.Printf("%d %d\n", left, right)
			return true
		}
	}
	return false
}

func isHexagonalNumber(num int64) bool {
	start := int64(math.Floor(math.Sqrt(float64(num / 2))))
	end := int64(math.Floor(math.Sqrt(float64(num))))
	var i int64
	for i = start; i <= end; i++ {
		left := i
		right := 2*left - 1
		if (left * right) == num {
			//fmt.Printf("%d %d\n", left, right)
			return true
		}
	}
	return false
}
