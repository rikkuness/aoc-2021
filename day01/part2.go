package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

func part2(f io.Reader) {
	s := bufio.NewScanner(f)
	var (
		w       [3]int
		t, i, l int = 0, 0, int(math.Inf(1))
	)
	for s.Scan() {
		c, _ := strconv.Atoi(s.Text())
		w[2], w[1], w[0] = w[1], w[0], c
		i++
		if i < 3 {
			continue
		}
		a := w[0] + w[1] + w[2]
		if a > l {
			t++
		}
		l = a
	}
	fmt.Println(t)
}
