package challenge17

import "testing"

func TestWordCount(t *testing.T) {
	sum := WorldsCount(1000)
	if sum != 21124 {
		t.Error("Expect sum to be 21124", sum)
	}
}
