package main

import (
	"fmt"
	"sort"
)

func dfs(g *[][]byte, size *int, y, x int) {
	if y < 0 || y >= len(*g) || x < 0 || x >= len((*g)[0]) || (*g)[y][x] >= 9 {
		return
	}
	(*size)++
	(*g)[y][x] = (*g)[y][x] + 128 // add a bit at the end to say we were here
	dfs(g, size, y+1, x)
	dfs(g, size, y-1, x)
	dfs(g, size, y, x+1)
	dfs(g, size, y, x-1)
}

func getIslands(g [][]byte) (islands []int) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if g[y][x] < 9 {
				var sz int
				dfs(&g, &sz, y, x)
				islands = append(islands, sz)
			}
		}
	}
	return
}

func part2(g [][]byte) {
	a := getIslands(g)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a[0] * a[1] * a[2])
}
