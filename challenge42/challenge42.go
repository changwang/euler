package challenge42

import (
	"euler/utils"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

func GetNumOfTriangleWords(url string) int64 {
	resp, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	words := strings.Split(body, ",")

	var triangles []string
	for _, word := range words {
		w := strings.Replace(word, "\"", "", 2)
		if isTriangleWord(w) {
			triangles = append(triangles, w)
		}
	}
	return int64(len(triangles))
}

func isTriangleWord(input string) bool {
	score := utils.StringScore(input)
	return isTriangleNumber(score)
}

func isTriangleNumber(num int64) bool {
	if num < 0 {
		return false
	}

	double := num * 2
	smaller := int64(math.Floor(math.Sqrt(float64(double))))
	larger := smaller + 1
	if smaller*larger != double {
		return false
	}
	return true
}
