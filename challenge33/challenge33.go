package challenge33

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
			p := Pair{i, j}
			if isCancellingPair(i, j) && !contains(pairs, p) {
				pairs = append(pairs, p)
			}
		}
	}
	return pairs
}

func isCancellingPair(numerator, demoninator int) bool {
	copyNumer := numerator
	copyDemon := demoninator

	numRem := numerator % 10
	numerator /= 10
	if numRem == 0 {
		return false
	}

	demRem := demoninator % 10
	demoninator /= 10
	if demRem == 0 {
		return false
	}

	if isEqual(numRem, demRem, numerator, demoninator, copyNumer, copyDemon) {
		return true
	}

	if isEqual(numerator, demRem, numRem, demoninator, copyNumer, copyDemon) {
		return true
	}

	if isEqual(numRem, demoninator, numerator, demRem, copyNumer, copyDemon) {
		return true
	}

	if isEqual(numerator, demoninator, numRem, demRem, copyNumer, copyDemon) {
		return true
	}

	return false
}

func isEqual(smp1, smp2, rem1, rem2, orig1, orig2 int) bool {
	return smp1 == smp2 && float64(rem1)/float64(rem2) == float64(orig1)/float64(orig2)
}

func contains(pairs []Pair, p Pair) bool {
	for _, pa := range pairs {
		if pa.numer == p.numer && pa.demon == p.demon {
			return true
		}
	}
	return false
}
