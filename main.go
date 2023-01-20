package main

import (
	"fmt"
)

func main() {
	m := initialize(5, 2)

	addRow(m, []int{3, 5, 6})
	addRow(m, []int{1, 4, 7})
	addRow(m, []int{2, 3, 6})
	addRow(m, []int{1, 4})
	addRow(m, []int{2, 7})
	addRow(m, []int{4, 5, 7})

	n := initialize(4, 0)
	addRow(n, []int{1, 2, 4})
	addRow(n, []int{3})
	addRow(n, []int{2, 3})
	addRow(n, []int{1, 4})

	coll := findAll(m)

	fmt.Println(coll)
}
