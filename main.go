package main

import (
	algorithmX "nienna/Algorithm-X"
)

func main() {
	// pentominoes
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

	algorithmX.SolvePacking(pentominoes, 10, 6)
}
