package main

import (
	"fmt"

	"github.com/sumanchapai/nqueens/nqueens"
)

func main() {
	for size := 1; size < 9; size++ {
		fmt.Println("Board size", size)
		b := nqueens.New(size)
		fmt.Println(b)
		b.Solve()
		if !b.HasSolution() {
			fmt.Println("No solution for size", size)
		} else {
			fmt.Println("Solution:")
			fmt.Println(b)
		}
		fmt.Printf("\n\n")
	}
}
