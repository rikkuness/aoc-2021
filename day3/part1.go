package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var (
		c [12]int
		x uint16
	)
	for s.Scan() {
		r := s.Text()
		for i := 0; i < len(r); i++ {
			if !(r[i]&1<<5 == 0) {
				c[i] += 1
			} else {
				c[i] -= 1
			}
		}
	}
	for i := 0; i < len(c); i++ {
		if c[i] > 0 {
			x |= 1 << (len(c) - 1 - i)
		} else {
			x &= ^(1 << (len(c) - 1 - i))
		}
	}
	fmt.Println(int(x) * int(^x&4095))
}
