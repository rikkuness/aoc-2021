package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func part1() {
	f, _ := os.Open("input.txt")
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
