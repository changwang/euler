package challenge39

import (
	"fmt"
)

type Triangle struct {
	a int64
	b int64
	c int64
}

func (tri *Triangle) isRightAngle() bool {
	return (tri.a*tri.a + tri.b*tri.b) == tri.c*tri.c
}

func (tri *Triangle) String() string {
	return fmt.Sprintf("%d - %d - %d : (%d - %d) - %d", tri.a, tri.b, tri.c, tri.a*tri.a, tri.b*tri.b, tri.c*tri.c)
}

func getAllRightAngleTrianglesWith(perimeter int64) []Triangle {
	triangles := make([]Triangle, 0)
	var i, j int64

	for i = 1; i < perimeter-2; i++ {
		for j = i; j < perimeter-2; j++ {
			tri := Triangle{i, j, perimeter - i - j}
			if tri.isRightAngle() {
				triangles = append(triangles, tri)
			}
		}
	}

	return triangles
}

func GetMaxSolutionRightAngleTriangle(target int64) (int64, []Triangle) {
	var max, p, i int64
	var triangles []Triangle
	for i = 12; i <= target; i++ {
		tris := getAllRightAngleTrianglesWith(i)
		if int64(len(tris)) > max {
			max = int64(len(tris))
			p = i
			triangles = tris
		}
	}
	return p, triangles
}
