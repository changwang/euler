package challenge22

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNameScores(t *testing.T) {
	resp, _ := http.Get("https://projecteuler.net/project/resources/p022_names.txt")
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	names := strings.Split(body, ",")

	for i, name := range names {
		names[i] = strings.Replace(name, "\"", "", 2)
	}

	scores := NameScores(names)
	if scores != 871198282 {
		t.Error("Expect name scores to be 871198282", scores)
	}
}

func TestNameScore(t *testing.T) {
	score := nameScore("COLIN")
	if score != 53 {
		t.Error("Expect score to be 53", score)
	}
}
