package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type rule struct {
	op   byte
	arg1 byte
	arg2 int
	next string
}

type workflow struct {
	// name  string
	rules []rule
}

func parse() (map[string]workflow, []map[byte]int) {
	input := strings.TrimSpace(string(lo.Must(os.ReadFile("./input"))))
	top, bottom, _ := strings.Cut(input, "\n\n")

	workflows := make(map[string]workflow)
	for _, s := range strings.Split(top, "\n") {
		name, rest, _ := strings.Cut(s, "{")
		var rules []rule
		for _, ss := range strings.Split(rest[:len(rest)-1], ",") {
			if idx := strings.IndexByte(ss, ':'); idx != -1 {
				num, _ := strconv.Atoi(ss[2:idx])
				rules = append(rules, rule{
					op:   ss[1],
					arg1: ss[0],
					arg2: num,
					next: ss[idx+1:],
				})
			} else {
				// final
				rules = append(rules, rule{
					next: ss,
				})
			}
		}
		workflows[name] = workflow{
			rules: rules,
		}
	}

	var parts []map[byte]int
	for _, s := range strings.Split(bottom, "\n") {
		part := make(map[byte]int)
		for _, ss := range strings.Split(s[1:len(s)-1], ",") {
			name, val, _ := strings.Cut(ss, "=")
			if len(name) != 1 {
				panic("foo")
			}
			num, _ := strconv.Atoi(val)
			part[name[0]] = num
		}
		parts = append(parts, part)
	}
	return workflows, parts
}

func main() {
	workflows, parts := parse()

	var passed []map[byte]int

	for _, part := range parts {
		curr := "in"

	nextWorkflow:
		for {
			if curr == "A" {
				passed = append(passed, part)
				break
			}
			if curr == "R" {
				break
			}
			w := workflows[curr]
			for _, r := range w.rules {
				switch r.op {
				case '<':
					if part[r.arg1] < r.arg2 {
						curr = r.next
						continue nextWorkflow
					}
				case '>':
					if part[r.arg1] > r.arg2 {
						curr = r.next
						continue nextWorkflow
					}
				default:
					curr = r.next
					continue nextWorkflow
				}
			}
		}
	}

	var sum int
	for _, part := range passed {
		for _, v := range part {
			sum += v
		}
	}
	fmt.Println(sum)
}
