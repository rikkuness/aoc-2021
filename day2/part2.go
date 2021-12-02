package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var h, d, a int
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		v, _ := strconv.Atoi(p[1])
		switch p[0][0] {
		case 'f':
			h += v
			d += (a * v)
		case 'd':
			a += v
		case 'u':
			a -= v
		}
	}
	fmt.Println(h * d)
}
