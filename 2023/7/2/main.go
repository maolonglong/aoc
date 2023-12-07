package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

var _handStrength = map[byte]int{
	'J': -1,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 11,
	'K': 12,
	'A': 13,
}

const (
	_ = iota
	High
	One
	Two
	Three
	Full
	Four
	Five
)

func handType(hand []byte) int {
	cnt := make(map[byte]int)
	for _, b := range hand {
		cnt[b]++
	}
	j := cnt['J']
	freq := make(map[int]int)
	for _, v := range cnt {
		freq[v]++
	}
	if freq[5] == 1 {
		return Five
	}
	if freq[4] == 1 {
		if j > 0 {
			return Five
		}
		return Four
	}
	if freq[3] == 1 {
		if j == 2 || j == 3 && freq[2] == 1 {
			return Five
		}
		if j == 1 || j == 3 {
			return Four
		}
		if freq[2] == 1 {
			return Full
		}
		return Three
	}
	if freq[2] == 2 {
		if j == 2 {
			return Four
		}
		if j == 1 {
			return Full
		}
		return Two
	}
	if freq[2] == 1 {
		if j > 0 {
			return Three
		}
		return One
	}
	if j == 1 {
		return One
	}
	return High
}

func less(h1, h2 []byte) bool {
	for i := 0; i < 5; i++ {
		if _handStrength[h1[i]] != _handStrength[h2[i]] {
			return _handStrength[h1[i]] < _handStrength[h2[i]]
		}
	}
	return false
}

type Hand struct {
	b   []byte
	typ int
	bid int
}

type Hands []Hand

func (a Hands) Len() int      { return len(a) }
func (a Hands) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Hands) Less(i, j int) bool {
	if a[i].typ != a[j].typ {
		return a[i].typ < a[j].typ
	}
	return less(a[i].b, a[j].b)
}

func main() {
	f := lo.Must(os.Open("./input"))
	defer f.Close()

	var hands Hands

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// (hand, bid)
		pair := bytes.Fields(line)
		bid, _ := strconv.Atoi(string(pair[1]))
		hands = append(hands, Hand{
			b:   pair[0],
			typ: handType(pair[0]),
			bid: bid,
		})
	}

	sort.Sort(hands)

	var sum int
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}
	fmt.Println(sum)
}
