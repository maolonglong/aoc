package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
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

type node struct {
	loss int
	state
}

type state struct {
	pos   lo.Tuple2[int, int]
	steps int
	dir   lo.Tuple2[int, int]
}

var _ heap.Interface = (*nodes)(nil)

type nodes []node

func (a *nodes) Len() int           { return len(*a) }
func (a *nodes) Swap(i, j int)      { (*a)[i], (*a)[j] = (*a)[j], (*a)[i] }
func (a *nodes) Less(i, j int) bool { return (*a)[i].loss < (*a)[j].loss }
func (a *nodes) Pop() any           { x := (*a)[len(*a)-1]; (*a) = (*a)[:len(*a)-1]; return x }
func (a *nodes) Push(x any)         { *a = append(*a, x.(node)) }

func part1(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	dirs := []lo.Tuple2[int, int]{
		{A: -1, B: 0},
		{A: 1, B: 0},
		{A: 0, B: -1},
		{A: 0, B: 1},
	}

	var h nodes
	heap.Push(&h, node{
		loss: 0,
		state: state{
			pos:   lo.T2(0, 0),
			steps: 0,
			dir:   lo.T2(-1, -1), // dummy
		},
	})

	vis := mapset.NewThreadUnsafeSet[state]()
	for {
		curr := heap.Pop(&h).(node)
		if curr.pos.A == m-1 && curr.pos.B == n-1 && curr.steps >= 4 {
			return curr.loss
		}

		for _, nextDir := range dirs {
			if -nextDir.A == curr.dir.A && -nextDir.B == curr.dir.B {
				// reverse
				continue
			}

			nextPos := lo.T2(curr.pos.A+nextDir.A, curr.pos.B+nextDir.B)
			if nextPos.A < 0 || nextPos.A >= m || nextPos.B < 0 || nextPos.B >= n {
				continue
			}

			if curr.steps != 0 && nextDir != curr.dir && curr.steps < 4 {
				continue
			}

			nextState := state{
				pos:   nextPos,
				steps: 1,
				dir:   nextDir,
			}
			if nextDir == curr.dir {
				steps := curr.steps + 1
				if steps == 11 {
					continue
				}
				nextState.steps = steps
			}

			if vis.Contains(nextState) {
				continue
			}

			vis.Add(nextState)
			heap.Push(&h, node{
				loss:  curr.loss + int(grid[nextPos.A][nextPos.B]-'0'),
				state: nextState,
			})
		}
	}
}

func main() {
	grid := parse()
	fmt.Println(part1(grid))
}
