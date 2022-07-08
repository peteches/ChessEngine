package main

import (
	"fmt"
	"strconv"
	"strings"
)

const totalSquares = 64

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

// nolint:gochecknoglobals // this is a pseudo const
var boardMatrix = map[uint64]string{
	1:  "A8",
	2:  "B8",
	3:  "C8",
	4:  "D8",
	5:  "E8",
	6:  "F8",
	7:  "G8",
	8:  "H8",
	9:  "A7",
	10: "B7",
	11: "C7",
	12: "D7",
	13: "E7",
	14: "F7",
	15: "G7",
	16: "H7",
	17: "A6",
	18: "B6",
	19: "C6",
	20: "D6",
	21: "E6",
	22: "F6",
	23: "G6",
	24: "H6",
	25: "A5",
	26: "B5",
	27: "C5",
	28: "D5",
	29: "E5",
	30: "F5",
	31: "G5",
	32: "H5",
	33: "A4",
	34: "B4",
	35: "C4",
	36: "D4",
	37: "E4",
	38: "F4",
	39: "G4",
	40: "H4",
	41: "A3",
	42: "B3",
	43: "C3",
	44: "D3",
	45: "E3",
	46: "F3",
	47: "G3",
	48: "H3",
	49: "A2",
	50: "B2",
	51: "C2",
	52: "D2",
	53: "E2",
	54: "F2",
	55: "G2",
	56: "H2",
	57: "A1",
	58: "B1",
	59: "C1",
	60: "D1",
	61: "E1",
	62: "F1",
	63: "G1",
	64: "H1",
}

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
	bb.board ^= (1 << (bit - 1))
}

func (bb *BITBOARD) Squares() []string {
	squares := []string{}

	if (bb.board & 1) > 0 {
		squares = append(squares, boardMatrix[1])
	}

	// nolint:gomnd // have to skip 1 which is deal with above due to bit
	// shift.
	for square := uint64(2); square <= totalSquares; square++ {
		if (bb.board & (1 << square)) > 0 {
			squares = append(squares, boardMatrix[square])
		}
	}

	return squares
}

func (bb *BITBOARD) Occupied(sqr uint8) bool {
	return (bb.board & (1 << (sqr - 1))) > 0
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

// nolint:cyclop // not sure how to simplify this further
func (p *PiecePositions) Occupied(sqr uint8) bool {
	return p.WhiteKing.Occupied(sqr) ||
		p.WhiteQueen.Occupied(sqr) ||
		p.WhiteBishop.Occupied(sqr) ||
		p.WhiteKnight.Occupied(sqr) ||
		p.WhiteRook.Occupied(sqr) ||
		p.WhitePawn.Occupied(sqr) ||
		p.BlackKing.Occupied(sqr) ||
		p.BlackQueen.Occupied(sqr) ||
		p.BlackBishop.Occupied(sqr) ||
		p.BlackKnight.Occupied(sqr) ||
		p.BlackRook.Occupied(sqr) ||
		p.BlackPawn.Occupied(sqr)
}

// nolint:cyclop // not sure how to simplify this further
func (p *PiecePositions) OccupiedBy(sqr uint8) string {
	if !p.Occupied(sqr) {
		return ""
	}

	if p.WhiteKing.Occupied(sqr) {
		return "K"
	}

	if p.WhiteQueen.Occupied(sqr) {
		return "Q"
	}

	if p.WhiteBishop.Occupied(sqr) {
		return "B"
	}

	if p.WhiteKnight.Occupied(sqr) {
		return "N"
	}

	if p.WhiteRook.Occupied(sqr) {
		return "R"
	}

	if p.WhitePawn.Occupied(sqr) {
		return "P"
	}

	if p.BlackKing.Occupied(sqr) {
		return "k"
	}

	if p.BlackQueen.Occupied(sqr) {
		return "q"
	}

	if p.BlackBishop.Occupied(sqr) {
		return "b"
	}

	if p.BlackKnight.Occupied(sqr) {
		return "n"
	}

	if p.BlackRook.Occupied(sqr) {
		return "r"
	}

	if p.BlackPawn.Occupied(sqr) {
		return "p"
	}

	return "Bla"
}

func (p *PiecePositions) String() string {
	fen := ""
	unoccupied := 0

	for index := uint8(1); index <= totalSquares; index++ {
		if p.Occupied(index) {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			fen += p.OccupiedBy(index)
		} else {
			unoccupied++
		}

		if index%8 == 0 {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			if index < totalSquares {
				fen += "/"
			}
		}
	}

	return fen
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

func (p *Position) String() string {
	fen := ""

	fen += p.Pieces.String()
	fen += " "

	switch p.SideToMove {
	case WHITE:
		fen += "w"
	case BLACK:
		fen += "b"
	}

	fen += " "

	if p.CastlingRights > 0 {
		for rights := p.CastlingRights; rights != 0; {
			switch {
			case rights&WhiteKingSideAllowed > 0:
				rights ^= WhiteKingSideAllowed
				fen += "K"
			case rights&WhiteQueenSideAllowed > 0:
				rights ^= WhiteQueenSideAllowed
				fen += "Q"
			case rights&BlackKingSideAllowed > 0:
				rights ^= BlackKingSideAllowed
				fen += "k"
			case rights&BlackQueenSideAllowed > 0:
				rights ^= BlackQueenSideAllowed
				fen += "q"
			}
		}
	} else {
		fen += "-"
	}

	fen += " "

	if p.EnPassantTarget == 0 {
		fen += "-"
	} else {
		fen += boardMatrix[p.EnPassantTarget]
	}

	fen += " "

	fen += fmt.Sprintf("%d %d", p.HalfmoveClock, p.FullMoveCounter)

	return fen
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
