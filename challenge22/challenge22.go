package challenge22

import (
	"euler/utils"
	"sort"
)

func NameScores(names sort.StringSlice) int64 {
	var totalNameScores int64
	sort.Sort(names)
	for i, name := range names {
		score := utils.StringScore(name)
		totalNameScores += (score * int64(i+1))
	}
	return totalNameScores
}
