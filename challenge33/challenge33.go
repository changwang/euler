package challenge33

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	numer int
	demon int
}

const (
	zero = "0"
)

var nullPair = Pair{0, 0}

func GetAllCancelingPairs() []Pair {
	pairs := make([]Pair, 0)

	for i := 10; i < 99; i++ {
		for j := i + 1; j < 99; j++ {
			p := simplify(i, j)
			origPair := Pair{i, j}
			if p != nullPair && !contains(pairs, p) && isEqual(p, origPair) {
				pairs = append(pairs, origPair)
			}
		}
	}
	return pairs
}

func isEqual(sPair, oPair Pair) bool {
	return float64(sPair.numer)/float64(sPair.demon) == float64(oPair.numer)/float64(oPair.demon)
}

func simplify(numerator, demoninator int) Pair {
	numerString := fmt.Sprintf("%d", numerator)
	demonString := fmt.Sprintf("%d", demoninator)

	if string(numerString[1]) == zero && string(demonString[1]) == zero {
		return nullPair
	}

	for i, s := range numerString {
		idx := strings.Index(demonString, string(s))
		if idx != -1 {
			subNum := substringToI(numerString, i)
			subDem := substringToI(demonString, idx)
			if subNum == 0 || subDem == 0 {
				return nullPair
			}
			return Pair{subNum, subDem}
		}

	}
	return nullPair
}

func contains(pairs []Pair, p Pair) bool {
	for _, pa := range pairs {
		if pa.numer == p.numer && pa.demon == p.demon {
			return true
		}
	}
	return false
}

func substringToI(str string, idx int) int {
	var subInt int
	if idx == 0 {
		subInt, _ = strconv.Atoi(string(str[1]))
	} else if idx == 1 {
		subInt, _ = strconv.Atoi(string(str[0]))
	}
	return subInt
}
