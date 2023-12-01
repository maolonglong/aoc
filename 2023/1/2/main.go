package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/samber/lo"
)

var (
	first int
	last  int
	flag  bool
	sum   int
)

func hit(b byte) {
	if !flag {
		flag = true
		first = int(b - '0')
		last = int(b - '0')
	} else {
		last = int(b - '0')
	}
}

func main() {
	f := lo.Must(os.Open("./input"))
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		n := len(line)
		for i, b := range line {
			if b >= '0' && b <= '9' {
				hit(b)
			} else {
				switch b {
				case 'o':
					if i+2 < n && line[i+1] == 'n' && line[i+2] == 'e' {
						hit('1')
					}
				case 't':
					if i+2 < n && line[i+1] == 'w' && line[i+2] == 'o' {
						hit('2')
					} else if i+4 < n && line[i+1] == 'h' && line[i+2] == 'r' &&
						line[i+3] == 'e' && line[i+4] == 'e' {
						hit('3')
					}
				case 'f':
					if i+3 < n && line[i+1] == 'o' && line[i+2] == 'u' &&
						line[i+3] == 'r' {
						hit('4')
					} else if i+3 < n && line[i+1] == 'i' && line[i+2] == 'v' &&
						line[i+3] == 'e' {
						hit('5')
					}
				case 's':
					if i+2 < n && line[i+1] == 'i' && line[i+2] == 'x' {
						hit('6')
					} else if i+4 < n && line[i+1] == 'e' && line[i+2] == 'v' &&
						line[i+3] == 'e' && line[i+4] == 'n' {
						hit('7')
					}
				case 'e':
					if i+4 < n && line[i+1] == 'i' && line[i+2] == 'g' &&
						line[i+3] == 'h' && line[i+4] == 't' {
						hit('8')
					}
				case 'n':
					if i+3 < n && line[i+1] == 'i' && line[i+2] == 'n' &&
						line[i+3] == 'e' {
						hit('9')
					}
				case '\n':
					sum += first*10 + last
					flag = false
				default:
				}
			}
		}
	}

	fmt.Println(sum)
}
