package main

import (
	"fmt"
)

// root node is columns[0]
type matrix struct {
	columns []column
	numCols int
	numRows int
}

type column struct {
	left, right *column
	head        node
	length      int
	debug       *column // TODO
}

type node struct {
	up, down, left, right *node
	col                   *column
	value                 int
	debug                 *node // TODO
}

func initialize(primary int, secondary int) *matrix {
	numCols := primary + secondary

	// setup first header row of matrix
	m := matrix{columns: make([]column, numCols+1), numCols: numCols, numRows: 0}

	// root node has an empty column
	m.columns[0].head.down = &m.columns[0].head
	m.columns[0].head.up = &m.columns[0].head

	m.columns[0].right = &m.columns[1]
	m.columns[0].left = &m.columns[primary]

	m.columns[0].debug = &m.columns[0] // TODO

	// initialize primary columns
	for i := 1; i <= primary; i++ {
		m.columns[i] = column{left: &m.columns[i-1], right: &m.columns[(i+1)%(primary+1)]}
		m.columns[i].debug = &m.columns[i]           // TODO
		m.columns[i].head.debug = &m.columns[i].head // TODO
		m.columns[i].head.down = &m.columns[i].head
		m.columns[i].head.up = &m.columns[i].head
	}

	// initialize secondary columns (not linked into header row)
	for i := primary + 1; i <= numCols; i++ {
		m.columns[i] = column{left: &m.columns[i], right: &m.columns[i]}
		m.columns[i].debug = &m.columns[i]           // TODO
		m.columns[i].head.debug = &m.columns[i].head // TODO
		m.columns[i].head.down = &m.columns[i].head
		m.columns[i].head.up = &m.columns[i].head
	}

	return &m
}

// add sparse row encoded as indices
func addRow(m *matrix, indices []int) {

	// input verification
	last := -1
	for _, element := range indices {
		if element <= 0 || element > m.numCols {
			fmt.Println("Error: Index out of range.")
		}
		if element <= last {
			fmt.Println("Error: Indices not ordered")
		}
		last = element
	}

	m.numRows++

	for _, e := range indices {
		// insert new node in last row of column
		current := node{value: m.numRows, col: &m.columns[e], down: &m.columns[e].head, up: m.columns[e].head.up}
		current.debug = &current // TODO

		m.columns[e].head.up.down = &current
		m.columns[e].head.up = &current

		m.columns[e].length++
	}

	// link new nodes to left/right neighbour (head.up still points to new node)
	for i, e := range indices {
		m.columns[e].head.up.right = m.columns[indices[(i+1)%len(indices)]].head.up
		m.columns[indices[(i+1)%len(indices)]].head.up.left = m.columns[e].head.up
	}
}

func cover(c *column) {
	// remove from header row
	c.left.right = c.right
	c.right.left = c.left

	// cover all rows covered by column c
	for row := c.head.down; row != &(c.head); row = row.down {
		for e := row.right; e != row; e = e.right {
			e.up.down = e.down
			e.down.up = e.up

			// decrease size of column
			e.col.length--
		}
	}
}

func coverCol(n *node) {
	// cover all other columns of selected row

	for e := n.right; e != n; e = e.right {
		cover(e.col)
	}
}

func uncover(c *column) {
	// uncover all rows covered by column c
	for row := c.head.up; row != &(c.head); row = row.up {
		for e := row.left; e != row; e = e.left {
			e.up.down = e
			e.down.up = e

			e.col.length++
		}
	}

	c.left.right = c
	c.right.left = c
}

func uncoverCol(n *node) {
	// cover all other columns of selected row

	for e := n.left; e != n; e = e.left {
		uncover(e.col)
	}
}

func heuristic(m *matrix) *column {
	// Knuths MRV Heuristic, choose column which is fulfilled by least number of rows

	// start with column 1
	minLen := m.columns[1].length
	minCol := &m.columns[1]

	// check if smaller column exists
	for c := m.columns[1].right; c != &m.columns[0]; c = c.right {
		if c.length < minLen {
			minLen = c.length
			minCol = c
		}
	}

	return minCol
}
