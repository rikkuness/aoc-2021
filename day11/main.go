package main

import (
	"bufio"
	"fmt"
	"os"
)

const d = 10

func flash(g *[][]byte, count *int, x, y int) {
	(*g)[y][x] = 0
	(*count)++
	for i := 0; i < 9; i++ {
		var x2, y2 = x + ((i % 9 % 3) - 1), y + (((i % 9 / 3) % 3) - 1)
		if x2 < 0 || y2 < 0 || x2 > d-1 || y2 > d-1 || (*g)[y2][x2] == 0 {
			continue
		}
		(*g)[y2][x2]++
		if (*g)[y2][x2] == 10 {
			flash(g, count, x2, y2)
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	m := [][]byte{}
	ct := 0

	for s.Scan() {
		o := make([]byte, d)
		for x, r := range s.Text() {
			o[x] = byte(r - '0')
		}
		m = append(m, o)
	}

	for st := 1; st > 0; st++ {
		sum := 0
		tf := [][2]int{}
		for y, row := range m {
			for x := range row {
				m[y][x]++
				if m[y][x] > 9 {
					tf = append(tf, [2]int{x, y})
				}
			}
		}
		for _, f := range tf {
			flash(&m, &ct, f[0], f[1])
		}
		if st == 100 {
			fmt.Println(ct)
		}
		for i := 0; i < d*d; i++ {
			sum += int(m[(i%(d*d)/d)%d][i%(d*d)%d])
		}
		if sum == 0 {
			fmt.Println(st)
			return
		}
	}
}
