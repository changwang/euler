package challenge45

import "testing"

func TestIsTriangleNum(t *testing.T) {
	var num int64
	num = 6
	if !isTriangleNumber(num) {
		t.Error("Expect 6 to be triangle number")
	}

	num = 40755
	if !isTriangleNumber(num) {
		t.Error("Expect 40755 to be triangle number")
	}
}

func TestIsPentagonalNumber(t *testing.T) {
	var num int64
	num = 5
	if !isPentagonalNumber(num) {
		t.Error("Expect 5 to be pentagonal number")
	}

	num = 40755
	if !isPentagonalNumber(num) {
		t.Error("Expect 40755 to be pentagonal number")
	}
}

func TestIsHexagonalNumber(t *testing.T) {
	var num int64
	num = 6
	if !isHexagonalNumber(num) {
		t.Error("Expect 6 to be hexagonal number")
	}

	num = 40755
	if !isHexagonalNumber(num) {
		t.Error("Expect 40755 to be hexagonal number")
	}
}

func TestNextSatisfiesAll(t *testing.T) {
	next := NextNumberSatisfiesAll(40755)
	if next != 1533776805 {
		t.Error("Expect next number to be 1533776805", next)
	}
}
