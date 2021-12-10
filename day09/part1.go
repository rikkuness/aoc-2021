package main

import (
	"fmt"
	"math"
)

func valAt(g *[][]byte, y, x int) int {
	if y < 0 || y >= len(*g) || x < 0 || x >= len((*g)[0]) {
		return int(math.Inf(1))
	}
	return int((*g)[y][x])
}

func isLow(g *[][]byte, y, x int) bool {
	v := valAt(g, y, x)
	return valAt(g, y+1, x) > v && valAt(g, y-1, x) > v && valAt(g, y, x+1) > v && valAt(g, y, x-1) > v
}

func part1(g [][]byte) {
	var risk int
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if isLow(&g, y, x) {
				risk += valAt(&g, y, x) + 1
			}
		}
	}
	fmt.Println(risk)
}
