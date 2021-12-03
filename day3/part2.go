package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func filter(in []uint16, p int, mc bool) (out []uint16) {
	t := 0
	for _, v := range in {
		t += int(v >> (bs - (p + 1)) & 1)
	}

	var doo uint16 = 0
	if float32(t) >= float32(len(in))/2 {
		doo = 1
	}

	if !mc {
		doo = 1 ^ doo
	}

	for _, v := range in {
		if (v >> (bs - (p + 1)) & 1) == doo {
			out = append(out, v)
		}
	}

	return
}

func part2(f io.Reader) {
	s := bufio.NewScanner(f)
	var all []uint16
	for s.Scan() {
		r := s.Text()
		v, _ := strconv.ParseUint(r, 2, bs)
		all = append(all, uint16(v))
	}

	var oxy, co2 = all, all
	for i := 0; i < bs; i++ {
		oxy = filter(oxy, i, true)
		if len(oxy) > 1 {
			continue
		}
		break
	}

	for i := 0; i < bs; i++ {
		co2 = filter(co2, i, false)
		if len(co2) > 1 {
			continue
		}
		break
	}

	fmt.Println(int(oxy[0]) * int(co2[0]))
}
