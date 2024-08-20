package main

import (
	"fmt"
)

func render(maze [][]string) {
	// TODO figure out a way to clear the command line
	// syscall.Exec("clear")
	// for i := 0; i < 1; i++ {

	// }
	for _, row := range maze {
		for _, val := range row {	
			fmt.Print(val)
		}
		fmt.Println()
	}

	for i := 0; i < len(maze); i++ {
		fmt.Print(escUp)
	}
}
