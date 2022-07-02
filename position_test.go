package main

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	. "github.com/smartystreets/goconvey/convey"
)

var DEBUG = false

var VALID_FENSTRINGS = map[string]Position{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E2, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      WHITE,
		CastlingAbility: 0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0 ^ (WhiteKingSideAllowed | WhiteQueenSideAllowed | BlackKingSideAllowed | BlackQueenSideAllowed),
		EnPassantTarget: 0,
		HalfmoveClock:   0,
		FullMoveCounter: 1,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: A3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: B3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: C3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: D3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: E3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: F3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: G3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H3 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: H3,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - A6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: A6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - B6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: B6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - C6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: C6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - D6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: D6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - E6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: E6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - F6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: F6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - G6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: G6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": {
		Pieces: &PiecePositions{
			W_KING:   NewBitboard(E1),
			W_QUEEN:  NewBitboard(D1),
			W_BISHOP: NewBitboard(C1, F1),
			W_KNIGHT: NewBitboard(B1, G1),
			W_ROOK:   NewBitboard(A1, H1),
			W_PAWN:   NewBitboard(A2, B2, C2, D2, E4, F2, G2, H2),
			B_KING:   NewBitboard(E8),
			B_QUEEN:  NewBitboard(D8),
			B_BISHOP: NewBitboard(C8, F8),
			B_KNIGHT: NewBitboard(B8, G8),
			B_ROOK:   NewBitboard(A8, H8),
			B_PAWN:   NewBitboard(A7, B7, C7, D7, E7, F7, G7, H7),
		},
		SideToMove:      BLACK,
		CastlingAbility: 0,
		EnPassantTarget: H6,
		HalfmoveClock:   4,
		FullMoveCounter: 4,
	},
}

var INVALID_FENSTRINGS = map[string]error{
	"rnbfkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": &piecePositionError{
		fen:      "rnbfkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4",
		errPiece: 'f',
	},
	"rnbqkbnr/pppppppp/9/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4": &piecePositionError{
		fen:      "rnbqkbnr/pppppppp/9/8/4P3/8/PPPP1PPP/RNBQKBNR b - H6 4 4",
		errPiece: '9',
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR x - H6 4 4": &sideToMoveError{
		fen:     "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR x - H6 4 4",
		errSide: "x",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - H2 4 4": &enPassantTargetError{
		fen:       "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - H2 4 4",
		errTarget: "H2",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 e 4": &halfMoveClockError{
		fen:           "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 e 4",
		halfMoveClock: "e",
		err:           "strconv.ParseUint: parsing \"e\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1 4": &halfMoveClockError{
		fen:           "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 -1 4",
		halfMoveClock: "-1",
		err:           "strconv.ParseUint: parsing \"-1\": invalid syntax",
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 1 -1": &fullMoveCounterError{
		fen:             "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w - E3 1 -1",
		fullMoveCounter: "-1",
		err:             "strconv.ParseUint: parsing \"-1\": invalid syntax",
	},
}

func TestConstants(t *testing.T) {
	if DEBUG == true {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	Convey("Constants Set Correctly", t, func() {
		So(A8, ShouldEqual, uint64(1))
		So(B8, ShouldEqual, uint64(2))
		So(C8, ShouldEqual, uint64(3))
		So(D8, ShouldEqual, uint64(4))
		So(E8, ShouldEqual, uint64(5))
		So(F8, ShouldEqual, uint64(6))
		So(G8, ShouldEqual, uint64(7))
		So(H8, ShouldEqual, uint64(8))
		So(A7, ShouldEqual, uint64(9))
		So(B7, ShouldEqual, uint64(10))
		So(C7, ShouldEqual, uint64(11))
		So(D7, ShouldEqual, uint64(12))
		So(E7, ShouldEqual, uint64(13))
		So(F7, ShouldEqual, uint64(14))
		So(G7, ShouldEqual, uint64(15))
		So(H7, ShouldEqual, uint64(16))
		So(A6, ShouldEqual, uint64(17))
		So(B6, ShouldEqual, uint64(18))
		So(C6, ShouldEqual, uint64(19))
		So(D6, ShouldEqual, uint64(20))
		So(E6, ShouldEqual, uint64(21))
		So(F6, ShouldEqual, uint64(22))
		So(G6, ShouldEqual, uint64(23))
		So(H6, ShouldEqual, uint64(24))
		So(A5, ShouldEqual, uint64(25))
		So(B5, ShouldEqual, uint64(26))
		So(C5, ShouldEqual, uint64(27))
		So(D5, ShouldEqual, uint64(28))
		So(E5, ShouldEqual, uint64(29))
		So(F5, ShouldEqual, uint64(30))
		So(G5, ShouldEqual, uint64(31))
		So(H5, ShouldEqual, uint64(32))
		So(A4, ShouldEqual, uint64(33))
		So(B4, ShouldEqual, uint64(34))
		So(C4, ShouldEqual, uint64(35))
		So(D4, ShouldEqual, uint64(36))
		So(E4, ShouldEqual, uint64(37))
		So(F4, ShouldEqual, uint64(38))
		So(G4, ShouldEqual, uint64(39))
		So(H4, ShouldEqual, uint64(40))
		So(A3, ShouldEqual, uint64(41))
		So(B3, ShouldEqual, uint64(42))
		So(C3, ShouldEqual, uint64(43))
		So(D3, ShouldEqual, uint64(44))
		So(E3, ShouldEqual, uint64(45))
		So(F3, ShouldEqual, uint64(46))
		So(G3, ShouldEqual, uint64(47))
		So(H3, ShouldEqual, uint64(48))
		So(A2, ShouldEqual, uint64(49))
		So(B2, ShouldEqual, uint64(50))
		So(C2, ShouldEqual, uint64(51))
		So(D2, ShouldEqual, uint64(52))
		So(E2, ShouldEqual, uint64(53))
		So(F2, ShouldEqual, uint64(54))
		So(G2, ShouldEqual, uint64(55))
		So(H2, ShouldEqual, uint64(56))
		So(A1, ShouldEqual, uint64(57))
		So(B1, ShouldEqual, uint64(58))
		So(C1, ShouldEqual, uint64(59))
		So(D1, ShouldEqual, uint64(60))
		So(E1, ShouldEqual, uint64(61))
		So(F1, ShouldEqual, uint64(62))
		So(G1, ShouldEqual, uint64(63))
		So(H1, ShouldEqual, uint64(64))
	})

}

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
					bn := NewBitboard(1)
					So(bn.board, ShouldEqual, 1)
					bn = NewBitboard(3)
					So(bn.board, ShouldEqual, 8)
					bn = NewBitboard(1, 2)
					So(bn.board, ShouldEqual, 5)
					bn = NewBitboard(A8, H8)
					So(bn.board, ShouldEqual, 257)
				})
			})
		})
	})
	Convey("Given an existing BitBoard", t, func() {
		bb := NewBitboard()
		Convey("When FlipBit method called", func() {
			Convey("It should update its board attribute", func() {
				bb.FlipBit(4)
				So(bb.board, ShouldEqual, 16)
				bb.FlipBit(2)
				So(bb.board, ShouldEqual, 20)
				bb.FlipBit(2)
				So(bb.board, ShouldEqual, 16)
				bb.FlipBit(0)
				So(bb.board, ShouldEqual, 16)
			})
		})
	})
}

func TestPosition(t *testing.T) {

	Convey("Given a NewPosition", t, func() {
		Convey("returns a NewPosition struct", func() {
			pos := NewPosition()
			So(pos, ShouldHaveSameTypeAs, &Position{})
			So(pos, ShouldResemble, &Position{
				Pieces:          NewPiecePositions(),
				SideToMove:      WHITE,
				CastlingAbility: 0,
				EnPassantTarget: 0,
				HalfmoveClock:   0,
				FullMoveCounter: 1,
			})
		})
		Convey("SetPositionFromFen should accept FenString and set position accordingly", func() {
			for fen, position := range VALID_FENSTRINGS {
				pos := NewPosition()
				err := pos.SetPositionFromFen(fen)
				So(*pos, ShouldResemble, position)
				So(err, ShouldEqual, nil)
			}
			for fen, expectedErr := range INVALID_FENSTRINGS {
				pos := NewPosition()
				err := pos.SetPositionFromFen(fen)
				So(err, ShouldResemble, expectedErr)
			}

		})
	})
}
