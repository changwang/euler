package challenge15

import (
	"fmt"
)

func LatticePaths(gridSize int) int64 {
	grid := make([][]int64, gridSize+1)

	for i := 0; i <= gridSize; i++ {
		grid[i] = make([]int64, gridSize+1)
	}

	for i := 0; i <= gridSize; i++ {
		grid[gridSize][i] = 1
	}

	for i := 0; i < gridSize; i++ {
		grid[i][gridSize] = 1
	}

	grid[gridSize][gridSize] = 0

	for i := gridSize - 1; i >= 0; i-- {
		for j := gridSize - 1; j >= 0; j-- {
			grid[i][j] = grid[i][j+1] + grid[i+1][j]
		}
	}

	printGrid(grid)

	return grid[0][0]
}

func printGrid(grid [][]int64) {
	length := len(grid)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Printf("%d\t", grid[i][j])
		}
		fmt.Println()
	}
}
