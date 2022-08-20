package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/errors"
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

type Position struct {
	Board           *board.Board
	EnPassantTarget board.Square
	SideToMove      uint8
	CastlingRights  uint8
	HalfmoveClock   uint8
	FullMoveCounter uint8
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

	pieceErr := p.Board.SetPieces(fenElements[0])
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

	fen += p.Board.String()
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
		fen += p.EnPassantTarget.String()
	}

	fen += " "

	fen += fmt.Sprintf("%d %d", p.HalfmoveClock, p.FullMoveCounter)

	return fen
}

func NewPosition() *Position {
	pos := Position{
		Board:           board.NewBoard(),
		SideToMove:      WHITE,
		CastlingRights:  0,
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	}

	return &pos
}
