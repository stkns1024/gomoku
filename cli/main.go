package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stkns1024/gomoku"
)

const size = (gomoku.Length+1)*(gomoku.Length+1)*2 - 1

type board struct {
	*gomoku.Board
	str []byte
}

func newBoard() *board {
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

	return &board{gomoku.NewBoard(), str}
}

func (b *board) place(stone byte, x, y uint8) error {
	err := b.Board.Place(stone, x, y)
	if err != nil {
		return err
	}

	pos := (y+1)*(gomoku.Length*2+2) + x*2 + 2
	b.str[pos] = stone

	return nil
}

func (b *board) String() string {
	return string(b.str)
}

func main() {
	var (
		board      = newBoard()
		scanner    = bufio.NewScanner(os.Stdin)
		moveCursor = fmt.Sprintf("\033[G\033[%dA", gomoku.Length+2)
	)

	for i := 0; i < gomoku.Size; i++ {
		fmt.Println(board)

		var stone byte
		if i%2 == 0 {
			stone = 'X'
		} else {
			stone = 'O'
		}

		var x, y uint8
		for {
			// 標準入力の読み込み
			fmt.Printf("\n%c>\033[K", stone)
			scanner.Scan()
			fmt.Print("\033[G\033[2A\033[2K")
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "入力エラー:", err)
				continue
			}
			pos := scanner.Text()
			if pos == "q" || pos == "quit" {
				fmt.Print("終了します\n\n")
				return
			} else if len(pos) != 2 {
				fmt.Fprint(os.Stderr, "位置は\"列行\"で指定してください。実際の入力:", pos)
				continue
			}
			x = pos[1] - 97
			y = pos[0] - 97

			err := board.place(stone, x, y)
			if err != nil {
				switch err.(type) {
				case gomoku.OutOfRangeError:
					fmt.Fprintf(os.Stderr, "%sは範囲外です", pos)
				case gomoku.AlreadyExistError:
					fmt.Fprintf(os.Stderr, "%sには既に石が存在します", pos)
				}
				continue
			}

			break
		}

		isChain, _ := board.IsChain(x, y)
		if isChain {
			fmt.Printf("%cの勝利\n\n", stone)
			return
		}

		fmt.Println(moveCursor)
	}
}
