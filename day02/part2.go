package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func part2(f io.Reader) {
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
