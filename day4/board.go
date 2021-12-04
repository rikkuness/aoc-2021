package main

type Cell struct {
	val  uint8
	tick bool
}

type Board struct {
	data [d][d]Cell
	done bool
	sum  int
}

func (b *Board) Mark(i uint8) {
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			if b.data[x][y].val == i {
				b.data[x][y].tick = true
				b.sum -= int(i)
				return
			}
		}
	}
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

func (b *Board) Bingo() bool {
	for i := 0; i < d; i++ {
		if b.rowDone(i) {
			b.done = true
			return true
		}
		if b.colDone(i) {
			b.done = true
			return true
		}
	}
	return false
}
