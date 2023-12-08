package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/samber/lo"
)

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

	var i, steps int
	curr := "AAA"
	for {
		if curr == "ZZZ" {
			break
		}

		inst := instructions[i]
		steps++
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

	fmt.Println(steps)
}
