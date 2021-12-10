package main

import "os"

const bs = 12

func main() {
	f, _ := os.Open("input.txt")
	part1(f)
	f.Seek(0, 0)
	part2(f)
}
