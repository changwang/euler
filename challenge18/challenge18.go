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
	var previous []int
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
		}
		return previous[:]
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
