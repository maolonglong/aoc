package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/samber/lo"
)

func parse() [][]byte {
	input := lo.Must(os.ReadFile("./input"))
	grid := bytes.Split(input, []byte{'\n'})
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	return grid
}

func main() {
	grid := parse()

	m := len(grid)
	n := len(grid[0])

	flags := make([]int, n)
	for j := range flags {
		flags[j] = -1
	}

	for i, r := range grid {
		for j, c := range r {
			switch c {
			case '.':
				continue
			case '#':
				flags[j] = i
			default:
				if i-flags[j] > 1 {
					flags[j]++
					grid[i][j] = '.'
					grid[flags[j]][j] = 'O'
				} else {
					flags[j] = i
				}
			}
		}
	}

	var sum int

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if grid[i][j] == 'O' {
				sum += m - i
			}
		}
	}

	fmt.Println(sum)
}
