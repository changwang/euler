package challenge42

import "testing"

func TestGetNumOfTriangleWords(t *testing.T) {
	num := GetNumOfTriangleWords("https://projecteuler.net/project/resources/p042_words.txt")
	if num != 162 {
		t.Error("Expect 162 triangle words", num)
	}
}

func TestIsTriangleNumber(t *testing.T) {
	var num int64
	num = 1
	if !isTriangleNumber(num) {
		t.Errorf("Expect %d to be triangle number\n", num)
	}

	num = 5
	if isTriangleNumber(num) {
		t.Errorf("Expect %d not to be triangle number\n", num)
	}
}

func TestIsTriangleWord(t *testing.T) {
	input := "SKY"
	if !isTriangleWord(input) {
		t.Errorf("Expect %s to be triangle word\n", input)
	}
}
