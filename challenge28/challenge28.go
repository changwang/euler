package challenge28

func SpiralDiagonalSum(span int) float64 {
	var sum float64
	starter := 1
	sum += float64(starter)

	for i := 1; i < span/2+1; i++ {
		for corner := 0; corner < 4; corner++ {
			starter += (i * 2)
			sum += float64(starter)
		}
	}
	return sum
}
