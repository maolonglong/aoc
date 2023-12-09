package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	f := lo.Must(os.Open("./input"))
	defer f.Close()

	var sum int

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		items := strings.Fields(line)
		var seq [][]int
		var n int
		seq = append(seq, make([]int, 0, len(items)))
		for _, it := range items {
			x, _ := strconv.Atoi(it)
			seq[0] = append(seq[0], x)
		}

		for {
			seq = append(seq, make([]int, len(seq[n])-1))
			for i := 0; i+1 < len(seq[n]); i++ {
				seq[n+1][i] = seq[n][i+1] - seq[n][i]
			}
			n++

			if lo.EveryBy(seq[n], func(item int) bool { return item == 0 }) {
				break
			}
		}

		var end int // 0
		for i := n - 1; i >= 0; i-- {
			end = seq[i][len(seq[i])-1] + end
		}
		sum += end
	}

	fmt.Println(sum)
}
