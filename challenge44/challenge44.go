package challenge44

import "euler/utils"

func GetMinimumDiff() int64 {
	var i, j int64
	var min int64
	min = 1<<63 - 1
	for i = 2; i < 3000; i++ {
		pentagonalI := pentagonalNumber(i)
		for j = i - 1; j > 0; j-- {
			pentagonalJ := pentagonalNumber(j)
			if utils.IsPentagonalNumber(pentagonalI+pentagonalJ) && utils.IsPentagonalNumber(pentagonalI-pentagonalJ) {
				if (pentagonalI - pentagonalJ) < min {
					min = pentagonalI - pentagonalJ
				}
			}
		}
	}
	return min
}

func pentagonalNumber(index int64) int64 {
	return index * (3*index - 1) / 2
}
