package main

import algorithmX "nienna/Algorithm-X"

func main() {
	/*
		m := algorithmX.Initialize(5, 2)

		algorithmX.AddRow(m, []int{3, 5, 6})
		algorithmX.AddRow(m, []int{1, 4, 7})
		algorithmX.AddRow(m, []int{2, 3, 6})
		algorithmX.AddRow(m, []int{1, 4})
		algorithmX.AddRow(m, []int{2, 7})
		algorithmX.AddRow(m, []int{4, 5, 7})

		algorithmX.ForceOption(m, 2, 3)

		fmt.Println(algorithmX.FindFirst(m))

	*/

	//algorithmX.PrintMaxQueens(60)

	//algorithmX.CountMaxQueens(13)

	mat1 := [9][9]int{{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9}}

	mat2 := [9][9]int{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 3, 0, 8, 5},
		{0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 7, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 1, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 7, 3},
		{0, 0, 2, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 9}}

	s := algorithmX.Sudoku{Board: mat1}

	r := algorithmX.Sudoku{Board: mat2}

	algorithmX.ShowSudoku(s)
	algorithmX.SolveSudoku(&s)
	algorithmX.ShowSudoku(s)

	algorithmX.ShowSudoku(r)
	algorithmX.SolveSudoku(&r)
	algorithmX.ShowSudoku(r)

}
