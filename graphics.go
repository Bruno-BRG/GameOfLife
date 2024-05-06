package main

import (
	"os/exec"
	"time"
)

// display the grid in the terminal updating it every 2 seconds with the next generations//
func displayGraphics(grid [][]int, M int, N int) {
	var cmd *exec.Cmd
	cmd = exec.Command("powershell.exe", "-NoExit")
	for {
		exec.Command("powershell.exe", "clear").Run()
		nextGeneration(grid, M, N)
		time.Sleep(3)
	}
}
