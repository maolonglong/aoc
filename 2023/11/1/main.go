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

	flagI := make([]bool, m)
	flagJ := make([]bool, n)

	var galaxies []lo.Tuple2[int, int]
	for i, r := range grid {
		for j, c := range r {
			if c == '#' {
				flagI[i] = true
				flagJ[j] = true
				galaxies = append(galaxies, lo.T2(i, j))
			}
		}
	}

	var sum int
	k := len(galaxies)
	for i := 0; i < k-1; i++ {
		for j := i + 1; j < k; j++ {
			minI := min(galaxies[i].A, galaxies[j].A)
			maxI := max(galaxies[i].A, galaxies[j].A)
			minJ := min(galaxies[i].B, galaxies[j].B)
			maxJ := max(galaxies[i].B, galaxies[j].B)

			path := (maxI - minI) + (maxJ - minJ)
			for expand := minI; expand <= maxI; expand++ {
				if !flagI[expand] {
					path++
				}
			}
			for expand := minJ; expand <= maxJ; expand++ {
				if !flagJ[expand] {
					path++
				}
			}

			sum += path
		}
	}

	fmt.Println(sum)
}
