package main

import (
	"fmt"
)

func main() {
	m := initialize(3, 4)

	t := []int{1, 2, 4, 6}

	addRow(m, t)
	addRow(m, t)

	fmt.Println(m.numCols)
}
