package main

import (
	"fmt"
	"math/rand/v2"
	"os/exec"
	"time"
)

func main() {
	M, N := 10, 10
	generations, randomNumber := 0, 0

	// Initializing the grid.
	grid := make([][]int, M)
	for i := range grid {
		grid[i] = make([]int, N)
	}
	//i want the code to run in a dedicated terminal//

	fmt.Println("Original Generation")
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			randomNumber = (rand.IntN(100))
			if grid[i][j] == 0 && randomNumber < 25 {
				grid[i][j] = 1
				fmt.Print("*")
			} else if grid[i][j] == 0 {
				grid[i][j] = 0
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	time.Sleep(7500 * time.Millisecond)
	exec.Command("powershell", "clear").Output()
	nextGeneration(grid, M, N, generations)
}
