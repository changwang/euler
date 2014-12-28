package utils

import (
	"bytes"
	"errors"
	"math"
	"strconv"
)

var occurred = map[int64]bool{
	1: false,
	2: false,
	3: false,
	4: false,
	5: false,
	6: false,
	7: false,
	8: false,
	9: false,
}

const (
	MinCharValue = 64
)

func IsPrime(num int64) bool {
	if num < 2 {
		return false
	}

	var sqrt, i int64
	sqrt = int64(math.Sqrt(float64(num)))
	for i = 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Contains(intSlice []int64, num int64) bool {
	for _, d := range intSlice {
		if num == d {
			return true
		}
	}
	return false
}

func SliceSum(intSlice []int64) int64 {
	var sum int64

	for _, d := range intSlice {
		sum += d
	}
	return sum
}

func ProperDivisors(num int64) []int64 {
	if num == 1 {
		return []int64{0}
	}

	stopper := int64(math.Ceil(math.Sqrt(float64(num))))
	var i int64
	var divisors []int64
	divisors = append(divisors, 1)

	for i = 2; i <= stopper; i++ {
		if num%i == 0 {
			if !Contains(divisors, i) && i < num {
				divisors = append(divisors, i)
			}
			if !Contains(divisors, num/i) && num/i < num {
				divisors = append(divisors, num/i)
			}
		}
	}
	return divisors
}

func Factorial(num int64) float64 {
	if num == 0 || num == 1 {
		return 1
	}

	var factorial float64
	factorial = 1
	for num > 1 {
		factorial *= float64(num)
		num--
	}
	return factorial
}

func IsPalindrome(input string) bool {
	length := len(input)
	if length == 1 {
		return true
	}

	var start, end int
	if length%2 != 0 {
		start = length/2 - 1
		end = length/2 + 1
	} else {
		start, end = length/2-1, length/2
	}

	for start >= 0 && end < length {
		if input[start] != input[end] {
			return false
		}
		start--
		end++
	}

	return true
}

func DigitLength(num int64) int {
	var length int

	if num == 0 {
		return 1
	}

	if num < 0 {
		num -= num
	}

	for num > 0 {
		num /= 10
		length++
	}
	return length
}

func CreatePandigitals(start, span int64) []int64 {
	if start+span > 10 {
		return nil
	}

	var i int64

	var digitSlice []string
	for i = start; i < (start + span); i++ {
		ds := strconv.FormatInt(i, 10)
		digitSlice = append(digitSlice, ds)
	}
	digits, err := createPandigits(digitSlice, span)
	if err != nil {
		return nil
	}
	return digits
}

func createPandigits(digits []string, span int64) ([]int64, error) {
	if int64(len(digits)) != span {
		return nil, errors.New("number of digits mismatches span")
	}

	var pandigits []int64
	if span == 1 {
		digit, err := strconv.ParseInt(digits[0], 10, 64)
		if err != nil {
			return nil, err
		}
		pandigits = append(pandigits, digit)
		return pandigits, nil
	}

	for i := 0; i < len(digits); i++ {
		head := digits[i]
		rest := make([]string, len(digits))
		copy(rest, digits)
		rest = append(rest[:i], rest[i+1:]...)
		restdigits, err := createPandigits(rest, int64(len(rest)))
		if err != nil {
			return nil, err
		}

		for _, digit := range restdigits {
			var buffer bytes.Buffer
			buffer.Write([]byte(head))
			ds := strconv.FormatInt(digit, 10)
			buffer.Write([]byte(ds))
			permu, err := strconv.ParseInt(buffer.String(), 10, 64)
			if err != nil {
				return nil, err
			}
			pandigits = append(pandigits, permu)
		}
	}

	return pandigits, nil
}

func IsPandigital(num, span int64) bool {
	length := int64(DigitLength(num))
	defer resetCheckMap()
	if length != span {
		return false
	}

	for num > 0 {
		rem := num % 10
		if rem > span {
			return false
		}

		num /= 10
		found, ok := occurred[rem]
		if !ok || found {
			return false
		}

		occurred[rem] = true
	}
	return true
}

func resetCheckMap() {
	for k := range occurred {
		occurred[k] = false
	}
}

func StringScore(input string) int64 {
	var score int64
	for _, c := range input {
		score += int64(c - MinCharValue)
	}
	return score
}

func IsTriangleNumber(num int64) bool {
	num *= 2
	left := int64(math.Floor(math.Sqrt(float64(num))))
	right := left + 1
	//fmt.Printf("%d %d\n", left, right)
	return num == (left * right)
}

func IsPentagonalNumber(num int64) bool {
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

func IsHexagonalNumber(num int64) bool {
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
