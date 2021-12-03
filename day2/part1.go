package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func part1(f io.Reader) {
	s := bufio.NewScanner(f)
	var h, d int
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		v, _ := strconv.Atoi(p[1])
		switch p[0][0] {
		case 'f':
			h += v
		case 'd':
			d += v
		case 'u':
			d -= v
		}
	}
	fmt.Println(h * d)
}
