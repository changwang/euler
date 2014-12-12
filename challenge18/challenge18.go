package challenge18

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	space = " "
)

func MaximumPathSum(lines []string) int {
	previous := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		current := extractDigits(line)
		previous = combineLevels(previous, current)
	}
	return maxSum(previous)
}

func maxSum(digits []int) int {
	var max int

	for _, v := range digits {
		if max < v {
			max = v
		}
	}
	return max
}

func extractDigits(line string) []int {
	line = strings.TrimSpace(line)
	digitsString := strings.Split(line, space)
	digits := make([]int, len(digitsString))

	for i, ds := range digitsString {
		digits[i], _ = strconv.Atoi(ds)
	}

	return digits
}

func combineLevels(previous, current []int) []int {
	if len(previous) == 0 || len(current) == 0 {
		if len(previous) == 0 {
			return current[:]
		} else {
			return previous[:]
		}
	}

	newCurrent := make([]int, len(current))
	for i := range newCurrent {
		if i == 0 {
			newCurrent[i] = current[i] + previous[i]
		} else if i == (len(newCurrent) - 1) {
			newCurrent[i] = current[i] + previous[i-1]
		} else {
			left := current[i] + previous[i-1]
			right := current[i] + previous[i]
			newCurrent[i] = int(math.Max(float64(left), float64(right)))
		}
	}
	fmt.Printf("Combined digits array is %v\n", newCurrent)
	return newCurrent
}

/*
func main() {
	triangle :=
		`75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

	buf := strings.Split(triangle, "\n")
	result := challenge18.MaximumPathSum(buf)
	fmt.Printf("maximum sum path %d\n", result)
}
*/
