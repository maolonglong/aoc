package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/samber/lo"
)

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSym(b byte) bool {
	return !isNum(b) && b != '.'
}

func main() {
	input := bytes.Split(lo.Must(os.ReadFile("./input")), []byte{'\n'})
	input = input[:len(input)-1]
	m := len(input)
	n := len(input[0])
	// fmt.Printf("m: %v\n", m)
	// fmt.Printf("n: %v\n", n)

	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	isPart := func(i, j int) bool {
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && isSym(input[ni][nj]) {
				return true
			}
		}
		return false
	}

	maybe := make(map[int]map[int]int)
	findGear := func(i, j int) []int {
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && input[ni][nj] == '*' {
				return []int{ni, nj}
			}
		}
		return nil
	}

	var (
		// 0: init
		// 1: number
		// 2: part number
		state int

		pos []int // '*' pos
		num int
		sum int
	)

	for i, line := range input {
		// reset
		state = 0
		num = 0
		pos = nil

		for j, b := range line {
			switch state {
			case 0:
				if isNum(b) {
					num = int(b - '0')
					if isPart(i, j) {
						state = 2
						pos = findGear(i, j)
					} else {
						state = 1
					}
				}
			case 1:
				if isNum(b) {
					num = num*10 + int(b-'0')
					if isPart(i, j) {
						if pos == nil {
							pos = findGear(i, j)
						}
						state = 2
					}
				} else {
					state = 0
					num = 0
					pos = nil
				}
			case 2:
				if isNum(b) {
					num = num*10 + int(b-'0')
					if pos == nil {
						pos = findGear(i, j)
					}
				}
				if !isNum(b) || j == n-1 {
					if pos != nil {
						_, ok := maybe[pos[0]]
						if !ok {
							maybe[pos[0]] = make(map[int]int)
						}
						if maybe[pos[0]][pos[1]] > 0 {
							sum += maybe[pos[0]][pos[1]] * num
						} else {
							maybe[pos[0]][pos[1]] = num
						}
					}
					state = 0
					num = 0
					pos = nil
				}
			}
		}
	}
	fmt.Println(sum)
}
