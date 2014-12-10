package challenge17

import (
	"bytes"
)

var wordMap = map[int32]string{
	0:    "",
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

const (
	and = "and"
)

func WorldsCount(target int32) int32 {
	var sum, i int32
	for i = 1; i <= target; i++ {
		sum += int32(len(digitToWord(i)))
	}
	return sum
}

func digitToWord(digit int32) string {
	switch {
	case digit <= 20:
		return lessThanTwenty(digit)
	case digit > 20 && digit < 100:
		return lessThanHundred(digit)
	case digit >= 100:
		return lessThanThousand(digit)
	}
	return ""
}

func lessThanTwenty(digit int32) string {
	if digit > 20 {
		return ""
	}
	return wordMap[digit]
}

func lessThanHundred(digit int32) string {
	if digit > 100 {
		return ""
	}

	var words bytes.Buffer
	remainder := digit % 10
	digit -= remainder
	words.WriteString(wordMap[digit])
	words.WriteString(wordMap[remainder])

	return words.String()
}

func lessThanThousand(digit int32) string {
	if digit > 999 {
		return greaterThanThousand(digit)
	}

	var words bytes.Buffer
	hundredDigit := digit / 100
	words.WriteString(wordMap[hundredDigit])
	words.WriteString(wordMap[100])

	remainder := digit - hundredDigit*100
	if remainder > 0 && remainder <= 20 {
		words.WriteString(and)
		words.WriteString(lessThanTwenty(remainder))
	} else if remainder > 20 {
		words.WriteString(and)
		words.WriteString(lessThanHundred(remainder))
	}

	return words.String()
}

func greaterThanThousand(digit int32) string {
	thousandDigit := digit / 1000
	words := wordMap[thousandDigit] + wordMap[1000]
	return words
}
