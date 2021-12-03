package main

import (
	"bufio"
	"fmt"
	"io"
)

func part1(f io.Reader) {
	var (
		c [bs]int
		x uint16
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		r := s.Text()
		for i := 0; i < bs; i++ {
			if !(r[i]&1<<5 == 0) {
				c[i] += 1
			} else {
				c[i] -= 1
			}
			if c[i] > 0 {
				x |= 1 << (bs - 1 - i)
			} else {
				x &= ^(1 << (bs - 1 - i))
			}
		}
	}

	z := (^uint16(0)) >> (16 - bs)
	fmt.Println(int(x) * int(^x&z))
}
