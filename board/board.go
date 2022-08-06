package board

import (
	"strconv"
	"strings"

	"github.com/peteches/ChessEngine/errors"
)

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
	WhiteKing *King
	BlackKing *King

	WhiteQueens *Queens
	BlackQueens *Queens

	WhiteBishops *Bishops
	BlackBishops *Bishops

	WhiteKnights *Knights
	BlackKnights *Knights

	WhiteRooks *Rooks
	BlackRooks *Rooks

	WhitePawns *Pawns
	BlackPawns *Pawns
}

func NewBoard() *Board {
	return &Board{
		WhiteKing:    NewKing(White),
		BlackKing:    NewKing(Black),
		WhiteQueens:  NewQueens(White),
		BlackQueens:  NewQueens(Black),
		WhiteBishops: NewBishops(White),
		BlackBishops: NewBishops(Black),
		WhiteKnights: NewKnights(White),
		BlackKnights: NewKnights(Black),
		WhiteRooks:   NewRooks(White),
		BlackRooks:   NewRooks(Black),
		WhitePawns:   NewPawns(White),
		BlackPawns:   NewPawns(Black),
	}
}

// nolint: cyclop // can't be simplified any further
func (b *Board) OccupiedBy(sqr Square) string {
	switch {
	case b.WhiteKing.Positions.Occupied(sqr):
		return b.WhiteKing.String()
	case b.WhiteQueens.Positions.Occupied(sqr):
		return b.WhiteQueens.String()
	case b.WhiteBishops.Positions.Occupied(sqr):
		return b.WhiteBishops.String()
	case b.WhiteKnights.Positions.Occupied(sqr):
		return b.WhiteKnights.String()
	case b.WhiteRooks.Positions.Occupied(sqr):
		return b.WhiteRooks.String()
	case b.WhitePawns.Positions.Occupied(sqr):
		return b.WhitePawns.String()
	case b.BlackKing.Positions.Occupied(sqr):
		return b.BlackKing.String()
	case b.BlackQueens.Positions.Occupied(sqr):
		return b.BlackQueens.String()
	case b.BlackBishops.Positions.Occupied(sqr):
		return b.BlackBishops.String()
	case b.BlackKnights.Positions.Occupied(sqr):
		return b.BlackKnights.String()
	case b.BlackRooks.Positions.Occupied(sqr):
		return b.BlackRooks.String()
	case b.BlackPawns.Positions.Occupied(sqr):
		return b.BlackPawns.String()
	default:
		return ""
	}
}

func (b *Board) Occupied(sqr Square) bool {
	occupiedBy := b.OccupiedBy(sqr)

	return len(occupiedBy) > 0
}

func (b *Board) String() string {
	fen := ""
	unoccupied := 0

	for idx, sqr := range AllSquares {
		if b.Occupied(sqr) {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			fen += b.OccupiedBy(sqr)
		} else {
			unoccupied++
		}

		if (idx+1)%8 == 0 {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			if (idx + 1) < TotalSquares {
				fen += "/"
			}
		}
	}

	return fen
}

// nolint:funlen,cyclop // TODO look are refactoring this
func (b *Board) SetPieces(pieces string) *errors.PiecePositionError {
	// reset all positions
	b.BlackRooks.Positions.Board &= uint64(0)
	b.BlackKnights.Positions.Board &= uint64(0)
	b.BlackBishops.Positions.Board &= uint64(0)
	b.BlackQueens.Positions.Board &= uint64(0)
	b.BlackKing.Positions.Board &= uint64(0)
	b.BlackPawns.Positions.Board &= uint64(0)
	b.WhiteRooks.Positions.Board &= uint64(0)
	b.WhiteKnights.Positions.Board &= uint64(0)
	b.WhiteBishops.Positions.Board &= uint64(0)
	b.WhiteQueens.Positions.Board &= uint64(0)
	b.WhiteKing.Positions.Board &= uint64(0)
	b.WhitePawns.Positions.Board &= uint64(0)

	offset := 0

	// fen strings are odd. They move from A8-H8,A7-H7 etc.
	// I want them to go A1-H1,A2-H2 etc. This then makes more sense when
	// using the Constants A1-H8 where A1 is the smallest and H8 is the
	// largest. Using fenstrings as is A8 is smaller than A1, which isn't
	// very intuitive. This does mean I need to manipulate the fen string
	// slightly to allow more intuitive BitBoards
	pieceRanks := strings.Split(pieces, "/")
	piecesSensibleOrder := ""

	for i := 7; i >= 0; i-- {
		piecesSensibleOrder += pieceRanks[i]
	}

	// this case statement pushes over the limit for funlen
	// but is largely unavoidable. complexity in this function is as simple as can be made.
	for index, pos := range piecesSensibleOrder {
		switch pos {
		case 'r':
			{
				b.BlackRooks.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'n':
			{
				b.BlackKnights.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'b':
			{
				b.BlackBishops.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'q':
			{
				b.BlackQueens.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'k':
			{
				b.BlackKing.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'p':
			{
				b.BlackPawns.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'R':
			{
				b.WhiteRooks.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'N':
			{
				b.WhiteKnights.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'B':
			{
				b.WhiteBishops.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'Q':
			{
				b.WhiteQueens.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'K':
			{
				b.WhiteKing.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case 'P':
			{
				b.WhitePawns.Positions.FlipBit(Square(1 << (index + offset)))
			}
		case '1', '2', '3', '4', '5', '6', '7', '8':
			{
				offset += (int(pos-'0') - 1)
			}
		default:
			{
				return &errors.PiecePositionError{
					ErrPiece: pos,
				}
			}
		}
	}

	return nil
}
