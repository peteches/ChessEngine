package main

import (
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

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
	bb := BITBOARD{
		board: 0,
	}
	for _, x := range initPositions {
		bb.FlipBit(x)
	}
	return &bb
}

func (bb *BITBOARD) FlipBit(p uint64) {
	log.Debug().Uint64("bit", p).Msg("Flipping Bit")
	switch p {
	case 0:
		{
		}
	case 1:
		{
			oboard := bb.board
			nboard := bb.board ^ p
			bb.board = nboard
			log.Debug().Interface("OldBoard", oboard).
				Interface("NewBoard", nboard).
				Interface("UpdatedBoard", bb.board).
				Msg("BitFlipped")
		}
	default:
		{
			bb.board ^= (1 << p)
		}
	}
}

type PiecePositions struct {
	W_KING   *BITBOARD
	W_QUEEN  *BITBOARD
	W_BISHOP *BITBOARD
	W_KNIGHT *BITBOARD
	W_ROOK   *BITBOARD
	W_PAWN   *BITBOARD
	B_KING   *BITBOARD
	B_QUEEN  *BITBOARD
	B_BISHOP *BITBOARD
	B_KNIGHT *BITBOARD
	B_ROOK   *BITBOARD
	B_PAWN   *BITBOARD
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
	SideToMove      uint8
	CastlingAbility uint8
	EnPassantTarget uint64
	HalfmoveClock   uint8
	FullMoveCounter uint8
}

func (p *Position) SetPositionFromFen(fen string) error {
	llogger := log.With().Str("fen", fen).Logger()
	llogger.Info().Msg("Setting position to new fen")

	fenElements := strings.Split(fen, " ")

	// Setting Pieces to required Positions
	newPosition := NewPiecePositions()
	offset := 1
	for index, pos := range strings.ReplaceAll(fenElements[0], "/", "") {
		pllogger := llogger.With().
			Str("fenElement", strconv.QuoteRune(pos)).
			Int("index", index).
			Int("offset", offset).
			Logger()
		pllogger.Debug().Msg("")
		switch pos {
		case 'r':
			{
				newPosition.B_ROOK.FlipBit(uint64(index + offset))
			}
		case 'n':
			{
				newPosition.B_KNIGHT.FlipBit(uint64(index + offset))
			}
		case 'b':
			{
				newPosition.B_BISHOP.FlipBit(uint64(index + offset))
			}
		case 'q':
			{
				newPosition.B_QUEEN.FlipBit(uint64(index + offset))
			}
		case 'k':
			{
				newPosition.B_KING.FlipBit(uint64(index + offset))
			}
		case 'p':
			{
				newPosition.B_PAWN.FlipBit(uint64(index + offset))
			}
		case 'R':
			{
				newPosition.W_ROOK.FlipBit(uint64(index + offset))
			}
		case 'N':
			{
				newPosition.W_KNIGHT.FlipBit(uint64(index + offset))
			}
		case 'B':
			{
				newPosition.W_BISHOP.FlipBit(uint64(index + offset))
			}
		case 'Q':
			{
				newPosition.W_QUEEN.FlipBit(uint64(index + offset))
			}
		case 'K':
			{
				newPosition.W_KING.FlipBit(uint64(index + offset))
			}
		case 'P':
			{
				newPosition.W_PAWN.FlipBit(uint64(index + offset))
			}
		case '1', '2', '3', '4', '5', '6', '7', '8':
			{
				pllogger.Debug().
					Msg("Updating offset by fenElement -1")
				offset += (int(pos-'0') - 1)
				pllogger.Debug().
					Msg("Updated index by fenElement -1")
			}
		default:
			{
				return &piecePositionError{
					fen:      fen,
					errPiece: pos,
				}
			}
		}
	}
	llogger.Debug().Interface("NewPosition", newPosition).Msg("Updating Position")
	p.Pieces = newPosition

	// Setting which side is to move
	llogger.Debug().Str("SideToMove", fenElements[1]).Msg("Setting Side to Move")
	switch fenElements[1] {
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
			return &sideToMoveError{fen: fen, errSide: fenElements[1]}
		}
	}

	llogger.Debug().Str("CastlingRights", fenElements[2]).Msg("Setting Castling Rights")
	// Setting Castling Rights
	for _, pos := range fenElements[2] {
		pllogger := llogger.With().
			Uint8("CurrentCastlingAbility", p.CastlingAbility).
			Str("fenElement", strconv.QuoteRune(pos)).
			Logger()
		pllogger.Debug().Msg("")
		switch pos {
		case 'k':
			{
				p.CastlingAbility ^= BlackKingSideAllowed
			}
		case 'q':
			{
				p.CastlingAbility ^= BlackQueenSideAllowed
			}
		case 'K':
			{
				p.CastlingAbility ^= WhiteKingSideAllowed
			}
		case 'Q':
			{
				p.CastlingAbility ^= WhiteQueenSideAllowed
			}
		}
	}

	// Setting EnPassantTarget square
	switch fenElements[3] {
	case "A3":
		p.EnPassantTarget = A3
	case "B3":
		p.EnPassantTarget = B3
	case "C3":
		p.EnPassantTarget = C3
	case "D3":
		p.EnPassantTarget = D3
	case "E3":
		p.EnPassantTarget = E3
	case "F3":
		p.EnPassantTarget = F3
	case "G3":
		p.EnPassantTarget = G3
	case "H3":
		p.EnPassantTarget = H3
	case "A6":
		p.EnPassantTarget = A6
	case "B6":
		p.EnPassantTarget = B6
	case "C6":
		p.EnPassantTarget = C6
	case "D6":
		p.EnPassantTarget = D6
	case "E6":
		p.EnPassantTarget = E6
	case "F6":
		p.EnPassantTarget = F6
	case "G6":
		p.EnPassantTarget = G6
	case "H6":
		p.EnPassantTarget = H6
	case "-":
		p.EnPassantTarget = 0
	default:
		return &enPassantTargetError{fen: fen, errTarget: fenElements[3]}
	}

	// Set HalfmoveClock
	i, err := strconv.ParseUint(fenElements[4], 10, 8)
	if err != nil {
		return &halfMoveClockError{
			fen:           fen,
			err:           err.Error(),
			halfMoveClock: fenElements[4],
		}
	}
	p.HalfmoveClock = uint8(i)
	// Set FullMoveCounter
	i, err = strconv.ParseUint(fenElements[5], 10, 8)
	if err != nil {
		return &fullMoveCounterError{
			fen:             fen,
			err:             err.Error(),
			fullMoveCounter: fenElements[5],
		}
	}
	p.FullMoveCounter = uint8(i)
	return nil
}

func NewPosition() *Position {
	pos := Position{
		Pieces:          NewPiecePositions(),
		SideToMove:      WHITE,
		CastlingAbility: 0,
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	}
	return &pos
}
