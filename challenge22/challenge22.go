package challenge22

import (
	"sort"
)

const (
	MINIMUM_CHAR = 64
)

func NameScores(names sort.StringSlice) float64 {
	var totalNameScores float64
	sort.Sort(names)
	for i, name := range names {
		score := nameScore(name)
		totalNameScores += (score * float64(i+1))
	}
	return totalNameScores
}

func nameScore(name string) float64 {
	var score float64
	for _, c := range name {
		score += float64(c - MINIMUM_CHAR)
	}
	return score
}
