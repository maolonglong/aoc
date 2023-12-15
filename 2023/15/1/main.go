package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/samber/lo"
)

func hash(data []byte) int {
	var h int
	for _, b := range data {
		h += int(b)
		h *= 17
		h %= 256
	}
	return h
}

func main() {
	contents := lo.Must(os.ReadFile("./input"))
	steps := bytes.Split(bytes.TrimSpace(contents), []byte{','})
	var sum int
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Println(sum)
}
