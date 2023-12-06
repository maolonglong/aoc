package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

func main() {
	f := lo.Must(os.Open("./input"))
	defer f.Close()

	r := bufio.NewReader(f)
	st := mapset.NewThreadUnsafeSet[int]()
	var sum int
	for {
		st.Clear()

		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		l := parseLine(line)
		for _, x := range l.Win {
			st.Add(x)
		}
		var score int
		for _, x := range l.Have {
			if st.Contains(x) {
				score++
			}
		}
		if score > 0 {
			sum += 1 << (score - 1)
		}
	}

	fmt.Println(sum)
}

type Line struct {
	Win  []int
	Have []int
	Card int
}

var _re = regexp.MustCompile(`Card\s+(\d+):\s+(.*?)\|(.*)`)

func parseLine(line []byte) Line {
	matched := _re.FindSubmatch(line)
	var l Line
	l.Card, _ = strconv.Atoi(string(matched[1]))

	for _, x := range strings.Fields(string(matched[2])) {
		num, _ := strconv.Atoi(x)
		l.Win = append(l.Win, num)
	}

	for _, x := range strings.Fields(string(matched[3])) {
		num, _ := strconv.Atoi(x)
		l.Have = append(l.Have, num)
	}

	return l
}
