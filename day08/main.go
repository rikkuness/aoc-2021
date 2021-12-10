package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type wiringMap map[rune]byte

var (
	base   wiringMap      = map[rune]byte{'a': 1, 'b': 2, 'c': 4, 'd': 8, 'e': 16, 'f': 32, 'g': 64}
	digits map[string]int = map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}
)

func charsToByte(d wiringMap, in string) (out byte) {
	for _, c := range in {
		out = out + d[c]
	}
	return
}

func deduceWiring(in []string) wiringMap {
	var (
		six, five []byte
		known     [10]byte
		wm        = wiringMap{}
	)

	for _, seq := range in {
		switch len(seq) {
		case 2: // identified number 1
			known[1] = charsToByte(base, seq)
		case 3: // identified number 7
			known[7] = charsToByte(base, seq)
		case 4: // identified number 4
			known[4] = charsToByte(base, seq)
		case 5: // identified numbers 2,3,5
			five = append(five, charsToByte(base, seq))
		case 6: // identified numbers 0,6,9
			six = append(six, charsToByte(base, seq))
		case 7: // identified number 8
			known[8] = charsToByte(base, seq)
		}
	}

	wm['a'] = known[7] &^ known[1] // a is 7 with 1 removed
	wm['x'] = known[4] &^ known[1] // b,d are 4 with 1 removed

	for _, s := range six {
		bb := known[4] + wm['a'] ^ s // 6 segment sequence with 4 and a removed
		if bits.OnesCount(uint(bb)) == 1 {
			wm['g'] = bb // only one segment lit means that must be g
		}
	}

	known[9] = known[4] + wm['a'] + wm['g'] // 4 with a and g is a 9
	wm['e'] = known[9] ^ known[8]           // e is 8 with 9 removed

	for _, s := range six {
		bb := known[7] + wm['e'] + wm['g'] ^ s // 6 segment sequence with 7, e and g removed
		if bits.OnesCount(uint(bb)) == 1 {
			wm['b'] = bb // only one segment lit, must be a b
		}
	}

	wm['d'] = wm['x'] &^ wm['b'] // d is bd with b removed

	for _, f := range five {
		bb := wm['a'] + wm['b'] + wm['d'] + wm['g'] ^ f
		if bits.OnesCount(uint(bb)) == 1 {
			wm['f'] = bb // a,b,d and g with only one must be f
		}
	}

	wm['c'] = known[1] &^ wm['f'] // c is 1 with f removed

	return wm
}

func main() {
	var (
		times [10]int
		sum   int
	)

	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	for s.Scan() {
		m := strings.Split(s.Text(), " ")
		wiring := deduceWiring(m[0:10])
		solved := make(map[byte]int, 10)

		for seq, i := range digits {
			solved[charsToByte(wiring, seq)] = i
		}

		code := strings.Builder{}
		for _, dig := range m[11:15] {
			a := charsToByte(base, dig)
			times[solved[a]]++
			code.WriteString(fmt.Sprintf("%d", solved[a]))
		}
		a, _ := strconv.Atoi(code.String())
		sum += a
	}

	fmt.Println(times[1] + times[4] + times[7] + times[8])
	fmt.Println(sum)
}
