package challenge34

import (
	"bytes"
	"euler/utils"
	"fmt"
	"strconv"
)

const (
	space       = " "
	leftParen   = "("
	righttParen = ")"
)

func GetAllDigitFactorials() []int64 {
	all := make([]int64, 0)
	for i := 3; i < 100000; i++ {
		str, ok := IsDigitFactorial(int64(i))
		if ok {
			all = append(all, int64(i))
			fmt.Printf("%s\n", str)
		}
	}

	return all
}

func IsDigitFactorial(num int64) (string, bool) {
	var buffer bytes.Buffer

	copyNum := num
	buffer.Write([]byte(strconv.Itoa(int(copyNum)) + space))
	var sum int64

	for num > 0 {
		rem := num % 10
		fac := utils.Factorial(rem)
		if fac > float64(copyNum) || sum > copyNum {
			return "", false
		}
		buffer.Write([]byte(space + leftParen + strconv.Itoa(int(rem)) + space + strconv.Itoa(int(fac)) + righttParen))
		sum += int64(utils.Factorial(rem))
		num /= 10
	}

	buffer.Write([]byte(space + strconv.Itoa(int(sum))))

	if copyNum == sum {
		return buffer.String(), true
	}
	return "", false
}
