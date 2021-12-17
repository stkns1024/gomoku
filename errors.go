package gomoku

import "fmt"

type InvalidStoneError byte

func (e InvalidStoneError) Error() string {
	return fmt.Sprintf("stone should 'O' or 'X', but is '%v'", string(e))
}

type InvalidPositionError struct {
	x uint8
	y uint8
}

func (e InvalidPositionError) Error() string {
	return fmt.Sprintf("invalid position (%v, %v)", e.x, e.y)
}
