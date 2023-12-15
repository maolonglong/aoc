package main

import (
	"bytes"
	"container/list"
	"fmt"
	"os"
	"strconv"

	"github.com/samber/lo"
)

type step struct {
	label []byte
	op    byte
	num   int
}

type lens struct {
	label []byte
	focal int
}

func hash(data []byte) int {
	var h int
	for _, b := range data {
		h += int(b)
		h *= 17
		h %= 256
	}
	return h
}

func parse() []step {
	contents := lo.Must(os.ReadFile("./input"))
	xs := bytes.Split(bytes.TrimSpace(contents), []byte{','})
	var input []step
	for _, x := range xs {
		l, r, ok := bytes.Cut(x, []byte{'='})
		if ok {
			num, _ := strconv.Atoi(string(r))
			input = append(input, step{l, '=', num})
		} else {
			input = append(input, step{x[:len(x)-1], '-', -1})
		}
	}
	return input
}

//lint:ignore U1000 _
func dump(hashmap []*list.List) {
	for i, box := range hashmap {
		if box == nil || box.Len() == 0 {
			continue
		}
		fmt.Printf("Box %d:", i)
		for p := box.Front(); p != nil; p = p.Next() {
			v := p.Value.(*lens)
			fmt.Printf(" [%s %v]", v.label, v.focal)
		}
		fmt.Println()
	}
}

func main() {
	steps := parse()

	hashmap := make([]*list.List, 256)
nextStep:
	for _, step := range steps {
		idx := hash(step.label)
		if hashmap[idx] == nil {
			hashmap[idx] = list.New()
		}
		box := hashmap[idx]
		if step.op == '=' {
			for p := box.Front(); p != nil; p = p.Next() {
				lens := p.Value.(*lens)
				if bytes.Equal(lens.label, step.label) {
					lens.focal = step.num
					continue nextStep
				}
			}
			box.PushBack(&lens{step.label, step.num})
		} else {
			for p := box.Front(); p != nil; p = p.Next() {
				lens := p.Value.(*lens)
				if bytes.Equal(lens.label, step.label) {
					box.Remove(p)
					continue nextStep
				}
			}
		}
	}

	// dump(hashmap)

	var sum int
	for i, box := range hashmap {
		if box == nil || box.Len() == 0 {
			continue
		}
		var j int
		for p := box.Front(); p != nil; p = p.Next() {
			j++
			sum += (i + 1) * j * p.Value.(*lens).focal
		}
	}
	fmt.Println(sum)
}
