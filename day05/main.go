package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const dim = 1000

func calculateMoves(move string) (moves [][2]int) {
	var c [4]float64
	for i := 1; i < 5; i++ {
		p, _ := strconv.Atoi(
			regexp.MustCompile(`^(\d*),(\d*) -> (\d*),(\d*)$`).FindStringSubmatch(move)[i])
		c[i-1] = float64(p)
	}
	var yU, yL, xU, xL = math.Max(c[1], c[3]),
		math.Min(c[1], c[3]),
		math.Max(c[0], c[2]),
		math.Min(c[0], c[2])
	var mC = int(yU - yL)
	if mC == 0 {
		mC = int(xU - xL)
	}
	for i := 0; i <= mC; i++ {
		var x, y int
		switch c[0] {
		case c[2]:
			x = int(xL)
		case xU:
			x = int(xU) - i
		case xL:
			x = int(xL) + i
		}
		switch c[1] {
		case c[3]:
			y = int(yL)
		case yU:
			y = int(yU) - i
		case yL:
			y = int(yL) + i
		}
		moves = append(moves, [2]int{x, y})
	}
	return
}

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var m [dim][dim]int
	for s.Scan() {
		for _, cc := range calculateMoves(s.Text()) {
			m[cc[0]][cc[1]]++
		}
	}
	var c int
	for i := 0; i < dim; i++ {
		for _, n := range m[i] {
			if n > 1 {
				c++
			}
		}
	}
	fmt.Println(c)
}
