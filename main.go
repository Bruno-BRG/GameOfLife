package main

import (
	"fmt"
)

func main() {
	M := 10
	N := 10

	grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Println("Original Generation")
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	nextGeneration(grid, M, N)
}

func nextGeneration(grid [][]int, M int, N int) {
	future := make([][]int, M, N)
	for l := 0; l < M; l++ {
		for m := 0; m < N; m++ {
			aliveNeighbors := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if l+i >= 0 && l+i < M && m+j >= 0 && m+j < N && grid[l+i][m+j] == 1 {
						aliveNeighbors++
					}
				}
			}
			aliveNeighbors = aliveNeighbors - grid[l][m]

			if grid[l][m] == 1 && aliveNeighbors < 2 {
				future[l][m] = 0
			} else if grid[l][m] == 1 && aliveNeighbors > 3 {
				future[l][m] = 0
			} else if grid[l][m] == 0 && aliveNeighbors == 3 {
				future[l][m] = 1
			} else {
				future[l][m] = grid[l][m]
			}

		}
	}
	fmt.Println("Next Generation")
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if future[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
