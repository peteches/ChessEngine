package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const totalSquares = 64

const numFenElements = 6

const (
	firstRank uint8 = iota + 1
	secondRank
	thirdRank
	fourthRank
	fifthRank
	sixthRank
	seventhRank
	eighthRank
)

const (
	AFile uint8 = iota + 1
	BFile
	CFile
	DFile
	EFile
	FFile
	GFile
	HFile
)

type Square uint64

func (s *Square) File() uint8 {
	switch *s {
	case A1, A2, A3, A4, A5, A6, A7, A8:
		return AFile
	case B1, B2, B3, B4, B5, B6, B7, B8:
		return BFile
	case C1, C2, C3, C4, C5, C6, C7, C8:
		return CFile
	case D1, D2, D3, D4, D5, D6, D7, D8:
		return DFile
	case E1, E2, E3, E4, E5, E6, E7, E8:
		return EFile
	case F1, F2, F3, F4, F5, F6, F7, F8:
		return FFile
	case G1, G2, G3, G4, G5, G6, G7, G8:
		return GFile
	case H1, H2, H3, H4, H5, H6, H7, H8:
		return HFile
	default:
		return 0
	}
}

func (s *Square) Rank() uint8 {
	switch *s {
	case A1, B1, C1, D1, E1, F1, H1, G1:
		return firstRank
	case A2, B2, C2, D2, E2, F2, H2, G2:
		return secondRank
	case A3, B3, C3, D3, E3, F3, H3, G3:
		return thirdRank
	case A4, B4, C4, D4, E4, F4, H4, G4:
		return fourthRank
	case A5, B5, C5, D5, E5, F5, H5, G5:
		return fifthRank
	case A6, B6, C6, D6, E6, F6, H6, G6:
		return sixthRank
	case A7, B7, C7, D7, E7, F7, H7, G7:
		return seventhRank
	case A8, B8, C8, D8, E8, F8, H8, G8:
		return eighthRank
	default:
		return 0
	}
}

//nolint:varnamelen // these are board coordinates. longer names do not make sense
const (
	A8 Square = 1
	B8 Square = 1 << iota
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
var allSquares = [64]Square{
	A8, B8, C8, D8, E8, F8, G8, H8,
	A7, B7, C7, D7, E7, F7, G7, H7,
	A6, B6, C6, D6, E6, F6, G6, H6,
	A5, B5, C5, D5, E5, F5, G5, H5,
	A4, B4, C4, D4, E4, F4, G4, H4,
	A3, B3, C3, D3, E3, F3, G3, H3,
	A2, B2, C2, D2, E2, F2, G2, H2,
	A1, B1, C1, D1, E1, F1, G1, H1,
}

// nolint:gochecknoglobals // this is a pseudo const
var boardMatrixStoI = map[string]Square{
	"A8": A8,
	"B8": B8,
	"C8": C8,
	"D8": D8,
	"E8": E8,
	"F8": F8,
	"G8": G8,
	"H8": H8,
	"A7": A7,
	"B7": B7,
	"C7": C7,
	"D7": D7,
	"E7": E7,
	"F7": F7,
	"G7": G7,
	"H7": H7,
	"A6": A6,
	"B6": B6,
	"C6": C6,
	"D6": D6,
	"E6": E6,
	"F6": F6,
	"G6": G6,
	"H6": H6,
	"A5": A5,
	"B5": B5,
	"C5": C5,
	"D5": D5,
	"E5": E5,
	"F5": F5,
	"G5": G5,
	"H5": H5,
	"A4": A4,
	"B4": B4,
	"C4": C4,
	"D4": D4,
	"E4": E4,
	"F4": F4,
	"G4": G4,
	"H4": H4,
	"A3": A3,
	"B3": B3,
	"C3": C3,
	"D3": D3,
	"E3": E3,
	"F3": F3,
	"G3": G3,
	"H3": H3,
	"A2": A2,
	"B2": B2,
	"C2": C2,
	"D2": D2,
	"E2": E2,
	"F2": F2,
	"G2": G2,
	"H2": H2,
	"A1": A1,
	"B1": B1,
	"C1": C1,
	"D1": D1,
	"E1": E1,
	"F1": F1,
	"G1": G1,
	"H1": H1,
}

// nolint:gochecknoglobals,dupl // this is a pseudo const
var boardMatrixItoS = map[Square]string{
	A8: "A8",
	B8: "B8",
	C8: "C8",
	D8: "D8",
	E8: "E8",
	F8: "F8",
	G8: "G8",
	H8: "H8",
	A7: "A7",
	B7: "B7",
	C7: "C7",
	D7: "D7",
	E7: "E7",
	F7: "F7",
	G7: "G7",
	H7: "H7",
	A6: "A6",
	B6: "B6",
	C6: "C6",
	D6: "D6",
	E6: "E6",
	F6: "F6",
	G6: "G6",
	H6: "H6",
	A5: "A5",
	B5: "B5",
	C5: "C5",
	D5: "D5",
	E5: "E5",
	F5: "F5",
	G5: "G5",
	H5: "H5",
	A4: "A4",
	B4: "B4",
	C4: "C4",
	D4: "D4",
	E4: "E4",
	F4: "F4",
	G4: "G4",
	H4: "H4",
	A3: "A3",
	B3: "B3",
	C3: "C3",
	D3: "D3",
	E3: "E3",
	F3: "F3",
	G3: "G3",
	H3: "H3",
	A2: "A2",
	B2: "B2",
	C2: "C2",
	D2: "D2",
	E2: "E2",
	F2: "F2",
	G2: "G2",
	H2: "H2",
	A1: "A1",
	B1: "B1",
	C1: "C1",
	D1: "D1",
	E1: "E1",
	F1: "F1",
	G1: "G1",
	H1: "H1",
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

func NewBitboard(initPositions ...Square) *BITBOARD {
	bitBoard := BITBOARD{
		board: 0,
	}
	for _, x := range initPositions {
		bitBoard.FlipBit(x)
	}

	return &bitBoard
}

func (bb *BITBOARD) FlipBit(bit Square) {
	bb.board ^= uint64(bit)
}

func (bb *BITBOARD) Squares() []string {
	squares := []string{}

	for _, square := range allSquares {
		if (bb.board & uint64(square)) > 0 {
			squares = append(squares, boardMatrixItoS[square])
		}
	}

	return squares
}

func (bb *BITBOARD) Occupied(sqr Square) bool {
	return (bb.board & uint64(sqr)) > 0
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
func (p *PiecePositions) Occupied(sqr Square) bool {
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
func (p *PiecePositions) OccupiedBy(sqr Square) string {
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

	for idx, sqr := range allSquares {
		if p.Occupied(sqr) {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			fen += p.OccupiedBy(sqr)
		} else {
			unoccupied++
		}

		if (idx+1)%8 == 0 {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			if (idx + 1) < totalSquares {
				fen += "/"
			}
		}
	}

	return fen
}

type Move struct {
	piece       string
	srcSquare   Square
	dstSquare   Square
	capture     bool
	promotionTo string
}

func NewMove(lanMove string) (*Move, *MoveError) {
	r := "(?i)^(?P<piece>[NBRQK])?(?P<src>[A-H][1-8])(?P<capture>[-X])?(?P<dst>[A-H][1-8])(?P<promotionTo>[NBRQ])?$"
	moveRegex := regexp.MustCompile(r)
	matches := moveRegex.FindStringSubmatch(lanMove)
	pieceIndex := moveRegex.SubexpIndex("piece")
	srcIndex := moveRegex.SubexpIndex("src")
	capIndex := moveRegex.SubexpIndex("capture")
	dstIndex := moveRegex.SubexpIndex("dst")
	promoIndex := moveRegex.SubexpIndex("promotionTo")

	src := boardMatrixStoI[strings.ToUpper(matches[srcIndex])]
	dst := boardMatrixStoI[strings.ToUpper(matches[dstIndex])]

	var piece string

	switch strings.ToUpper(matches[pieceIndex]) {
	case "":
		piece = "P"
	default:
		piece = strings.ToUpper(matches[pieceIndex])
	}

	return &Move{
		piece:       piece,
		srcSquare:   src,
		dstSquare:   dst,
		capture:     strings.ToUpper(matches[capIndex]) == "X",
		promotionTo: strings.ToUpper(matches[promoIndex]),
	}, nil
}

type Position struct {
	Pieces          *PiecePositions
	EnPassantTarget Square
	SideToMove      uint8
	CastlingRights  uint8
	HalfmoveClock   uint8
	FullMoveCounter uint8
}

// nolint:funlen,cyclop // this case statement pushes over the limit but is largely
// unavoidable.
// complexity in this function is as simple as can be made.
func (p *PiecePositions) setPieces(pieces string) *PiecePositionError {
	// restet all positions
	p.BlackRook.board &= uint64(0)
	p.BlackKnight.board &= uint64(0)
	p.BlackBishop.board &= uint64(0)
	p.BlackQueen.board &= uint64(0)
	p.BlackKing.board &= uint64(0)
	p.BlackPawn.board &= uint64(0)
	p.WhiteRook.board &= uint64(0)
	p.WhiteKnight.board &= uint64(0)
	p.WhiteBishop.board &= uint64(0)
	p.WhiteQueen.board &= uint64(0)
	p.WhiteKing.board &= uint64(0)
	p.WhitePawn.board &= uint64(0)

	offset := 0

	for index, pos := range strings.ReplaceAll(pieces, "/", "") {
		switch pos {
		case 'r':
			{
				p.BlackRook.FlipBit(Square(1 << (index + offset)))
			}
		case 'n':
			{
				p.BlackKnight.FlipBit(Square(1 << (index + offset)))
			}
		case 'b':
			{
				p.BlackBishop.FlipBit(Square(1 << (index + offset)))
			}
		case 'q':
			{
				p.BlackQueen.FlipBit(Square(1 << (index + offset)))
			}
		case 'k':
			{
				p.BlackKing.FlipBit(Square(1 << (index + offset)))
			}
		case 'p':
			{
				p.BlackPawn.FlipBit(Square(1 << (index + offset)))
			}
		case 'R':
			{
				p.WhiteRook.FlipBit(Square(1 << (index + offset)))
			}
		case 'N':
			{
				p.WhiteKnight.FlipBit(Square(1 << (index + offset)))
			}
		case 'B':
			{
				p.WhiteBishop.FlipBit(Square(1 << (index + offset)))
			}
		case 'Q':
			{
				p.WhiteQueen.FlipBit(Square(1 << (index + offset)))
			}
		case 'K':
			{
				p.WhiteKing.FlipBit(Square(1 << (index + offset)))
			}
		case 'P':
			{
				p.WhitePawn.FlipBit(Square(1 << (index + offset)))
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
	enPassantMatrix := map[string]Square{
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

// nolint:funlen // Cannot really make this any simpler
func (p *Position) SetPositionFromFen(fen string) error {
	fenElements := strings.Split(fen, " ")

	if len(fenElements) != numFenElements {
		return &InvalidFenstringError{
			fen: fen,
			err: "Missing Fen elements",
		}
	}

	pieceErr := p.Pieces.setPieces(fenElements[0])
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
		fen += boardMatrixItoS[p.EnPassantTarget]
	}

	fen += " "

	fen += fmt.Sprintf("%d %d", p.HalfmoveClock, p.FullMoveCounter)

	return fen
}

func (p *Position) IsValidMove(move *Move) bool {
	return false
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
