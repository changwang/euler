package challenge14

func LongestCollatzSequence(target int64) (int64, int64) {
	var maxNum, maxLength, i int64

	for i = 1; i < target; i++ {
		seqSize := collatzSequence(i)
		if seqSize > maxLength {
			maxNum = i
			maxLength = seqSize
		}
	}
	return maxNum, maxLength
}

func collatzSequence(num int64) int64 {
	var sequence []int64
	sequence = append(sequence, num)
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1
		}
		sequence = append(sequence, num)
	}
	length := len(sequence)
	return int64(length)
}
