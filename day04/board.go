package main

type Cell struct {
	val  uint8
	tick bool
}

type Board struct {
	data [d][d]Cell
	win  bool
	sum  int
}

func (b *Board) Mark(i uint8) bool {
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			if b.data[x][y].val == i {
				b.data[x][y].tick = true
				b.sum -= int(i)
				if b.rowDone(x) || b.colDone(y) {
					b.win = true
				}
				return b.win
			}
		}
	}
	return b.win
}

func (b *Board) rowDone(rn int) bool {
	t := 0
	for i := 0; i < d; i++ {
		if b.data[rn][i].tick {
			t++
		}
	}
	return t == d
}

func (b *Board) colDone(cn int) bool {
	t := 0
	for i := 0; i < d; i++ {
		if b.data[i][cn].tick {
			t++
		}
	}
	return t == d
}
