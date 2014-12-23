package challenge12

import "testing"

func TestTriangleNumbers(t *testing.T) {
	triNum := TriangleNumbers(500)
	if triNum != 76576500 {
		t.Error("Expecting first triangle number to be 76576500", triNum)
	}
}
