package algorithmX

import (
	"fmt"
)

// root node is columns[0]
type Matrix struct {
	columns []Column
	numCols int
	numRows int
	sol     []*Node
}

type Column struct {
	left, right *Column
	head        Node
	length      int
	id          int
}

type Node struct {
	up, down, left, right *Node
	col                   *Column
	value                 int
}

func Initialize(primary int, secondary int) *Matrix {
	numCols := primary + secondary

	// setup first header row of matrix
	m := Matrix{columns: make([]Column, numCols+1), numCols: numCols, numRows: 0}

	// root node has an empty column
	m.columns[0].head.down = &m.columns[0].head
	m.columns[0].head.up = &m.columns[0].head

	m.columns[0].right = &m.columns[1]
	m.columns[0].left = &m.columns[primary]

	// initialize primary columns
	for i := 1; i <= primary; i++ {
		m.columns[i] = Column{left: &m.columns[i-1], right: &m.columns[(i+1)%(primary+1)], id: i}
		m.columns[i].head.down = &m.columns[i].head
		m.columns[i].head.up = &m.columns[i].head
	}

	// initialize secondary columns (not linked into header row)
	for i := primary + 1; i <= numCols; i++ {
		m.columns[i] = Column{left: &m.columns[i], right: &m.columns[i], id: i}
		m.columns[i].head.down = &m.columns[i].head
		m.columns[i].head.up = &m.columns[i].head
	}

	return &m
}

func AddRow(m *Matrix, indices []int) {
	// add sparse row encoded as indices

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
		current := Node{value: m.numRows, col: &m.columns[e], down: &m.columns[e].head, up: m.columns[e].head.up}

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

func cover(c *Column) {
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

func uncover(c *Column) {
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

func heuristic(m *Matrix) *Column {
	// Knuths MRV Heuristic, choose column which is fulfilled by least number of rows
	minLen := -1
	var minCol *Column

	// check if smaller column exists
	for c := m.columns[0].right; c != &m.columns[0]; c = c.right {
		if c.length < minLen || minLen == -1 {
			minLen = c.length
			minCol = c
		}
	}

	return minCol
}

func solve(m *Matrix, collector *[][]int, first bool) []*Node {
	// problem is solved
	if m.columns[0].left == &m.columns[0] {
		var solution []int
		for _, e := range m.sol {
			solution = append(solution, e.value)
		}

		*collector = append(*collector, solution)
		return m.sol
	}

	// MRV heuristic
	col := heuristic(m)
	if col.length == 0 {
		return nil
	}

	cover(col)

	// classic backtracking algorithm
	for r := col.head.down; r != &col.head; r = r.down {
		m.sol = append(m.sol, r)

		for n := r.right; n != r; n = n.right {
			cover(n.col)
		}

		result := solve(m, collector, first)

		undo := m.sol[len(m.sol)-1]
		col = undo.col
		m.sol = m.sol[:len(m.sol)-1]

		for n := undo.left; n != undo; n = n.left {
			uncover(n.col)
		}

		if first && result != nil {
			return result
		}
	}

	uncover(col)

	return nil
}

func FindFirst(m *Matrix) []int {
	var coll [][]int

	solve(m, &coll, true)

	if len(coll) == 0 {
		fmt.Println("Error: No solution found.")
		return nil
	}
	return coll[0]
}

func FindRows(m *Matrix) [][]bool {
	// Return row ids and their content

	sol := solve(m, &[][]int{}, true)

	var rows [][]bool
	for i := 0; i < len(sol); i++ {
		rows = append(rows, make([]bool, m.numCols))
	}

	for i, e := range sol {
		rows[i][e.col.id-1] = true
		for n := e.right; n != e; n = n.right {
			rows[i][n.col.id-1] = true
		}
	}

	return rows
}

func FindAll(m *Matrix) [][]int {
	var coll [][]int

	solve(m, &coll, false)

	return coll
}

func PrettyPrint(m *Matrix) {
	var elements [][]bool
	for i := 0; i < m.numRows; i++ {
		elements = append(elements, make([]bool, m.numCols))
	}

	for i := 1; i <= m.numCols; i++ {
		h := m.columns[i]
		for e := h.head.down; *e != h.head; e = e.down {
			elements[e.value-1][i-1] = true
		}
	}

	numPrimary := 0
	for c := m.columns[0].right; c != &m.columns[0]; c = c.right {
		numPrimary++
	}

	for x := 1; x <= m.numRows; x++ {
		fmt.Print("|")
		for y := 1; y <= m.numCols; y++ {
			if elements[x-1][y-1] {
				fmt.Print("X|")
			} else {
				fmt.Print(" |")
			}
			if y == numPrimary {
				fmt.Print("|")
			}
		}
		fmt.Println("")
	}
}

// force row to solution before starting search
func ForceOption(m *Matrix, id int, value int) {
	// remove row value from m, needs column id with link to row
	col := &m.columns[id]

	// collect columns covered by forced option to cover together later, otherwise not all are found
	var columns []*Column
	columns = append(columns, col)

	for r := col.head.down; r != &col.head; r = r.down {
		if r.value == value {
			m.sol = append(m.sol, r)

			for n := r.right; n != r; n = n.right {
				columns = append(columns, n.col)
			}

			for _, e := range columns {
				cover(e)
			}

			continue
		}
	}
}
