package main

import (
	"fmt"
)

func main() {
	m := initialize(3, 4)

	addRow(m, []int{1, 2, 4, 6})
	addRow(m, []int{3, 5})
	addRow(m, []int{2, 5})

	c := &m.columns[2]

	cover(c)

	uncover(c)

	fmt.Println(m.numCols)
}
