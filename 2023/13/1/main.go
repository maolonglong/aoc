package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
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

func findMirror(a []int64) int {
	n := len(a)
	for i := 0; i+1 < n; i++ {
		l, r := i, i+1
		for l >= 0 && r < n && a[l] == a[r] {
			l--
			r++
		}
		if l == -1 || r == n {
			return l + 1 + (r-l-1)/2
		}
	}
	return 0
}

func summarize(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	rows := make([]int64, m)

	for i, r := range grid {
		s := string(r)
		s = strings.ReplaceAll(s, ".", "0")
		s = strings.ReplaceAll(s, "#", "1")
		rows[i], _ = strconv.ParseInt(s, 2, 64)
	}
	res0 := findMirror(rows)
	if res0 != 0 {
		return res0 * 100
	}

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

	return findMirror(cols)
}

func main() {
	grids := parse()
	var sum int
	for _, g := range grids {
		sum += summarize(g)
	}
	fmt.Println(sum)
}
