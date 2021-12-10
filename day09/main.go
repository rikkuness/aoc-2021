package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var m [][]byte
	for y := 0; s.Scan(); y++ {
		row := s.Text()
		o := make([]byte, len(row))
		for x, r := range row {
			o[x] = byte(r - '0')
		}
		m = append(m, o)
	}

	part1(m)
	part2(m)
}
