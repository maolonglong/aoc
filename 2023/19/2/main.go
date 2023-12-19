package main

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type rule struct {
	op   func(int, int) bool
	arg1 byte
	arg2 int
	next string
}

type workflow struct {
	// name  string
	rules []rule
}

func _range(start, end int) []int {
	a := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		a[i-start] = i
	}
	return a
}

func lt(x, y int) bool {
	return x < y
}

func gt(x, y int) bool {
	return x > y
}

func parse() map[string]workflow {
	input := strings.TrimSpace(string(lo.Must(os.ReadFile("./input"))))
	top, _, _ := strings.Cut(input, "\n\n")

	workflows := make(map[string]workflow)
	for _, s := range strings.Split(top, "\n") {
		name, rest, _ := strings.Cut(s, "{")
		var rules []rule
		for _, ss := range strings.Split(rest[:len(rest)-1], ",") {
			if idx := strings.IndexByte(ss, ':'); idx != -1 {
				num, _ := strconv.Atoi(ss[2:idx])
				rules = append(rules, rule{
					op:   lo.If(ss[1] == '>', gt).Else(lt),
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

	return workflows
}

func dfs(workflows map[string]workflow, wName string, rIdx int, part map[byte][]int) int {
	if wName == "A" {
		ret := 1
		for _, v := range part {
			ret *= len(v)
		}
		return ret
	}
	if wName == "R" {
		return 0
	}

	r := workflows[wName].rules[rIdx]
	if r.op == nil {
		return dfs(workflows, r.next, 0, part)
	}
	part1 := maps.Clone(part)
	part2 := maps.Clone(part)
	k := r.arg1
	values := part[k]
	var succ []int
	var fail []int
	for _, v := range values {
		if r.op(v, r.arg2) {
			succ = append(succ, v)
		} else {
			fail = append(fail, v)
		}
	}
	part1[k] = succ
	part2[k] = fail

	return dfs(workflows, r.next, 0, part1) +
		dfs(workflows, wName, rIdx+1, part2)
}

func main() {
	workflows := parse()
	a := _range(1, 4000)
	parts := map[byte][]int{
		'x': a,
		'm': a,
		'a': a,
		's': a,
	}

	fmt.Println(dfs(workflows, "in", 0, parts))
}
