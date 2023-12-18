package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/samber/lo"
)

type plan struct {
	// color string
	dir lo.Tuple2[int, int]
	n   int64
}

func parse() []plan {
	input := string(lo.Must(os.ReadFile("./input.test")))
	lines := strings.Split(input, "\n")
	var plans []plan
	for i, line := range lines {
		if line == "" {
			continue
		}
		a := strings.Fields(line)
		var p plan
		p.n, _ = strconv.ParseInt(a[2][2:7], 16, 64)
		switch a[2][7] {
		case '3':
			p.dir = lo.T2(-1, 0)
		case '1':
			p.dir = lo.T2(1, 0)
		case '2':
			p.dir = lo.T2(0, -1)
		case '0':
			p.dir = lo.T2(0, 1)
		default:
			panic(i)
		}
		plans = append(plans, p)
	}
	return plans
}

func inPoly(poly []lo.Tuple2[int, int], i, j, minI, minJ, maxI, maxJ int) bool {
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
		next := lo.T2(curr.A+plan.dir.A*int(plan.n),
			curr.B+plan.dir.B*int(plan.n))
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

	// TODO:
	var sum atomic.Int64
	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if inPoly(poly, i, j, minI, minJ, maxI, maxJ) {
				sum.Add(1)
			}
		}
	}
	fmt.Println(sum.Load())
}
