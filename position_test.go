package main

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/errors"
	"github.com/peteches/ChessEngine/moves"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:gochecknoglobals // this is for testing purposes
var validPiecePositions = map[string]PiecePositions{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR": {
		WhiteKing:   board.NewBitboard(board.E1),
		WhiteQueen:  board.NewBitboard(board.D1),
		WhiteBishop: board.NewBitboard(board.C1, board.F1),
		WhiteKnight: board.NewBitboard(board.B1, board.G1),
		WhiteRook:   board.NewBitboard(board.A1, board.H1),
		WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E2, board.F2, board.G2, board.H2),
		BlackKing:   board.NewBitboard(board.E8),
		BlackQueen:  board.NewBitboard(board.D8),
		BlackBishop: board.NewBitboard(board.C8, board.F8),
		BlackKnight: board.NewBitboard(board.B8, board.G8),
		BlackRook:   board.NewBitboard(board.A8, board.H8),
		BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR": {
		WhiteKing:   board.NewBitboard(board.E1),
		WhiteQueen:  board.NewBitboard(board.D1),
		WhiteBishop: board.NewBitboard(board.C1, board.F1),
		WhiteKnight: board.NewBitboard(board.B1, board.G1),
		WhiteRook:   board.NewBitboard(board.A1, board.H1),
		WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
		BlackKing:   board.NewBitboard(board.E8),
		BlackQueen:  board.NewBitboard(board.D8),
		BlackBishop: board.NewBitboard(board.C8, board.F8),
		BlackKnight: board.NewBitboard(board.B8, board.G8),
		BlackRook:   board.NewBitboard(board.A8, board.H8),
		BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
	},
	"rnbqkbnr/ppp1pppp/3p4/8/4P3/8/PPPP1PPP/RNBQKBNR": {
		WhiteKing:   board.NewBitboard(board.E1),
		WhiteQueen:  board.NewBitboard(board.D1),
		WhiteBishop: board.NewBitboard(board.C1, board.F1),
		WhiteKnight: board.NewBitboard(board.B1, board.G1),
		WhiteRook:   board.NewBitboard(board.A1, board.H1),
		WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
		BlackKing:   board.NewBitboard(board.E8),
		BlackQueen:  board.NewBitboard(board.D8),
		BlackBishop: board.NewBitboard(board.C8, board.F8),
		BlackKnight: board.NewBitboard(board.B8, board.G8),
		BlackRook:   board.NewBitboard(board.A8, board.H8),
		BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D6, board.E7, board.F7, board.G7, board.H7),
	},
}

//nolint:gochecknoglobals // this is for testing purposes
var validFenstrings = map[string]Position{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E2, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      WHITE,
		CastlingRights:  0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.A3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.B3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.C3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.D3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.E3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.F3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.G3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.H3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.A6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.B6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.C6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.D6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.E6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.F6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.G6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   board.NewBitboard(board.E1),
			WhiteQueen:  board.NewBitboard(board.D1),
			WhiteBishop: board.NewBitboard(board.C1, board.F1),
			WhiteKnight: board.NewBitboard(board.B1, board.G1),
			WhiteRook:   board.NewBitboard(board.A1, board.H1),
			WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:   board.NewBitboard(board.E8),
			BlackQueen:  board.NewBitboard(board.D8),
			BlackBishop: board.NewBitboard(board.C8, board.F8),
			BlackKnight: board.NewBitboard(board.B8, board.G8),
			BlackRook:   board.NewBitboard(board.A8, board.H8),
			BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.H6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
}

//nolint:gochecknoglobals // this is for testing purposes
var invalidFenstrings = map[string]error{
	"rnbfkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": &errors.PiecePositionError{
		Fen:      "rnbfkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4",
		ErrPiece: 'f',
	},
	"rnbqkbnr/pppppppp/9/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": &errors.PiecePositionError{
		Fen:      "rnbqkbnr/pppppppp/9/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4",
		ErrPiece: '9',
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR x - H6 4 4": &errors.SideToMoveError{
		Fen:     "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR x - H6 4 4",
		ErrSide: "x",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - H2 4 4": &errors.EnPassantTargetError{
		Fen:       "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - H2 4 4",
		ErrTarget: "H2",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 e 4": &errors.HalfMoveClockError{
		Fen:           "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 e 4",
		HalfMoveClock: "e",
		Err:           "strconv.ParseUint: parsing \"e\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1 4": &errors.HalfMoveClockError{
		Fen:           "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1 4",
		HalfMoveClock: "-1",
		Err:           "strconv.ParseUint: parsing \"-1\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 1 -1": &errors.FullMoveCounterError{
		Fen:             "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 1 -1",
		FullMoveCounter: "-1",
		Err:             "strconv.ParseUint: parsing \"-1\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR": &errors.InvalidFenstringError{
		Fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR",
		Err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w": &errors.InvalidFenstringError{
		Fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w",
		Err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w -": &errors.InvalidFenstringError{
		Fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w -",
		Err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3": &errors.InvalidFenstringError{
		Fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3",
		Err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1": &errors.InvalidFenstringError{
		Fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1",
		Err: "Missing Fen elements",
	},
}

//nolint:funlen // convey testing is verbose
func TestPiecePositions(t *testing.T) {
	Convey("Given a PiecePositions struct", t, func() {
		pieces := NewPiecePositions()
		Convey("It should have BITBOARD fields for all types of piece", func() {
			So(pieces.WhiteKing, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.WhiteQueen, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.WhiteKnight, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.WhiteBishop, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.WhiteRook, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.WhitePawn, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.BlackKing, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.BlackQueen, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.BlackKnight, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.BlackBishop, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.BlackRook, ShouldHaveSameTypeAs, &board.BitBoard{})
			So(pieces.BlackPawn, ShouldHaveSameTypeAs, &board.BitBoard{})
		})
		Convey("the Occupied() method should return true if any piece is in square", func() {
			for _, sqr := range board.AllSquares {
				So(pieces.Occupied(sqr), ShouldEqual, false)
			}
			pieces.BlackBishop.FlipBit(board.A8)
			So(pieces.Occupied(board.A8), ShouldEqual, true)
			pieces.WhiteQueen.FlipBit(board.E5)
			So(pieces.Occupied(board.E5), ShouldEqual, true)
			pieces.WhiteQueen.FlipBit(board.H1)
			So(pieces.Occupied(board.H1), ShouldEqual, true)
		})
		Convey("the OccupiedBy() method should return which Piece is occupying the square", func() {
			pieces.BlackBishop.FlipBit(board.A8)
			So(pieces.OccupiedBy(board.A8), ShouldEqual, "b")
			pieces.BlackBishop.FlipBit(board.F3)
			So(pieces.OccupiedBy(board.F3), ShouldEqual, "b")
			pieces.WhiteQueen.FlipBit(board.H3)
			So(pieces.OccupiedBy(board.H3), ShouldEqual, "Q")
			pieces.WhiteRook.FlipBit(board.H1)
			So(pieces.OccupiedBy(board.H1), ShouldEqual, "R")
		})
		Convey("the String() method returns the string representation of the pieces ala fen notation", func() {
			pieces := NewPiecePositions()
			So(pieces.String(), ShouldEqual, "8/8/8/8/8/8/8/8")
			pieces.BlackRook.FlipBit(board.A8)
			pieces.BlackKnight.FlipBit(board.B8)
			pieces.BlackBishop.FlipBit(board.C8)
			pieces.BlackQueen.FlipBit(board.D8)
			pieces.BlackKing.FlipBit(board.E8)
			pieces.BlackBishop.FlipBit(board.F8)
			pieces.BlackKnight.FlipBit(board.G8)
			pieces.BlackRook.FlipBit(board.H8)
			pieces.BlackPawn.FlipBit(board.A7)
			pieces.BlackPawn.FlipBit(board.B7)
			pieces.BlackPawn.FlipBit(board.C7)
			pieces.BlackPawn.FlipBit(board.D7)
			pieces.BlackPawn.FlipBit(board.E7)
			pieces.BlackPawn.FlipBit(board.F7)
			pieces.BlackPawn.FlipBit(board.G7)
			pieces.BlackPawn.FlipBit(board.H7)
			pieces.WhitePawn.FlipBit(board.A2)
			pieces.WhitePawn.FlipBit(board.B2)
			pieces.WhitePawn.FlipBit(board.C2)
			pieces.WhitePawn.FlipBit(board.D2)
			pieces.WhitePawn.FlipBit(board.E2)
			pieces.WhitePawn.FlipBit(board.F2)
			pieces.WhitePawn.FlipBit(board.G2)
			pieces.WhitePawn.FlipBit(board.H2)
			pieces.WhiteRook.FlipBit(board.A1)
			pieces.WhiteKnight.FlipBit(board.B1)
			pieces.WhiteBishop.FlipBit(board.C1)
			pieces.WhiteQueen.FlipBit(board.D1)
			pieces.WhiteKing.FlipBit(board.E1)
			pieces.WhiteBishop.FlipBit(board.F1)
			pieces.WhiteKnight.FlipBit(board.G1)
			pieces.WhiteRook.FlipBit(board.H1)
			So(pieces.String(), ShouldEqual, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
		})
		Convey("The setPieces method should update Pieces bitboards", func() {
			for piecePositions, expectedPiecePosition := range validPiecePositions {
				err := pieces.setPieces(piecePositions)
				So(err, ShouldEqual, nil)
				So(*pieces, ShouldResemble, expectedPiecePosition)
			}
		})
	})
}

//nolint:funlen // Convey testing is verbose
func TestPosition(t *testing.T) {
	Convey("Given a NewPosition", t, func() {
		pos := NewPosition()
		Convey("returns a NewPosition struct", func() {
			So(pos, ShouldHaveSameTypeAs, &Position{})
			So(pos, ShouldResemble, &Position{
				Pieces:          NewPiecePositions(),
				SideToMove:      WHITE,
				CastlingRights:  0,
				EnPassantTarget: 0,
				HalfmoveClock:   0,
				FullMoveCounter: 1,
			})
		})
		Convey("The setCastlingRights method should update position.CastlingRights field", func() {
			CastlingRights := map[string]uint8{
				"kqKQ": 0 ^ (BlackKingSideAllowed | BlackQueenSideAllowed | WhiteKingSideAllowed | WhiteQueenSideAllowed),
				"kqK":  0 ^ (BlackKingSideAllowed | BlackQueenSideAllowed | WhiteKingSideAllowed),
				"kqQ":  0 ^ (BlackKingSideAllowed | BlackQueenSideAllowed | WhiteQueenSideAllowed),
				"kq":   0 ^ (BlackKingSideAllowed | BlackQueenSideAllowed),
				"-":    0,
				"K":    0 ^ WhiteKingSideAllowed,
			}
			for k, v := range CastlingRights {
				err := pos.setCastlingRights(k)
				So(err, ShouldEqual, nil)
				So(pos.CastlingRights, ShouldEqual, v)
			}
		})
		Convey("The setEnPassantTarget method should update the EnPassantTarget field", func() {
			enpassantTargets := map[string]board.Square{
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
			for target, expectedTarget := range enpassantTargets {
				enpassantErr := pos.setEnPassantTarget(target)
				So(enpassantErr, ShouldEqual, nil)
				So(pos.EnPassantTarget, ShouldEqual, expectedTarget)
			}
			enpassantErr := pos.setEnPassantTarget("board.H1")
			So(enpassantErr, ShouldResemble, &errors.EnPassantTargetError{ErrTarget: "board.H1"})
		})
		Convey("The setSideToMove method should update the side to move field", func() {
			err := pos.setSideToMove("w")
			So(err, ShouldEqual, nil)
			So(pos.SideToMove, ShouldEqual, WHITE)
			err = pos.setSideToMove("b")
			So(err, ShouldEqual, nil)
			So(pos.SideToMove, ShouldEqual, BLACK)
		})
		Convey("The SetPositionFromFen method should ", func() {
			Convey("Accept FenString and set position accordingly", func() {
				for fen, position := range validFenstrings {
					pos := NewPosition()
					err := pos.SetPositionFromFen(fen)
					So(err, ShouldEqual, nil)
					So(*pos, ShouldResemble, position)
				}
			})
			Convey("Return a errors.PositionError if given position is invalid", func() {
				for fen, expectedErr := range invalidFenstrings {
					pos := NewPosition()
					err := pos.SetPositionFromFen(fen)
					So(err, ShouldResemble, expectedErr)
				}
			})
		})
		Convey("the String() method should accept no arguments and return the position as a fen string", func() {
			for fen := range validFenstrings {
				pos := NewPosition()
				err := pos.SetPositionFromFen(fen)
				So(err, ShouldEqual, nil)
				So(pos.String(), ShouldEqual, fen)
			}
		})
		Convey("The IsValidMove() method should", func() {
			SkipConvey("Accept a Move{} and return bool", func() {
				testCases := map[string]map[string]bool{
					"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - E3 1 4": {
						"e2e4": true,
					},
				}
				for startFen, v := range testCases {
					for move, moveIsValid := range v {
						err := pos.SetPositionFromFen(startFen)
						So(err, ShouldEqual, nil)
						m, _ := moves.NewMove(move)
						So(pos.IsValidMove(m), ShouldEqual, moveIsValid)
					}
				}
			})
		})
	})
}

func FuzzSetPositionFromFen(f *testing.F) {
	for fen := range validFenstrings {
		f.Add(fen)
	}

	f.Fuzz(func(t *testing.T, fen string) {
		pos := NewPosition()
		err := pos.SetPositionFromFen(fen)
		if err != nil {
			t.Fail()
		}
	})
}
