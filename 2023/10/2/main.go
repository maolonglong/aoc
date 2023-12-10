package main

import (
	"bytes"
	"fmt"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
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

func dfs(grid [][]byte, i, j int, dir Pair, path, poly *[]Pair) bool {
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
			if newdir != d {
				*poly = append(*poly, lo.T2(ni, nj))
			}
			if ok && dfs(grid, ni, nj, newdir, path, poly) {
				return true
			}
			*path = (*path)[:len(*path)-1]
			if newdir != d {
				*poly = (*poly)[:len(*poly)-1]
			}
		}
	}

	return false
}

func inPoly(poly []Pair, i, j int) bool {
	minI := poly[0].A
	minJ := poly[0].B
	maxI := poly[0].A
	maxJ := poly[0].B

	for _, p := range poly[1:] {
		minI = min(minI, p.A)
		minJ = min(minJ, p.B)
		maxI = max(maxI, p.A)
		maxJ = max(maxJ, p.B)
	}

	if i < minI || i > maxI || j < minJ || j > maxJ {
		return false
	}

	var res bool
	for k := 0; k < len(poly); k++ {
		p1 := poly[k]
		p2 := poly[(k+1)%len(poly)]
		p3 := poly[(k+2)%len(poly)]
		p4 := poly[(k+3)%len(poly)]

		if p1.A != i && p2.A != i && (p1.A > i) != (p2.A > i) &&
			j < (p2.B-p1.B)*(i-p1.A)/(p2.A-p1.A)+p1.B {
			res = !res
			continue
		}

		if p1.A != i && p2.A == i && p2.B > j {
			if p3.A == i && p3.B > j {
				if p1.A != i && p4.A != i && (p1.A > i) != (p4.A > i) {
					res = !res
					continue
				}
			} else {
				if p1.A != i && p3.A != i && (p1.A > i) != (p3.A > i) {
					res = !res
					continue
				}
			}
		}
	}

	return res
}

func main() {
	grid := parse()

	si, sj := loc(grid, 'S')
	poly := []Pair{lo.T2(si, sj)}
	path := []Pair{lo.T2(si, sj)}
	_ = dfs(grid, si, sj, _dummy, &path, &poly)

	vis := mapset.NewThreadUnsafeSet(path...)

	var sum int
	for i, r := range grid {
		for j := range r {
			if vis.Contains(lo.T2(i, j)) {
				continue
			}
			if inPoly(poly, i, j) {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
