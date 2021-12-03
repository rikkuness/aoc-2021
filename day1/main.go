package main

import "os"

func main() {
	f, _ := os.Open("input.txt")
	part1(f)
	f.Seek(0, 0)
	part2(f)
}
