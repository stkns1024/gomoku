package gomoku

import "fmt"

type InvalidStoneError byte

func (e InvalidStoneError) Error() string {
	return fmt.Sprintf("stone should 'o' or 'x', but is '%v'", string(e))
}

type InvalidPositionError struct {
	x uint8
	y uint8
}

func (e InvalidPositionError) Error() string {
	return fmt.Sprintf("invalid area (%v, %v)", e.x, e.y)
}
