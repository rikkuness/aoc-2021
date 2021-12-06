package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func sim(w [9]int, days int) (c int) {
	for d := 1; d < days; d++ {
		w[(d+7)%9] += w[d%9]
	}
	for _, d := range w {
		c += d
	}
	return
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	var w [9]int
	for _, i := range strings.Split(string(f), ",") {
		a, _ := strconv.Atoi(i)
		w[int8(a)] += 1
	}
	fmt.Println(sim(w, 80))
	fmt.Println(sim(w, 256))
}
