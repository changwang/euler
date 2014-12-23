package challenge14

import "testing"

func TestLongestCollatzSequence(t *testing.T) {
	maxNum, _ := LongestCollatzSequence(1000000)
	if maxNum != 837799 {
		t.Error("Expect maximum number to be 837799", maxNum)
	}
}

func TestCollatzSequence(t *testing.T) {
	length := collatzSequence(13)
	if length != 10 {
		t.Error("Expect 10 terms", length)
	}
}
