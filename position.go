package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/errors"
	"github.com/peteches/ChessEngine/moves"
)

const numFenElements = 6

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

type PiecePositions struct {
	WhiteKing   *board.BitBoard
	WhiteQueen  *board.BitBoard
	WhiteBishop *board.BitBoard
	WhiteKnight *board.BitBoard
	WhiteRook   *board.BitBoard
	WhitePawn   *board.BitBoard
	BlackKing   *board.BitBoard
	BlackQueen  *board.BitBoard
	BlackBishop *board.BitBoard
	BlackKnight *board.BitBoard
	BlackRook   *board.BitBoard
	BlackPawn   *board.BitBoard
}

func NewPiecePositions() *PiecePositions {
	return &PiecePositions{
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
		board.NewBitboard(),
	}
}

//nolint:cyclop // not sure how to simplify this further
func (p *PiecePositions) Occupied(sqr board.Square) bool {
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

//nolint:cyclop // not sure how to simplify this further
func (p *PiecePositions) OccupiedBy(sqr board.Square) string {
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

	for idx, sqr := range board.AllSquares {
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

			if (idx + 1) < board.TotalSquares {
				fen += "/"
			}
		}
	}

	return fen
}

type Position struct {
	Pieces          *PiecePositions
	EnPassantTarget board.Square
	SideToMove      uint8
	CastlingRights  uint8
	HalfmoveClock   uint8
	FullMoveCounter uint8
}

//nolint:funlen,cyclop
func (p *PiecePositions) setPieces(pieces string) *errors.PiecePositionError {
	// restet all positions
	p.BlackRook.Board &= uint64(0)
	p.BlackKnight.Board &= uint64(0)
	p.BlackBishop.Board &= uint64(0)
	p.BlackQueen.Board &= uint64(0)
	p.BlackKing.Board &= uint64(0)
	p.BlackPawn.Board &= uint64(0)
	p.WhiteRook.Board &= uint64(0)
	p.WhiteKnight.Board &= uint64(0)
	p.WhiteBishop.Board &= uint64(0)
	p.WhiteQueen.Board &= uint64(0)
	p.WhiteKing.Board &= uint64(0)
	p.WhitePawn.Board &= uint64(0)

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
				p.BlackRook.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'n':
			{
				p.BlackKnight.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'b':
			{
				p.BlackBishop.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'q':
			{
				p.BlackQueen.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'k':
			{
				p.BlackKing.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'p':
			{
				p.BlackPawn.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'R':
			{
				p.WhiteRook.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'N':
			{
				p.WhiteKnight.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'B':
			{
				p.WhiteBishop.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'Q':
			{
				p.WhiteQueen.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'K':
			{
				p.WhiteKing.FlipBit(board.Square(1 << (index + offset)))
			}
		case 'P':
			{
				p.WhitePawn.FlipBit(board.Square(1 << (index + offset)))
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

func (p *Position) setSideToMove(side string) *errors.SideToMoveError {
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
			return &errors.SideToMoveError{ErrSide: side}
		}
	}

	return nil
}

func (p *Position) setCastlingRights(rights string) *errors.CastlingRightsError {
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
				return &errors.CastlingRightsError{ErrChar: pos}
			}
		}
	}

	p.CastlingRights = 0 ^ (blackKingSide | blackQueenSide | whiteKingSide | whiteQueenSide)

	return nil
}

func (p *Position) setEnPassantTarget(targetSquare string) *errors.EnPassantTargetError {
	enPassantMatrix := map[string]board.Square{
		"-":  0,
		"A3": board.A3,
		"B3": board.B3,
		"C3": board.C3,
		"D3": board.D3,
		"E3": board.E3,
		"F3": board.F3,
		"G3": board.G3,
		"H3": board.H3,
		"A6": board.A6,
		"B6": board.B6,
		"C6": board.C6,
		"D6": board.D6,
		"E6": board.E6,
		"F6": board.F6,
		"G6": board.G6,
		"H6": board.H6,
	}

	var ok bool
	p.EnPassantTarget, ok = enPassantMatrix[targetSquare]

	if ok {
		return nil
	}

	return &errors.EnPassantTargetError{ErrTarget: targetSquare}
}

//nolint:funlen // Cannot really make this any simpler
func (p *Position) SetPositionFromFen(fen string) error {
	fenElements := strings.Split(fen, " ")

	if len(fenElements) != numFenElements {
		return &errors.InvalidFenstringError{
			Fen: fen,
			Err: "Missing Fen elements",
		}
	}

	pieceErr := p.Pieces.setPieces(fenElements[0])
	if pieceErr != nil {
		pieceErr.Fen = fen

		return pieceErr
	}

	stmErr := p.setSideToMove(fenElements[1])
	if stmErr != nil {
		stmErr.Fen = fen

		return stmErr
	}

	castlingErr := p.setCastlingRights(fenElements[2])
	if castlingErr != nil {
		castlingErr.Fen = fen

		return castlingErr
	}

	// Setting EnPassantTarget square
	enPassantErr := p.setEnPassantTarget(fenElements[3])
	if enPassantErr != nil {
		enPassantErr.Fen = fen

		return enPassantErr
	}

	//nolint:gomnd // parse uint wants base and bits which arn't easily
	// derived
	halfMoveClock, err := strconv.ParseUint(fenElements[4], 10, 8)
	if err != nil {
		return &errors.HalfMoveClockError{
			Fen:           fen,
			Err:           err.Error(),
			HalfMoveClock: fenElements[4],
		}
	}

	p.HalfmoveClock = uint8(halfMoveClock)

	//nolint:gomnd // parse uint wants base and bits which arn't easily
	// derived
	halfMoveClock, err = strconv.ParseUint(fenElements[5], 10, 8)
	if err != nil {
		return &errors.FullMoveCounterError{
			Fen:             fen,
			Err:             err.Error(),
			FullMoveCounter: fenElements[5],
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
		fen += board.BoardMatrixItoS[p.EnPassantTarget]
	}

	fen += " "

	fen += fmt.Sprintf("%d %d", p.HalfmoveClock, p.FullMoveCounter)

	return fen
}

func (p *Position) IsValidMove(move *moves.Move) bool {
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
