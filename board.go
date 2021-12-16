package gomoku

const Length = 15
const Size = 225
const ChainSize = 5

type Board [Size]byte

func NewBoard() *Board {
	var board Board
	for i := 0; i < Size; i++ {
		board[i] = ' '
	}

	return &board
}

func (b *Board) Place(stone byte, x, y uint8) error {
	if stone != 'o' && stone != 'x' {
		return InvalidStoneError(stone)
	}

	if x >= Length || y >= Length {
		return InvalidPositionError{x, y}
	}

	position := x + y*Length
	if b[position] != ' ' {
		return InvalidPositionError{x, y}
	}

	b[position] = stone

	return nil
}

func (b *Board) IsChain(x, y uint8) (bool, error) {
	if x >= Length || y >= Length {
		err := InvalidPositionError{x, y}
		return false, err
	}

	pos := x + y*Length
	stone := b[pos]
	shifts := [4]uint8{1, 7, 8, 9}
	for _, shift := range shifts {
		length := 1
		newPos := pos
		for ; length < ChainSize; length++ {
			isEdge := newPos/Length == 0 || newPos%Length == 0
			if isEdge {
				break
			}

			newPos -= shift

			if b[newPos] != stone {
				break
			}
		}

		for ; length < ChainSize; length++ {
			isEdge := newPos/Length == Length || newPos%Length == Length
			if isEdge {
				break
			}

			newPos += shift

			if b[newPos] != stone {
				break
			}
		}

		if length >= ChainSize {
			return true, nil
		}
	}

	return false, nil
}
