package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

func part1(f io.Reader) {
	s := bufio.NewScanner(f)
	var c, l, t int = 0, int(math.Inf(1)), 0
	for s.Scan() {
		c, _ = strconv.Atoi(s.Text())
		if c > l {
			t++
		}
		l = c
	}
	fmt.Println(t)
}
