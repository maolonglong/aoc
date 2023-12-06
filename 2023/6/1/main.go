package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func check(t, d, hold int) bool {
	cost := int(math.Ceil((float64(d) / float64(hold))))
	if cost*hold == d {
		cost++
	}
	return cost+hold <= t
}

func possible(t, d int) int {
	var cnt int
	for i := 1; i < t; i++ {
		if check(t, d, i) {
			cnt++
		}
	}
	return cnt
}

func main() {
	input := string(lo.Must(os.ReadFile("./input")))
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

	// fmt.Printf("times: %v\n", times)
	// fmt.Printf("distance: %v\n", distance)

	res := 1
	for i := 0; i < len(times); i++ {
		res *= possible(times[i], distance[i])
	}
	fmt.Println(res)
}
