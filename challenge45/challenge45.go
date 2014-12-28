package challenge45

import "euler/utils"

func NextNumberSatisfiesAll(start int64) int64 {
	var i int64
	for i = start + 1; ; i++ {
		if utils.IsTriangleNumber(i) && utils.IsPentagonalNumber(i) && utils.IsHexagonalNumber(i) {
			return i
		}
	}
}
