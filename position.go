package main

import (
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

//nolint:varnamelen // these are board coordinates. longer names do not make sense
const (
	A8 uint64 = iota + 1
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

// type MAILBOX [64]uint32

// const (
// 	W_KING uint32 = iota
// 	W_QUEEN
// 	W_BISHOP_1
// 	W_BISHOP_2
// 	W_KNIGHT_1
// 	W_KNIGHT_2
// 	W_ROOK_1
// 	W_ROOK_2
// 	W_PAWN_1
// 	W_PAWN_2
// 	W_PAWN_3
// 	W_PAWN_4
// 	W_PAWN_5
// 	W_PAWN_6
// 	W_PAWN_7
// 	W_PAWN_8
// 	B_KING
// 	B_QUEEN
// 	B_BISHOP_1
// 	B_BISHOP_2
// 	B_KNIGHT_1
// 	B_KNIGHT_2
// 	B_ROOK_1
// 	B_ROOK_2
// 	B_PAWN_1
// 	B_PAWN_2
// 	B_PAWN_3
// 	B_PAWN_4
// 	B_PAWN_5
// 	B_PAWN_6
// 	B_PAWN_7
// 	B_PAWN_8
// )

const (
	WHITE uint8 = iota
	BLACK
)

const (
	//	NoCastlingAllowed     = 0
	WhiteKingSideAllowed  = 1
	WhiteQueenSideAllowed = 2
	BlackKingSideAllowed  = 4
	BlackQueenSideAllowed = 8
)

type BITBOARD struct {
	board uint64
}

func NewBitboard(initPositions ...uint64) *BITBOARD {
	bitBoard := BITBOARD{
		board: 0,
	}
	for _, x := range initPositions {
		bitBoard.FlipBit(x)
	}

	return &bitBoard
}

func (bb *BITBOARD) FlipBit(bit uint64) {
	log.Debug().Uint64("bit", bit).Msg("Flipping Bit")

	switch bit {
	case 0:
		{
		}
	case 1:
		{
			oboard := bb.board
			nboard := bb.board ^ bit
			bb.board = nboard
			log.Debug().Interface("OldBoard", oboard).
				Interface("NewBoard", nboard).
				Interface("UpdatedBoard", bb.board).
				Msg("BitFlipped")
		}
	default:
		{
			bb.board ^= (1 << bit)
		}
	}
}

type PiecePositions struct {
	WhiteKing   *BITBOARD
	WhiteQueen  *BITBOARD
	WhiteBishop *BITBOARD
	WhiteKnight *BITBOARD
	WhiteRook   *BITBOARD
	WhitePawn   *BITBOARD
	BlackKing   *BITBOARD
	BlackQueen  *BITBOARD
	BlackBishop *BITBOARD
	BlackKnight *BITBOARD
	BlackRook   *BITBOARD
	BlackPawn   *BITBOARD
}

func NewPiecePositions() *PiecePositions {
	return &PiecePositions{
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
		NewBitboard(),
	}
}

type Position struct {
	Pieces          *PiecePositions
	EnPassantTarget uint64
	SideToMove      uint8
	CastlingRights  uint8
	HalfmoveClock   uint8
	FullMoveCounter uint8
}

// nolint:funlen // this case statement pushes over the limit but is largely
// unavoidable.
// complexity in this function is as simple as can be made.
func (p *Position) setPieces(pieces string) *PiecePositionError { // nolint:cyclop
	// Setting Pieces to required Positions
	newPosition := NewPiecePositions()
	offset := 1

	for index, pos := range strings.ReplaceAll(pieces, "/", "") {
		switch pos {
		case 'r':
			{
				newPosition.BlackRook.FlipBit(uint64(index + offset))
			}
		case 'n':
			{
				newPosition.BlackKnight.FlipBit(uint64(index + offset))
			}
		case 'b':
			{
				newPosition.BlackBishop.FlipBit(uint64(index + offset))
			}
		case 'q':
			{
				newPosition.BlackQueen.FlipBit(uint64(index + offset))
			}
		case 'k':
			{
				newPosition.BlackKing.FlipBit(uint64(index + offset))
			}
		case 'p':
			{
				newPosition.BlackPawn.FlipBit(uint64(index + offset))
			}
		case 'R':
			{
				newPosition.WhiteRook.FlipBit(uint64(index + offset))
			}
		case 'N':
			{
				newPosition.WhiteKnight.FlipBit(uint64(index + offset))
			}
		case 'B':
			{
				newPosition.WhiteBishop.FlipBit(uint64(index + offset))
			}
		case 'Q':
			{
				newPosition.WhiteQueen.FlipBit(uint64(index + offset))
			}
		case 'K':
			{
				newPosition.WhiteKing.FlipBit(uint64(index + offset))
			}
		case 'P':
			{
				newPosition.WhitePawn.FlipBit(uint64(index + offset))
			}
		case '1', '2', '3', '4', '5', '6', '7', '8':
			{
				offset += (int(pos-'0') - 1)
			}
		default:
			{
				return &PiecePositionError{
					errPiece: pos,
				}
			}
		}
	}

	p.Pieces = newPosition

	return nil
}

func (p *Position) setSideToMove(side string) *SideToMoveError {
	switch side {
	case "w":
		{
			p.SideToMove = WHITE
		}
	case "b":
		{
			p.SideToMove = BLACK
		}
	default:
		{
			return &SideToMoveError{errSide: side}
		}
	}

	return nil
}

func (p *Position) setCastlingRights(rights string) *CastlingRightsError {
	var blackKingSide uint8

	var blackQueenSide uint8

	var whiteKingSide uint8

	var whiteQueenSide uint8

	for _, pos := range rights {
		switch pos {
		case 'k':
			{
				blackKingSide = BlackKingSideAllowed
			}
		case 'q':
			{
				blackQueenSide = BlackQueenSideAllowed
			}
		case 'K':
			{
				whiteKingSide = WhiteKingSideAllowed
			}
		case 'Q':
			{
				whiteQueenSide = WhiteQueenSideAllowed
			}
		case '-':
			{
				p.CastlingRights = 0

				return nil
			}
		default:
			{
				return &CastlingRightsError{errChar: pos}
			}
		}
	}

	p.CastlingRights = 0 ^ (blackKingSide | blackQueenSide | whiteKingSide | whiteQueenSide)

	return nil
}

func (p *Position) setEnPassantTarget(targetSquare string) *EnPassantTargetError {
	enPassantMatrix := map[string]uint64{
		"-":  0,
		"A3": A3,
		"B3": B3,
		"C3": C3,
		"D3": D3,
		"E3": E3,
		"F3": F3,
		"G3": G3,
		"H3": H3,
		"A6": A6,
		"B6": B6,
		"C6": C6,
		"D6": D6,
		"E6": E6,
		"F6": F6,
		"G6": G6,
		"H6": H6,
	}

	var ok bool
	p.EnPassantTarget, ok = enPassantMatrix[targetSquare]

	if ok {
		return nil
	}

	return &EnPassantTargetError{errTarget: targetSquare}
}

func (p *Position) SetPositionFromFen(fen string) error {
	fenElements := strings.Split(fen, " ")

	pieceErr := p.setPieces(fenElements[0])
	if pieceErr != nil {
		pieceErr.fen = fen

		return pieceErr
	}

	stmErr := p.setSideToMove(fenElements[1])
	if stmErr != nil {
		stmErr.fen = fen

		return stmErr
	}

	castlingErr := p.setCastlingRights(fenElements[2])
	if castlingErr != nil {
		castlingErr.fen = fen

		return castlingErr
	}

	// Setting EnPassantTarget square
	enPassantErr := p.setEnPassantTarget(fenElements[3])
	if enPassantErr != nil {
		enPassantErr.fen = fen

		return enPassantErr
	}

	// nolint:gomnd // parse uint wants base and bits which arn't easily
	// derived
	halfMoveClock, err := strconv.ParseUint(fenElements[4], 10, 8)
	if err != nil {
		return &HalfMoveClockError{
			fen:           fen,
			err:           err.Error(),
			halfMoveClock: fenElements[4],
		}
	}

	p.HalfmoveClock = uint8(halfMoveClock)

	// nolint:gomnd // parse uint wants base and bits which arn't easily
	// derived
	halfMoveClock, err = strconv.ParseUint(fenElements[5], 10, 8)
	if err != nil {
		return &FullMoveCounterError{
			fen:             fen,
			err:             err.Error(),
			fullMoveCounter: fenElements[5],
		}
	}

	p.FullMoveCounter = uint8(halfMoveClock)

	return nil
}

func NewPosition() *Position {
	pos := Position{
		Pieces:          NewPiecePositions(),
		SideToMove:      WHITE,
		CastlingRights:  0,
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	}

	return &pos
}
