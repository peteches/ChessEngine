package main

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/errors"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:gochecknoglobals // this is for testing purposes
var validFenstrings = map[string]Position{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E2, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      WHITE,
		CastlingRights:  0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.A3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.B3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.C3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.D3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.E3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.F3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.G3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H3 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.H3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.A6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.B6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.C6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.D6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.E6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.F6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: board.G6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": {
		Board: &board.Board{
			WhiteKing:    board.NewKing(board.White, board.E1),
			WhiteQueens:  board.NewQueens(board.White, board.D1),
			WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
			WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
			WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
			WhitePawns:   board.NewPawns(board.White, board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
			BlackKing:    board.NewKing(board.Black, board.E8),
			BlackQueens:  board.NewQueens(board.Black, board.D8),
			BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
			BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
			BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
			BlackPawns:   board.NewPawns(board.Black, board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
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

//nolint:funlen // Convey testing is verbose
func TestPosition(t *testing.T) {
	Convey("Given a NewPosition", t, func() {
		pos := NewPosition()
		Convey("returns a NewPosition struct", func() {
			So(pos, ShouldHaveSameTypeAs, &Position{})
			So(pos, ShouldResemble, &Position{
				Board:           board.NewBoard(),
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
