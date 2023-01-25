package algorithmX

import (
	"fmt"
	"time"
)

func maxQueensMatrix(n int) *Matrix {
	m := Initialize(2*n, 4*n-2)

	/*
		+-----------------+------------------+---------------------+---------------------------------+------------------------------+
		|        -        | Queen on row (n) | Queen on column (n) | Queen on down diagonal (2n - 1) | Queen on up diagonal (2n -1) |
		+-----------------+------------------+---------------------+---------------------------------+------------------------------+
		| Queen on (1, 1) |                  |                     |                                 |                              |
		| Queen on (1, 2) |                  |                     |                                 |                              |
		| ...             |                  |                     |                                 |                              |
		| Queen on (2, 1) |                  |                     |                                 |                              |
		| ...             |                  |                     |                                 |                              |
		| Queen on (n, n) |                  |                     |                                 |                              |
		+-----------------+------------------+---------------------+---------------------------------+------------------------------+

		down diagonals: c = x - y (c in -(n-1):n-1), so translation: ((x - y) + n) + 2n
		up diagonals: c = x + y (c in 2:2n), so translation ((x + y) - 1) + (2n + 2n - 1)
	*/

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			AddRow(m, []int{x, y + n, ((x - y) + n) + 2*n, ((x + y) - 1) + (2*n + 2*n - 1)})
		}
	}

	return m
}

func PrintMaxQueens(n int) {
	// solve exact cover problem
	m := maxQueensMatrix(n)

	start := time.Now()
	solution := FindFirst(m)
	elapsed := time.Since(start)

	for x := 1; x <= n; x++ {
		fmt.Print("|")
		for y := 1; y <= n; y++ {
			if vecContains(solution, x+(y-1)*n) {
				fmt.Print("X|")
			} else {
				fmt.Print(" |")
			}
		}
		fmt.Println("")
	}

	fmt.Printf("Time: %s\n\n", elapsed)
}

func CountMaxQueens(n int) {
	m := maxQueensMatrix(n)
	start := time.Now()
	solution := FindAll(m)
	elapsed := time.Since(start)

	fmt.Printf("Solutions: %d", len(solution))
	fmt.Println("")

	fmt.Printf("Time: %s\n\n", elapsed)
}

func vecContains(vec []int, el int) bool {
	for _, e := range vec {
		if e == el {
			return true
		}
	}
	return false
}
