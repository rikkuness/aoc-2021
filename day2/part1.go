package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var h, d int
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		v, _ := strconv.Atoi(p[1])
		switch p[0] {
		case "forward":
			h += v
		case "down":
			d += v
		case "up":
			d -= v
		}
	}
	fmt.Printf("hPos: %d\tvPos: %d\tproduct: %d\n", h, d, h*d)
}
