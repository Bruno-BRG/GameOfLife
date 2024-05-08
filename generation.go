package main

import (
	"fmt"
	"os/exec"
	"time"
)

func nextGeneration(grid [][]int, M int, N int, generations int) {
	// Display the grid

	Generations := generations + 1

	fmt.Println("Generations: ", Generations)

	future := make([][]int, M)
	for i := range future {
		future[i] = make([]int, N)
	}

	// Loop through every cell
	for l := 0; l < M; l++ {
		for m := 0; m < N; m++ {
			// finding no Of Neighbours that are alive
			aliveNeighbours := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if l+i >= 0 && l+i < M && m+j >= 0 && m+j < N {
						aliveNeighbours += grid[l+i][m+j]
					}
				}
			}

			// The cell needs to be subtracted from
			// its neighbours as it was counted before
			aliveNeighbours -= grid[l][m]

			// Implementing the Rules of Life

			// Cell is lonely and dies
			if grid[l][m] == 1 && aliveNeighbours < 2 {
				future[l][m] = 0
			} else if grid[l][m] == 1 && aliveNeighbours > 3 { // Cell dies due to over population
				future[l][m] = 0
			} else if grid[l][m] == 0 && aliveNeighbours == 3 { // A new cell is born
				future[l][m] = 1
			} else { // Remains the same
				future[l][m] = grid[l][m]
			}
		}
	}

	fmt.Println("-----------------------")
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
	time.Sleep(5000 * time.Millisecond)
	exec.Command("powershell", "clear").Output()
	nextGeneration(future, M, N, Generations)
}
