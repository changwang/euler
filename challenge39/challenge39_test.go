package challenge39

import (
	"testing"
)

func TestIsRightAngleTriangle(t *testing.T) {
	tri := Triangle{20, 48, 52}
	if !tri.isRightAngle() {
		t.Error("Expect to be right angle triangle")
	}
}

func TestGetAllRightAngleTriangles(t *testing.T) {
	triangles := getAllRightAngleTrianglesWith(120)
	if len(triangles) != 3 {
		t.Error("Expect 3 different combinations")
	}
}

func TestGetAllSolutions(t *testing.T) {
	perimeter, tris := GetMaxSolutionRightAngleTriangle(1000)

	if perimeter != 840 {
		t.Errorf("Expect maximum %d solutions for %d\n", len(tris), perimeter)
	}
}
