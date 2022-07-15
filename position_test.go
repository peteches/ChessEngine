package main

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	. "github.com/smartystreets/goconvey/convey"
)

// nolint:gochecknoglobals // this is for testing purposes
var DEBUG = false

// nolint:gochecknoglobals // this is for testing purposes
var validPiecePositions = map[string]PiecePositions{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR": {
		WhiteKing:   NewBitboard(E1),
		WhiteQueen:  NewBitboard(D1),
		WhiteBishop: NewBitboard(C1, F1),
		WhiteKnight: NewBitboard(B1, G1),
		WhiteRook:   NewBitboard(A1, H1),
		WhitePawn:   NewBitboard(A2, B2, C2, D2, E2, F2, G2, H2),
		BlackKing:   NewBitboard(E8),
		BlackQueen:  NewBitboard(D8),
		BlackBishop: NewBitboard(C8, F8),
		BlackKnight: NewBitboard(B8, G8),
		BlackRook:   NewBitboard(A8, H8),
		BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR": {
		WhiteKing:   NewBitboard(E1),
		WhiteQueen:  NewBitboard(D1),
		WhiteBishop: NewBitboard(C1, F1),
		WhiteKnight: NewBitboard(B1, G1),
		WhiteRook:   NewBitboard(A1, H1),
		WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
		BlackKing:   NewBitboard(E8),
		BlackQueen:  NewBitboard(D8),
		BlackBishop: NewBitboard(C8, F8),
		BlackKnight: NewBitboard(B8, G8),
		BlackRook:   NewBitboard(A8, H8),
		BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
	},
	"rnbqkbnr/ppp1pppp/3p4/8/4P3/8/PPPP1PPP/RNBQKBNR": {
		WhiteKing:   NewBitboard(E1),
		WhiteQueen:  NewBitboard(D1),
		WhiteBishop: NewBitboard(C1, F1),
		WhiteKnight: NewBitboard(B1, G1),
		WhiteRook:   NewBitboard(A1, H1),
		WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
		BlackKing:   NewBitboard(E8),
		BlackQueen:  NewBitboard(D8),
		BlackBishop: NewBitboard(C8, F8),
		BlackKnight: NewBitboard(B8, G8),
		BlackRook:   NewBitboard(A8, H8),
		BlackPawn:   NewBitboard(A7, B7, C7, D6, E7, F7, G7, H7),
	},
}

// nolint:gochecknoglobals // this is for testing purposes
var validFenstrings = map[string]Position{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E2, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      WHITE,
		CastlingRights:  0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: A3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: B3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: C3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: D3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: E3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: F3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: G3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H3 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: H3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: A6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: B6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: C6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: D6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: E6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: F6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: G6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": {
		Pieces: &PiecePositions{
			WhiteKing:   NewBitboard(E1),
			WhiteQueen:  NewBitboard(D1),
			WhiteBishop: NewBitboard(C1, F1),
			WhiteKnight: NewBitboard(B1, G1),
			WhiteRook:   NewBitboard(A1, H1),
			WhitePawn:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			BlackKing:   NewBitboard(E8),
			BlackQueen:  NewBitboard(D8),
			BlackBishop: NewBitboard(C8, F8),
			BlackKnight: NewBitboard(B8, G8),
			BlackRook:   NewBitboard(A8, H8),
			BlackPawn:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingRights:  0,
		EnPassantTarget: H6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
}

// nolint:gochecknoglobals // this is for testing purposes
var invalidFenstrings = map[string]error{
	"rnbfkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": &PiecePositionError{
		fen:      "rnbfkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4",
		errPiece: 'f',
	},
	"rnbqkbnr/pppppppp/9/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": &PiecePositionError{
		fen:      "rnbqkbnr/pppppppp/9/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4",
		errPiece: '9',
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR x - H6 4 4": &SideToMoveError{
		fen:     "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR x - H6 4 4",
		errSide: "x",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - H2 4 4": &EnPassantTargetError{
		fen:       "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - H2 4 4",
		errTarget: "H2",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 e 4": &HalfMoveClockError{
		fen:           "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 e 4",
		halfMoveClock: "e",
		err:           "strconv.ParseUint: parsing \"e\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1 4": &HalfMoveClockError{
		fen:           "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1 4",
		halfMoveClock: "-1",
		err:           "strconv.ParseUint: parsing \"-1\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 1 -1": &FullMoveCounterError{
		fen:             "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 1 -1",
		fullMoveCounter: "-1",
		err:             "strconv.ParseUint: parsing \"-1\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR": &InvalidFenstringError{
		fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR",
		err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w": &InvalidFenstringError{
		fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w",
		err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w -": &InvalidFenstringError{
		fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w -",
		err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3": &InvalidFenstringError{
		fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3",
		err: "Missing Fen elements",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1": &InvalidFenstringError{
		fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1",
		err: "Missing Fen elements",
	},
}

// nolint:funlen // convey testing is verbose
func TestConstants(t *testing.T) {
	if DEBUG == true {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	Convey("Constants Set Correctly", t, func() {
		So(A1, ShouldEqual, Square(1))
		So(B1, ShouldEqual, Square(1<<1))
		So(C1, ShouldEqual, Square(1<<2))
		So(D1, ShouldEqual, Square(1<<3))
		So(E1, ShouldEqual, Square(1<<4))
		So(F1, ShouldEqual, Square(1<<5))
		So(G1, ShouldEqual, Square(1<<6))
		So(H1, ShouldEqual, Square(1<<7))
		So(A2, ShouldEqual, Square(1<<8))
		So(B2, ShouldEqual, Square(1<<9))
		So(C2, ShouldEqual, Square(1<<10))
		So(D2, ShouldEqual, Square(1<<11))
		So(E2, ShouldEqual, Square(1<<12))
		So(F2, ShouldEqual, Square(1<<13))
		So(G2, ShouldEqual, Square(1<<14))
		So(H2, ShouldEqual, Square(1<<15))
		So(A3, ShouldEqual, Square(1<<16))
		So(B3, ShouldEqual, Square(1<<17))
		So(C3, ShouldEqual, Square(1<<18))
		So(D3, ShouldEqual, Square(1<<19))
		So(E3, ShouldEqual, Square(1<<20))
		So(F3, ShouldEqual, Square(1<<21))
		So(G3, ShouldEqual, Square(1<<22))
		So(H3, ShouldEqual, Square(1<<23))
		So(A4, ShouldEqual, Square(1<<24))
		So(B4, ShouldEqual, Square(1<<25))
		So(C4, ShouldEqual, Square(1<<26))
		So(D4, ShouldEqual, Square(1<<27))
		So(E4, ShouldEqual, Square(1<<28))
		So(F4, ShouldEqual, Square(1<<29))
		So(G4, ShouldEqual, Square(1<<30))
		So(H4, ShouldEqual, Square(1<<31))
		So(A5, ShouldEqual, Square(1<<32))
		So(B5, ShouldEqual, Square(1<<33))
		So(C5, ShouldEqual, Square(1<<34))
		So(D5, ShouldEqual, Square(1<<35))
		So(E5, ShouldEqual, Square(1<<36))
		So(F5, ShouldEqual, Square(1<<37))
		So(G5, ShouldEqual, Square(1<<38))
		So(H5, ShouldEqual, Square(1<<39))
		So(A6, ShouldEqual, Square(1<<40))
		So(B6, ShouldEqual, Square(1<<41))
		So(C6, ShouldEqual, Square(1<<42))
		So(D6, ShouldEqual, Square(1<<43))
		So(E6, ShouldEqual, Square(1<<44))
		So(F6, ShouldEqual, Square(1<<45))
		So(G6, ShouldEqual, Square(1<<46))
		So(H6, ShouldEqual, Square(1<<47))
		So(A7, ShouldEqual, Square(1<<48))
		So(B7, ShouldEqual, Square(1<<49))
		So(C7, ShouldEqual, Square(1<<50))
		So(D7, ShouldEqual, Square(1<<51))
		So(E7, ShouldEqual, Square(1<<52))
		So(F7, ShouldEqual, Square(1<<53))
		So(G7, ShouldEqual, Square(1<<54))
		So(H7, ShouldEqual, Square(1<<55))
		So(A8, ShouldEqual, Square(1<<56))
		So(B8, ShouldEqual, Square(1<<57))
		So(C8, ShouldEqual, Square(1<<58))
		So(D8, ShouldEqual, Square(1<<59))
		So(E8, ShouldEqual, Square(1<<60))
		So(F8, ShouldEqual, Square(1<<61))
		So(G8, ShouldEqual, Square(1<<62))
		So(H8, ShouldEqual, Square(1<<63))
	})
}

func TestSquare(t *testing.T) {
	Convey("Given a Square type", t, func() {
		// nolint:dupl // the Rank() and File() tests *could* be the same
		// function but I prefer them separate for clarity.
		Convey("It should have a File() method that reveals which file the square is in", func() {
			testCases := map[Square]uint8{
				A1: 1, A2: 1, A3: 1, A4: 1, A5: 1, A6: 1, A7: 1, A8: 1,
				B1: 2, B2: 2, B3: 2, B4: 2, B5: 2, B6: 2, B7: 2, B8: 2,
				C1: 3, C2: 3, C3: 3, C4: 3, C5: 3, C6: 3, C7: 3, C8: 3,
				D1: 4, D2: 4, D3: 4, D4: 4, D5: 4, D6: 4, D7: 4, D8: 4,
				E1: 5, E2: 5, E3: 5, E4: 5, E5: 5, E6: 5, E7: 5, E8: 5,
				F1: 6, F2: 6, F3: 6, F4: 6, F5: 6, F6: 6, F7: 6, F8: 6,
				G1: 7, G2: 7, G3: 7, G4: 7, G5: 7, G6: 7, G7: 7, G8: 7,
				H1: 8, H2: 8, H3: 8, H4: 8, H5: 8, H6: 8, H7: 8, H8: 8,
			}
			for sqr, expectedFile := range testCases {
				So(sqr.File(), ShouldEqual, expectedFile)
				So(sqr.File(), ShouldBeLessThan, expectedFile+1)
				So(sqr.File(), ShouldBeGreaterThan, expectedFile-1)
			}
		})
		// nolint:dupl // the Rank() and File() tests *could* be the same
		// function but I prefer them separate for clarity.
		Convey("It should have a Rank() method that reveals which rank the square is in", func() {
			testCases := map[Square]uint8{
				A1: 1, A2: 2, A3: 3, A4: 4, A5: 5, A6: 6, A7: 7, A8: 8,
				B1: 1, B2: 2, B3: 3, B4: 4, B5: 5, B6: 6, B7: 7, B8: 8,
				C1: 1, C2: 2, C3: 3, C4: 4, C5: 5, C6: 6, C7: 7, C8: 8,
				D1: 1, D2: 2, D3: 3, D4: 4, D5: 5, D6: 6, D7: 7, D8: 8,
				E1: 1, E2: 2, E3: 3, E4: 4, E5: 5, E6: 6, E7: 7, E8: 8,
				F1: 1, F2: 2, F3: 3, F4: 4, F5: 5, F6: 6, F7: 7, F8: 8,
				G1: 1, G2: 2, G3: 3, G4: 4, G5: 5, G6: 6, G7: 7, G8: 8,
				H1: 1, H2: 2, H3: 3, H4: 4, H5: 5, H6: 6, H7: 7, H8: 8,
			}
			for sqr, expectedRank := range testCases {
				So(sqr.Rank(), ShouldEqual, expectedRank)
				So(sqr.Rank(), ShouldBeLessThan, expectedRank+1)
				So(sqr.Rank(), ShouldBeGreaterThan, expectedRank-1)
			}
		})
	})
}

// nolint:funlen // convey testing is verbose
func TestBitboard(t *testing.T) {
	Convey("Given a NewBitboard function", t, func() {
		Convey("With no arguments", func() {
			Convey("It should return an empty board", func() {
				bb := NewBitboard()
				So(*bb, ShouldResemble, BITBOARD{})
				So(bb.board, ShouldEqual, 0)
			})

			Convey("With args", func() {
				Convey("It should return an initialised board", func() {
					bitBoard := NewBitboard(A8)
					So(bitBoard.board, ShouldEqual, A8)
					bitBoard = NewBitboard(E3)
					So(bitBoard.board, ShouldEqual, E3)
					bitBoard = NewBitboard(A2, B2)
					So(bitBoard.board, ShouldEqual, A2+B2)
					bitBoard = NewBitboard(A8, H1)
					So(bitBoard.board, ShouldEqual, A8+H1)
				})
			})
		})
	})
	Convey("Given an existing BitBoard", t, func() {
		bitboard := NewBitboard()
		Convey("Bit manipulation basics", func() {
			So(0^(1<<0), ShouldEqual, 1)
		})
		Convey("When FlipBit method called", func() {
			Convey("It should update its board attribute", func() {
				So(bitboard.board, ShouldEqual, 0)
				for _, sqr := range allSquares {
					bb := NewBitboard()
					So(bb.board, ShouldEqual, 0)
					bb.FlipBit(sqr)
					So(bb.board, ShouldEqual, sqr)
				}
			})
		})
		Convey("When Squares() method called returns list strings where bits are 1", func() {
			So(bitboard.Squares(), ShouldResemble, []string{})
			bitboard.FlipBit(A8)
			So(bitboard.Squares(), ShouldHaveLength, 1)
			So(bitboard.Squares(), ShouldContain, "A8")
			bitboard.FlipBit(H3)
			bitboard.FlipBit(H4)
			So(bitboard.Squares(), ShouldHaveLength, 3)
			So(bitboard.Squares(), ShouldContain, "A8")
			So(bitboard.Squares(), ShouldContain, "H4")
			So(bitboard.Squares(), ShouldContain, "H3")
		})

		Convey("When Occupied() method called with square, returns true if square occupied", func() {
			for _, sqr := range allSquares {
				So(bitboard.Occupied(sqr), ShouldEqual, false)
			}
			bitboard.FlipBit(B4)
			So(bitboard.Occupied(B4), ShouldEqual, true)
			bitboard.FlipBit(A8)
			So(bitboard.Occupied(B4), ShouldEqual, true)
			So(bitboard.Occupied(A8), ShouldEqual, true)
			So(bitboard.Occupied(H7), ShouldEqual, false)
			bitboard.FlipBit(H7)
			So(bitboard.Occupied(H7), ShouldEqual, true)
		})
	})
}

// nolint:funlen // convey testing is verbose
func TestPiecePositions(t *testing.T) {
	Convey("Given a PiecePositions struct", t, func() {
		pieces := NewPiecePositions()
		Convey("It should have BITBOARD fields for all types of piece", func() {
			So(pieces.WhiteKing, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.WhiteQueen, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.WhiteKnight, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.WhiteBishop, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.WhiteRook, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.WhitePawn, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.BlackKing, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.BlackQueen, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.BlackKnight, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.BlackBishop, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.BlackRook, ShouldHaveSameTypeAs, &BITBOARD{})
			So(pieces.BlackPawn, ShouldHaveSameTypeAs, &BITBOARD{})
		})
		Convey("the Occupied() method should return true if any piece is in square", func() {
			for _, sqr := range allSquares {
				So(pieces.Occupied(sqr), ShouldEqual, false)
			}
			pieces.BlackBishop.FlipBit(A8)
			So(pieces.Occupied(A8), ShouldEqual, true)
			pieces.WhiteQueen.FlipBit(E5)
			So(pieces.Occupied(E5), ShouldEqual, true)
			pieces.WhiteQueen.FlipBit(H1)
			So(pieces.Occupied(H1), ShouldEqual, true)
		})
		Convey("the OccupiedBy() method should return which Piece is occupying the square", func() {
			pieces.BlackBishop.FlipBit(A8)
			So(pieces.OccupiedBy(A8), ShouldEqual, "b")
			pieces.BlackBishop.FlipBit(F3)
			So(pieces.OccupiedBy(F3), ShouldEqual, "b")
			pieces.WhiteQueen.FlipBit(H3)
			So(pieces.OccupiedBy(H3), ShouldEqual, "Q")
			pieces.WhiteRook.FlipBit(H1)
			So(pieces.OccupiedBy(H1), ShouldEqual, "R")
		})
		Convey("the String() method returns the string representation of the pieces ala fen notation", func() {
			pieces := NewPiecePositions()
			So(pieces.String(), ShouldEqual, "8/8/8/8/8/8/8/8")
			pieces.BlackRook.FlipBit(A8)
			pieces.BlackKnight.FlipBit(B8)
			pieces.BlackBishop.FlipBit(C8)
			pieces.BlackQueen.FlipBit(D8)
			pieces.BlackKing.FlipBit(E8)
			pieces.BlackBishop.FlipBit(F8)
			pieces.BlackKnight.FlipBit(G8)
			pieces.BlackRook.FlipBit(H8)
			pieces.BlackPawn.FlipBit(A7)
			pieces.BlackPawn.FlipBit(B7)
			pieces.BlackPawn.FlipBit(C7)
			pieces.BlackPawn.FlipBit(D7)
			pieces.BlackPawn.FlipBit(E7)
			pieces.BlackPawn.FlipBit(F7)
			pieces.BlackPawn.FlipBit(G7)
			pieces.BlackPawn.FlipBit(H7)
			pieces.WhitePawn.FlipBit(A2)
			pieces.WhitePawn.FlipBit(B2)
			pieces.WhitePawn.FlipBit(C2)
			pieces.WhitePawn.FlipBit(D2)
			pieces.WhitePawn.FlipBit(E2)
			pieces.WhitePawn.FlipBit(F2)
			pieces.WhitePawn.FlipBit(G2)
			pieces.WhitePawn.FlipBit(H2)
			pieces.WhiteRook.FlipBit(A1)
			pieces.WhiteKnight.FlipBit(B1)
			pieces.WhiteBishop.FlipBit(C1)
			pieces.WhiteQueen.FlipBit(D1)
			pieces.WhiteKing.FlipBit(E1)
			pieces.WhiteBishop.FlipBit(F1)
			pieces.WhiteKnight.FlipBit(G1)
			pieces.WhiteRook.FlipBit(H1)
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

// nolint:funlen // Convey testing is verbose
func TestMoveStruct(t *testing.T) {
	Convey("Given a Move struct", t, func() {
		Convey("it should have relevant fields", func() {
			move := Move{}
			So(move.srcSquare, ShouldHaveSameTypeAs, Square(0))
			So(move.dstSquare, ShouldHaveSameTypeAs, Square(0))
			So(move.capture, ShouldHaveSameTypeAs, true)
			So(move.promotionTo, ShouldHaveSameTypeAs, "")
			So(move.piece, ShouldHaveSameTypeAs, "")
		})
		Convey("There should be a NewMove() func", func() {
			Convey("It should accept a LAN encoded move", func() {
				testCases := map[string]*Move{
					"e2e4": {
						piece:       "P",
						srcSquare:   E2,
						dstSquare:   E4,
						capture:     false,
						promotionTo: "",
					},
					"e2-e4": {
						piece:       "P",
						srcSquare:   E2,
						dstSquare:   E4,
						capture:     false,
						promotionTo: "",
					},
					"e2xe3": {
						piece:       "P",
						srcSquare:   E2,
						dstSquare:   E3,
						capture:     true,
						promotionTo: "",
					},
					"e7e8Q": {
						srcSquare:   E7,
						piece:       "P",
						dstSquare:   E8,
						capture:     false,
						promotionTo: "Q",
					},
					"E2E4": {
						piece:       "P",
						srcSquare:   E2,
						dstSquare:   E4,
						capture:     false,
						promotionTo: "",
					},
					"E2-E4": {
						piece:       "P",
						srcSquare:   E2,
						dstSquare:   E4,
						capture:     false,
						promotionTo: "",
					},
					"E2XE3": {
						srcSquare:   E2,
						piece:       "P",
						dstSquare:   E3,
						capture:     true,
						promotionTo: "",
					},
					"E7E8q": {
						piece:       "P",
						srcSquare:   E7,
						dstSquare:   E8,
						capture:     false,
						promotionTo: "Q",
					},
					"ke3-f5": {
						piece:       "K",
						srcSquare:   E3,
						dstSquare:   F5,
						capture:     false,
						promotionTo: "",
					},
					"qH3xe5": {
						piece:       "Q",
						srcSquare:   H3,
						dstSquare:   E5,
						capture:     true,
						promotionTo: "",
					},
				}
				for move, expectedMove := range testCases {
					m, err := NewMove(move)
					So(err, ShouldEqual, nil)
					So(m, ShouldResemble, expectedMove)
				}
			})
			SkipConvey("It should return an error if the move is invalid", func() {
				testCases := map[string]MoveError{
					"e2e7": {
						fen:  "",
						move: "e2e7",
						err:  "Pawns do not move like that",
					},
				}
				for move, err := range testCases {
					m, mError := NewMove(move)
					So(mError, ShouldResemble, err)
					So(m, ShouldEqual, nil)
				}
			})
		})
	})
}

// nolint:funlen // Convey testing is verbose
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
			enpassantTargets := map[string]Square{
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
			for target, expectedTarget := range enpassantTargets {
				enpassantErr := pos.setEnPassantTarget(target)
				So(enpassantErr, ShouldEqual, nil)
				So(pos.EnPassantTarget, ShouldEqual, expectedTarget)
			}
			enpassantErr := pos.setEnPassantTarget("H1")
			So(enpassantErr, ShouldResemble, &EnPassantTargetError{errTarget: "H1"})
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
			Convey("Return a PositionError if given position is invalid", func() {
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
						m, _ := NewMove(move)
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
