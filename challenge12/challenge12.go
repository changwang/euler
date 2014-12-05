package challenge12

import (
	"fmt"
	"math"
	"sort"
)

func TriangleNumbers(factorSize int) int {
	var triNum, i int
	triNum = 0
	i = 1
	for {
		triNum += i
		factors := getFactorSlice(triNum)
		length := len(factors)
		if length > factorSize {
			fmt.Printf("Triangle number %d has factors %v\n", triNum, factors)
			return triNum
		}
		i++
	}
}

func getFactorSlice(num int) []int {
	factorMap := make(map[int]int)
	sqrt := int(math.Sqrt(float64(num)))
	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			_, ok := factorMap[i]
			if !ok {
				factorMap[i] = i
			}
			_, ok = factorMap[num/i]
			if !ok {
				factorMap[num/i] = num / i
			}
		}
	}
	return getKeySlice(factorMap)
}

func getKeySlice(factorMap map[int]int) sort.IntSlice {
	var keys sort.IntSlice
	for k, _ := range factorMap {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}
