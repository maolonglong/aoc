package main

import (
	"bytes"
	"fmt"
	"math"
	"os"

	"github.com/samber/lo"
)

type Pair = lo.Tuple2[int, int]

var (
	_dummy = lo.Tuple2[int, int]{}

	E = lo.T2(0, 1)
	W = lo.T2(0, -1)
	S = lo.T2(1, 0)
	N = lo.T2(-1, 0)
)

var _dirs = []Pair{E, W, S, N}

var _m = map[Pair] /* from */ map[byte]Pair /* to */ {
	E: {
		'-': W,
		'L': N,
		'F': S,
	},
	W: {
		'-': E,
		'J': N,
		'7': S,
	},
	S: {
		'|': N,
		'7': W,
		'F': E,
	},
	N: {
		'|': S,
		'J': W,
		'L': E,
	},
}

func invert(dir Pair) Pair {
	switch dir {
	case E:
		return W
	case W:
		return E
	case S:
		return N
	case N:
		return S
	default:
		panic("foo")
	}
}

func parse() [][]byte {
	input := lo.Must(os.ReadFile("./input"))
	grid := bytes.Split(input, []byte{'\n'})
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	return grid
}

func loc(grid [][]byte, b byte) (int, int) {
	for i, r := range grid {
		for j, c := range r {
			if c == b {
				return i, j
			}
		}
	}
	panic("foo")
}

func valid(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func dfs(grid [][]byte, i, j int, dir Pair, path *[]Pair) bool {
	var dirs []Pair
	if dir == _dummy {
		dirs = _dirs
	} else {
		dirs = []Pair{dir}
	}

	for _, d := range dirs {
		ni, nj := i+d.A, j+d.B
		if valid(grid, ni, nj) {
			if grid[ni][nj] == 'S' {
				return true
			}

			newdir, ok := _m[invert(d)][grid[ni][nj]]
			*path = append(*path, lo.T2(ni, nj))
			if ok && dfs(grid, ni, nj, newdir, path) {
				return true
			}
			*path = (*path)[:len(*path)-1]
		}
	}

	return false
}

func main() {
	grid := parse()

	si, sj := loc(grid, 'S')
	var path []Pair
	_ = dfs(grid, si, sj, _dummy, &path)

	fmt.Println(math.Ceil(float64(len(path)) / 2))
}
