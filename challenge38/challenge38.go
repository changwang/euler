package challenge38

import (
	"bytes"
	"euler/utils"
	"fmt"
	"strconv"
)

func FindMaxPandigitalMultiples() int64 {
	var max, i int64

	for i = 1; i < 10000; i++ {
		concat := concatenatedProduct(i, 9)
		if !utils.IsPandigital(concat, 9) {
			continue
		}
		if max < concat {
			max = concat
		}
	}
	return max
}

func concatenatedProduct(num, length int64) int64 {
	var buffer bytes.Buffer
	var i int64
	for i = 1; i <= length; i++ {
		product := num * i
		buffer.Write([]byte(fmt.Sprintf("%d", product)))
		currentConcat := buffer.String()
		currentLen := int64(len(currentConcat))
		if currentLen > length {
			return -1
		}

		if currentLen == length {
			result, err := strconv.ParseInt(currentConcat, 10, 64)
			if err != nil {
				return -1
			}
			return result
		}
	}
	return -1
}
