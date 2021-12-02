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
		switch p[0] {
		case "forward":
			h += v
			d += (a * v)
		case "down":
			a += v
		case "up":
			a -= v
		}
	}
	fmt.Printf("hPos: %d\tvPos: %d\tproduct: %d\n", h, d, h*d)
}
