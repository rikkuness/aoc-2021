package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const d = 5

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	var callouts []uint8
	boards := make([]*Board, 0)

	for i := -1; s.Scan(); i++ {
		if i < 0 {
			for _, v := range strings.Split(s.Text(), ",") {
				ui, _ := strconv.ParseUint(v, 10, 8)
				callouts = append(callouts, uint8(ui))
			}
			continue
		}

		ui, _ := strconv.ParseUint(s.Text(), 10, 8)
		if len(boards) <= i/(d*d) {
			boards = append(boards, &Board{})
		}

		boards[i/(d*d)].data[(i%(d*d)/d)%d][i%(d*d)%d].val = uint8(ui)
		boards[i/(d*d)].sum += int(ui)
	}

	for _, c := range callouts {
		for bi, b := range boards {
			if !b.win {
				if b.Mark(c) {
					fmt.Printf("board %2d complete\tsum: %d\n", bi, b.sum*int(c))
				}
			}
		}
	}
}
