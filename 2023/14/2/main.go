package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
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

func tilt(grid [][]byte, dir byte) {
	m := len(grid)
	n := len(grid[0])

	var (
		nflag      int
		initFlag   int
		vertical   bool
		is, ie, io int
		js, je, jo int
	)
	switch dir {
	case 'N':
		nflag = n
		initFlag = -1
		vertical = true
		is, ie, io = 0, m, 1
		js, je, jo = 0, n, 1
	case 'W':
		nflag = m
		initFlag = -1
		vertical = false
		is, ie, io = 0, n, 1
		js, je, jo = 0, m, 1
	case 'S':
		nflag = n
		initFlag = m
		vertical = true
		is, ie, io = m-1, -1, -1
		js, je, jo = 0, n, 1
	case 'E':
		nflag = m
		initFlag = n
		vertical = false
		is, ie, io = n-1, -1, -1
		js, je, jo = 0, m, 1
	}

	flags := make([]int, nflag)
	for i := range flags {
		flags[i] = initFlag
	}

	for i := is; i != ie; i += io {
		for j := js; j != je; j += jo {
			var c byte
			if vertical {
				c = grid[i][j]
			} else {
				c = grid[j][i]
			}
			switch c {
			case '.':
				continue
			case '#':
				flags[j] = i
			default:
				if vertical {
					if io == 1 {
						if i-flags[j] > 1 {
							flags[j]++
							grid[i][j] = '.'
							grid[flags[j]][j] = 'O'
						} else {
							flags[j] = i
						}
					} else {
						if flags[j]-i > 1 {
							flags[j]--
							grid[i][j] = '.'
							grid[flags[j]][j] = 'O'
						} else {
							flags[j] = i
						}
					}
				} else {
					if io == 1 {
						if i-flags[j] > 1 {
							flags[j]++
							grid[j][i] = '.'
							grid[j][flags[j]] = 'O'
						} else {
							flags[j] = i
						}
					} else {
						if flags[j]-i > 1 {
							flags[j]--
							grid[j][i] = '.'
							grid[j][flags[j]] = 'O'
						} else {
							flags[j] = i
						}
					}
				}
			}
		}
	}
}

func cycle(grid [][]byte) {
	tilt(grid, 'N')
	tilt(grid, 'W')
	tilt(grid, 'S')
	tilt(grid, 'E')
}

// unsafe
func hash(grid [][]byte) uint32 {
	b := make([]byte, 0, len(grid)*len(grid[0]))
	for _, r := range grid {
		b = append(b, r...)
	}
	return crc32.ChecksumIEEE(b)
}

func main() {
	grid := parse()
	m := len(grid)
	n := len(grid[0])

	cached := make(map[uint32]int)
	var cur int
	target := 1000000000
	cached[hash(grid)] = 0
	for {
		cycle(grid)
		cur++

		h := hash(grid)
		pre, ok := cached[h]
		if ok {
			period := cur - pre
			cur = cur + period*((target-cur)/period)
		} else {
			cached[h] = cur
		}

		if cur == target {
			break
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
