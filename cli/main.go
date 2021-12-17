package main

import (
	"fmt"

	"github.com/stkns1024/gomoku"
)

const size = (gomoku.Length+1)*(gomoku.Length*2+2) - 1

type Board struct {
	*gomoku.Board
	str []byte
}

func NewBoard() *Board {
	//str := [size]byte{}
	str := make([]byte, size)

	// 列番号(1行目)
	str[0] = ' '
	idx := 1
	for i := 0; i < gomoku.Length; i++ {
		str[idx] = ' '
		str[idx+1] = byte(i + 97)
		idx += 2
	}
	str[idx] = '\n'
	idx++

	// 2行目以降
	for i := 0; ; i++ {
		str[idx] = byte(i + 97)
		idx++

		// 盤面
		for j := 1; j <= gomoku.Length; j++ {
			str[idx] = ' '
			str[idx+1] = '-'
			idx += 2
		}

		if idx >= size {
			break
		}

		str[idx] = '\n'
		idx++
	}

	board := gomoku.NewBoard()
	return &Board{board, str}
}

func (b *Board) String() string {
	return string(b.str)
}

func main() {
	board := NewBoard()
	board.Place('x', 0, 1)
	fmt.Println(board)
}
