package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(v int) int {
	return (v ^ v>>63) + (v >> 63 & 1)
}

func calc(in []int, p bool) int {
	mv := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		for _, v := range in {
			if p {
				mv[i] += abs(i-v) * (abs(i-v) + 1) / 2
			} else {
				mv[i] += abs(i - v)
			}
		}
	}
	sort.Ints(mv)
	return mv[0]
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)
	var all []int
	for _, s := range strings.Split(string(b), ",") {
		v, _ := strconv.Atoi(s)
		all = append(all, v)
	}

	fmt.Println(calc(all, false))
	fmt.Println(calc(all, true))
}
