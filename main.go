package main

import (
	algorithmX "nienna/Algorithm-X"
)

func main() {
	algorithmX.SolveMondrian(224)

	//fmt.Println(algorithmX.Divisors(20))

	/*
		// Example: n Queens problem

		// Solve the n Queens problem on a 60x60 board
		algorithmX.PrintMaxQueens(60)

		// Count the number of solutions to the n Queens problem on a 13x13 board
		algorithmX.CountMaxQueens(13)

		// Example: Sudoku solving

		// example sudoku
		mat := [9][9]int{{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 3, 0, 8, 5},
			{0, 0, 1, 0, 2, 0, 0, 0, 0},
			{0, 0, 0, 5, 0, 7, 0, 0, 0},
			{0, 0, 4, 0, 0, 0, 1, 0, 0},
			{0, 9, 0, 0, 0, 0, 0, 0, 0},
			{5, 0, 0, 0, 0, 0, 0, 7, 3},
			{0, 0, 2, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 4, 0, 0, 0, 9}}
		s := algorithmX.Sudoku{Board: mat}

		// Solve the sudoku s
		algorithmX.ShowSudoku(s)
		algorithmX.SolveSudoku(&s)
		algorithmX.ShowSudoku(s)

		fmt.Println("")

		// Example: Polyomino packing

		// the 12 pentominoes as example
		f := algorithmX.Polyomino{Tiles: [][]bool{{false, true, true}, {true, true, false}, {false, true, false}}}
		i := algorithmX.Polyomino{Tiles: [][]bool{{true}, {true}, {true}, {true}, {true}}}
		l := algorithmX.Polyomino{Tiles: [][]bool{{true, false}, {true, false}, {true, false}, {true, true}}}
		n := algorithmX.Polyomino{Tiles: [][]bool{{false, true}, {true, true}, {true, false}, {true, false}}}
		p := algorithmX.Polyomino{Tiles: [][]bool{{true, true}, {true, true}, {true, false}}}
		t := algorithmX.Polyomino{Tiles: [][]bool{{true, true, true}, {false, true, false}, {false, true, false}}}
		u := algorithmX.Polyomino{Tiles: [][]bool{{true, false, true}, {true, true, true}}}
		v := algorithmX.Polyomino{Tiles: [][]bool{{true, false, false}, {true, false, false}, {true, true, true}}}
		w := algorithmX.Polyomino{Tiles: [][]bool{{true, false, false}, {true, true, false}, {false, true, true}}}
		x := algorithmX.Polyomino{Tiles: [][]bool{{false, true, false}, {true, true, true}, {false, true, false}}}
		y := algorithmX.Polyomino{Tiles: [][]bool{{false, true}, {true, true}, {false, true}, {false, true}}}
		z := algorithmX.Polyomino{Tiles: [][]bool{{true, true, false}, {false, true, false}, {false, true, true}}}

		pentominoes := []algorithmX.Polyomino{f, i, l, n, p, t, u, v, w, x, y, z}

		// Find a way to pack the pentominoes into the 10x6 rectangle
		algorithmX.SolvePacking(pentominoes, 10, 6)
	*/
}
