package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/thoas/go-funk"
)

func parse() [][][]byte {
	input := lo.Must(os.ReadFile("./input"))
	lines := bytes.Split(input, []byte{'\n'})
	grids := [][][]byte{{}}
	for i, l := range lines {
		if len(l) == 0 {
			if i+1 != len(lines) {
				grids = append(grids, [][]byte{})
			}
			continue
		}
		grids[len(grids)-1] = append(grids[len(grids)-1], l)
	}
	return grids
}

func findMirror(a []int64) []int {
	n := len(a)
	var res []int
	for i := 0; i+1 < n; i++ {
		l, r := i, i+1
		for l >= 0 && r < n && a[l] == a[r] {
			l--
			r++
		}
		if l == -1 || r == n {
			res = append(res, l+1+(r-l-1)/2)
		}
	}
	return res
}

func summarize(grid [][]byte) []int {
	m := len(grid)
	n := len(grid[0])

	rows := make([]int64, m)

	for i, r := range grid {
		s := string(r)
		s = strings.ReplaceAll(s, ".", "0")
		s = strings.ReplaceAll(s, "#", "1")
		rows[i], _ = strconv.ParseInt(s, 2, 64)
	}

	var res []int
	res = append(res, lo.Map(findMirror(rows), func(x, _ int) int { return x * 100 })...)

	cols := make([]int64, n)
	b := make([]byte, m)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			b[i] = grid[i][j]
		}
		s := string(b)
		s = strings.ReplaceAll(s, ".", "0")
		s = strings.ReplaceAll(s, "#", "1")
		cols[j], _ = strconv.ParseInt(s, 2, 64)
	}

	res = append(res, findMirror(cols)...)
	return res
}

func invert(b byte) byte {
	if b == '.' {
		return '#'
	}
	return '.'
}

func diff(p1, p2 []int) int {
	_, v := funk.Difference(p1, p2)
	res := v.([]int)
	if len(res) > 0 {
		return res[0]
	}
	return 0
}

func main() {
	grids := parse()
	var sum int
loop:
	for _, g := range grids {
		part1 := summarize(g)
		for i, r := range g {
			for j, c := range r {
				g[i][j] = invert(c)
				if part2 := summarize(g); len(part2) > 0 {
					if v := diff(part1, part2); v != 0 {
						sum += v
						continue loop
					}
				}
				g[i][j] = c
			}
		}
	}
	fmt.Println(sum)
}
