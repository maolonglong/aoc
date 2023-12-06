package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func check(t, d, hold int) (bool, int) {
	cost := int(math.Ceil((float64(d) / float64(hold))))
	if cost*hold == d {
		cost++
	}
	return cost+hold <= t, hold - cost
}

func searchRange(t, d int) (int, int) {
	r0, r1 := -1, -1

	// var loops int

	left, right := 0, t-1
	for left+1 < right {
		// loops++
		mid := int(uint(left+right) >> 1)
		ok, cmp := check(t, d, mid)
		if ok {
			right = mid
		} else {
			if cmp > 0 {
				right = mid
			} else {
				left = mid
			}
		}
	}
	if ok, _ := check(t, d, left); ok {
		r0 = left
	} else if ok, _ := check(t, d, right); ok {
		r0 = right
	} else {
		return -1, -1
	}

	left, right = 0, t-1
	for left+1 < right {
		// loops++
		mid := int(uint(left+right) >> 1)
		ok, cmp := check(t, d, mid)
		if ok {
			left = mid
		} else {
			if cmp > 0 {
				right = mid
			} else {
				left = mid
			}
		}
	}
	if ok, _ := check(t, d, right); ok {
		r1 = right
	} else if ok, _ := check(t, d, left); ok {
		r1 = left
	}

	// fmt.Println(loops) // 50
	return r0, r1
}

func main() {
	input := string(lo.Must(os.ReadFile("./input")))
	input = strings.ReplaceAll(input, " ", "") // part2
	lines := strings.Split(input, "\n")

	var times []int
	for _, t := range strings.Fields(strings.TrimPrefix(lines[0], "Time:")) {
		x, _ := strconv.Atoi(t)
		times = append(times, x)
	}

	var distance []int
	for _, d := range strings.Fields(strings.TrimPrefix(lines[1], "Distance:")) {
		x, _ := strconv.Atoi(d)
		distance = append(distance, x)
	}

	l, r := searchRange(times[0], distance[0])
	fmt.Println(r - l + 1)
}
