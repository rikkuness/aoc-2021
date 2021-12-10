package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

var (
	p1score     = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	p2score     = map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	pair        = regexp.MustCompile(`(\(\)|\[\]|\<\>|\{\})`)
	p1      int = 0
	p2      []int
)

func removePairs(line *string) {
	if pair.MatchString(*line) {
		(*line) = pair.ReplaceAllString(*line, "")
		removePairs(line)
	}
}

func corruption(line *string) (corrupt *rune) {
	for _, r := range *line {
		if r == '}' || r == ']' || r == '>' || r == ')' {
			corrupt = &r
			return
		}
	}
	return
}

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		removePairs(&line)
		if r := corruption(&line); r != nil {
			p1 += p1score[*r]
		} else {
			var l = 0
			for i := len(line) - 1; i >= 0; i-- {
				l = ((l * 5) + p2score[rune(line[i])])
			}
			p2 = append(p2, l)
		}
	}
	fmt.Println(p1)
	sort.Ints(p2)
	fmt.Println(p2[len(p2)/2])
}
