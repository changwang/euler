package challenge11

type Grid [][]int32

type Result struct {
	X, Y    int
	Product int32
}

func LargestGridProduct(distance int, grid Grid, resultChan chan Result) {
	length := len(grid)

	go func() {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				maxUpProduct(i, j, distance, grid, resultChan)
				maxDownProduct(i, j, distance, grid, resultChan)
				maxLeftProduct(i, j, distance, grid, resultChan)
				maxRightProduct(i, j, distance, grid, resultChan)
				maxUpLeftDiagonalProduct(i, j, distance, grid, resultChan)
				maxUpRightDiagonalProduct(i, j, distance, grid, resultChan)
				maxDownLeftDiagonalProduct(i, j, distance, grid, resultChan)
				maxDownRightDiagonalProduct(i, j, distance, grid, resultChan)
			}
		}
		close(resultChan)
	}()
}

func maxUpProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	if x >= distance {
		for i := 0; i < distance; i++ {
			product *= grid[x-i][y]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxDownProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	length := len(grid)
	if (x + distance) < length {
		for i := 0; i < distance; i++ {
			product *= grid[x+i][y]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxLeftProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	if y >= distance {
		for i := 0; i < distance; i++ {
			product *= grid[x][y-i]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxRightProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	length := len(grid)
	if (y + distance) < length {
		for i := 0; i < distance; i++ {
			product *= grid[x][y+i]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxUpLeftDiagonalProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	if y >= distance && x >= distance {
		for i := 0; i < distance; i++ {
			product *= grid[x-i][y-i]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxUpRightDiagonalProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	length := len(grid)
	if (y+distance) <= length && x >= distance {
		for i := 0; i < distance; i++ {
			product *= grid[x-i][y+i]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxDownLeftDiagonalProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	length := len(grid)
	if y >= distance && (x+distance) <= length {
		for i := 0; i < distance; i++ {
			product *= grid[x+i][y-i]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}

func maxDownRightDiagonalProduct(x, y, distance int, grid Grid, productChan chan Result) {
	var product int32
	product = 1
	length := len(grid)
	if (y+distance) <= length && (x+distance) <= length {
		for i := 0; i < distance; i++ {
			product *= grid[x+i][y+i]
		}
	} else {
		product = 0
	}
	productChan <- Result{x, y, product}
}
