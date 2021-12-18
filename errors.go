package gomoku

import "fmt"

type InvalidStoneError byte

func (e InvalidStoneError) Error() string {
	return fmt.Sprint("使用できる石は'X'または'O'です。実際の値:", string(e))
}

type OutOfRangeError struct {
	x uint8
	y uint8
}

func (e OutOfRangeError) Error() string {
	return fmt.Sprintf("(%d, %d)は範囲外です", e.x, e.y)
}

type AlreadyExistError struct {
	x uint8
	y uint8
}

func (e AlreadyExistError) Error() string {
	return fmt.Sprintf("(%d, %d)には既に石が存在します", e.x, e.y)
}
