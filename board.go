package gomoku

const (
	Length      = 15
	Size        = 225
	ChainLength = 5
)

func min(vals ...uint8) uint8 {
	minVal := vals[0]
	for _, val := range vals[1:] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

type Board [Size]byte

func NewBoard() *Board {
	var board Board
	for i := 0; i < Size; i++ {
		board[i] = '-'
	}
	return &board
}

func (b *Board) Place(stone byte, x, y uint8) error {
	if stone != 'O' && stone != 'X' {
		return InvalidStoneError(stone)
	}

	if x >= Length || y >= Length {
		return InvalidPositionError{x, y}
	}

	pos := x + y*Length

	if b[pos] != '-' {
		return InvalidPositionError{x, y}
	}

	b[pos] = stone
	return nil
}

func (b *Board) shift(stone byte, pos uint8, step uint8, maxIter int) uint8 {
	var length uint8 = 0
	for i := 0; i < maxIter; i++ {
		pos += step
		if b[pos] != stone {
			break
		}
		length++
	}
	return length
}

func (b *Board) IsChain(x, y uint8) (bool, error) {
	if x >= Length || y >= Length {
		err := InvalidPositionError{x, y}
		return false, err
	}

	var (
		pos   = x + y*Length
		stone = b[pos]
	)

	if stone != 'O' && stone != 'X' {
		return false, InvalidStoneError(stone)
	}

	steps := [4]uint8{1, Length - 1, Length, Length + 1}
	for _, step := range steps {
		var (
			length         uint8 = 1
			maxShifts      uint8 = Length - 1
			shiftV, shiftH uint8
		)

		if step != 1 {
			shiftV = maxShifts - y
		} else {
			shiftV = Length
		}

		switch step % Length {
		case 0:
			shiftH = Length
		case 1:
			shiftH = maxShifts - x
		default:
			shiftH = x
		}

		maxIter := int(min(ChainLength-length, shiftV, shiftH))
		length += b.shift(stone, pos, step, maxIter)

		if shiftV != Length {
			shiftV = maxShifts - shiftV
		}

		if shiftH != Length {
			shiftH = maxShifts - shiftH
		}

		maxIter = int(min(ChainLength-length, shiftV, shiftH))
		length += b.shift(stone, pos, -step, maxIter)

		if length >= ChainLength {
			return true, nil
		}
	}

	return false, nil
}
