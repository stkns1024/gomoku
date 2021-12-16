package gomoku

const SIZE = 225

type Board [SIZE]byte

func NewBoard() *Board {
	var board Board
	for i := 0; i < SIZE; i++ {
		board[i] = ' '
	}

	return &board
}

func (b *Board) Place(stone byte, x, y uint8) error {
	if stone != 'o' && stone != 'x' {
		return InvalidStoneError(stone)
	}

	if x >= 15 || y >= 15 {
		return InvalidPositionError{x, y}
	}

	position := x + y*15
	if b[position] != ' ' {
		return InvalidPositionError{x, y}
	}

	b[position] = stone

	return nil
}

func (b *Board) IsChain(x, y uint8) (bool, error) {
	if x >= 15 || y >= 15 {
		err := InvalidPositionError{x, y}
		return false, err
	}

	return false, nil
}
