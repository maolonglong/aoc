package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/samber/lo"
)

func gcd(a, b int) int {
	var tmp int
	for b != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}

func main() {
	f := lo.Must(os.Open("./input"))
	defer f.Close()

	r := bufio.NewReader(f)

	instructions := lo.Must(r.ReadBytes('\n'))
	instructions = instructions[:len(instructions)-1]
	_ = lo.Must(r.ReadBytes('\n'))
	m := make(map[string][]string)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		from, to, _ := strings.Cut(line, " = ")
		to = to[1 : len(to)-1]
		left, right, _ := strings.Cut(to, ", ")

		m[from] = []string{left, right}
	}

	part1 := func(start string) int {
		var i, step int
		curr := start
		for {
			if curr[len(curr)-1] == 'Z' {
				break
			}

			inst := instructions[i]
			step++
			i++
			if i == len(instructions) {
				i = 0
			}

			if inst == 'L' {
				curr = m[curr][0]
			} else {
				curr = m[curr][1]
			}
		}

		return step
	}

	nodes := lo.Keys(m)
	startNodes := lo.Filter(nodes, func(item string, _ int) bool {
		return strings.HasSuffix(item, "A")
	})

	steps := lo.Map(startNodes, func(item string, _ int) int {
		return part1(item)
	})

	res := lo.Reduce(steps, func(agg, item, _ int) int {
		return agg / gcd(agg, item) * item
	}, 1)

	fmt.Println(res)
}
