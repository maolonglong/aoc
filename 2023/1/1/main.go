package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/samber/lo"
)

func main() {
	f := lo.Must(os.Open("./input"))
	defer f.Close()

	r := bufio.NewReader(f)

	var (
		first byte
		last  byte
		flag  bool
		sum   int
	)
	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if b >= '0' && b <= '9' {
			if !flag {
				flag = true
				first = b
				last = b
			} else {
				last = b
			}
		} else if b == '\n' {
			sum += int(first-'0')*10 + int(last-'0')
			flag = false
		}
	}

	fmt.Println(sum)
}
