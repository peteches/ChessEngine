package board

type BitBoard struct {
	Board uint64
}

func NewBitboard(initPositions ...Square) *BitBoard {
	bitBoard := BitBoard{
		Board: 0,
	}
	for _, x := range initPositions {
		bitBoard.FlipBit(x)
	}

	return &bitBoard
}

func (bb *BitBoard) FlipBit(bit Square) {
	bb.Board ^= uint64(bit)
}

func (bb *BitBoard) Squares() []Square {
	squares := []Square{}

	for _, square := range AllSquares {
		if (bb.Board & uint64(square)) > 0 {
			squares = append(squares, square)
		}
	}

	return squares
}

func (bb *BitBoard) Occupied(sqr Square) bool {
	return (bb.Board & uint64(sqr)) > 0
}

type Side uint8

const (
	White Side = iota
	Black
)

type Board struct {
	WhiteKing King
	BlackKing King

	WhiteQueens Queens
	BlackQueens Queens

	WhiteBishops Bishops
	BlackBishops Bishops

	WhiteKnights Knights
	BlackKnights Knights

	WhiteRooks Rooks
	BlackRooks Rooks

	WhitePawns Pawns
	BlackPawns Pawns
}
