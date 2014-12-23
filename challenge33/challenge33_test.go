package challenge33

import "testing"

func TestGetAllCancelingPairs(t *testing.T) {
	pairs := GetAllCancelingPairs()
	if len(pairs) != 4 {
		t.Error("Expect canceling pairs not to be 0", len(pairs))
	}
}

func TestIsCancelingPair(t *testing.T) {
	numerator := 49
	demoninator := 98

	if !isCancellingPair(numerator, demoninator) {
		t.Errorf("Expect %d and %d to be canceling pair", numerator, demoninator)
	}

	numerator = 30
	demoninator = 50
	if isCancellingPair(numerator, demoninator) {
		t.Errorf("Expect %d and %d not to be canceling pair", numerator, demoninator)
	}
}
