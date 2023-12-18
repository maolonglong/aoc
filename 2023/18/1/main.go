package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type plan struct {
	dir   lo.Tuple2[int, int]
	n     int
	color string
}

func parse() []plan {
	input := string(lo.Must(os.ReadFile("./input")))
	lines := strings.Split(input, "\n")
	var plans []plan
	for _, line := range lines {
		if line == "" {
			continue
		}
		a := strings.Fields(line)
		var p plan
		p.n, _ = strconv.Atoi(a[1])
		p.color = strings.Trim(a[2], "()")
		switch a[0][0] {
		case 'U':
			p.dir = lo.T2(-1, 0)
		case 'D':
			p.dir = lo.T2(1, 0)
		case 'L':
			p.dir = lo.T2(0, -1)
		case 'R':
			p.dir = lo.T2(0, 1)
		default:
			panic("foo")
		}
		plans = append(plans, p)
	}
	return plans
}

func inPoly(poly []lo.Tuple2[int, int], i, j int) bool {
	minI := poly[0].A
	minJ := poly[0].B
	maxI := poly[0].A
	maxJ := poly[0].B

	for _, p := range poly[1:] {
		minI = min(minI, p.A)
		minJ = min(minJ, p.B)
		maxI = max(maxI, p.A)
		maxJ = max(maxJ, p.B)
	}

	if i < minI || i > maxI || j < minJ || j > maxJ {
		return false
	}

	var res bool
	for k := 0; k < len(poly); k++ {
		p1 := poly[k]
		p2 := poly[(k+1)%len(poly)]
		p3 := poly[(k+2)%len(poly)]
		p4 := poly[(k+3)%len(poly)]

		// edge
		if i == p1.A && i == p2.A && j >= min(p1.B, p2.B) && j <= max(p1.B, p2.B) ||
			j == p1.B && j == p2.B && i >= min(p1.A, p2.A) && i <= max(p1.A, p2.A) {
			return true
		}

		if p1.A != i && p2.A != i && (p1.A > i) != (p2.A > i) &&
			j < (p2.B-p1.B)*(i-p1.A)/(p2.A-p1.A)+p1.B {
			res = !res
			continue
		}

		if p1.A != i && p2.A == i && p2.B > j {
			if p3.A == i && p3.B > j {
				if p1.A != i && p4.A != i && (p1.A > i) != (p4.A > i) {
					res = !res
					continue
				}
			} else {
				if p1.A != i && p3.A != i && (p1.A > i) != (p3.A > i) {
					res = !res
					continue
				}
			}
		}
	}

	return res
}

func main() {
	plans := parse()

	curr := lo.T2(0, 0)
	// poly := []lo.Tuple2[int, int]{curr}
	var poly []lo.Tuple2[int, int]

	for _, plan := range plans {
		next := lo.T2(curr.A+plan.dir.A*plan.n,
			curr.B+plan.dir.B*plan.n)
		poly = append(poly, next)
		curr = next
	}

	var minI, minJ, maxI, maxJ int
	for _, p := range poly[1:] {
		minI = min(minI, p.A)
		minJ = min(minJ, p.B)
		maxI = max(maxI, p.A)
		maxJ = max(maxJ, p.B)
	}

	// fixup
	if minI < 0 || minJ < 0 {
		offsetI := lo.If(minI < 0, -minI).Else(0)
		offsetJ := lo.If(minJ < 0, -minJ).Else(0)

		for i := range poly {
			poly[i] = lo.T2(poly[i].A+offsetI, poly[i].B+offsetJ)
		}

		minI += offsetI
		minJ += offsetJ
		maxI += offsetI
		maxJ += offsetJ
	}

	var sum int
	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if inPoly(poly, i, j) {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
