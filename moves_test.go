package main

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

// nolint:funlen // Convey testing is verbose
func TestMoveStruct(t *testing.T) {
	Convey("Given a Move struct", t, func() {
		Convey("it should have relevant fields", func() {
			move := Move{}
			So(move.srcSquare, ShouldHaveSameTypeAs, board.Square(0))
			So(move.dstSquare, ShouldHaveSameTypeAs, board.Square(0))
			So(move.capture, ShouldHaveSameTypeAs, true)
			So(move.promotionTo, ShouldHaveSameTypeAs, "")
			So(move.piece, ShouldHaveSameTypeAs, "")
		})
		Convey("There should be a NewMove() func", func() {
			Convey("It should accept a LAN encoded move", func() {
				testCases := map[string]*Move{
					"e2e4": {
						piece:       "P",
						srcSquare:   board.E2,
						dstSquare:   board.E4,
						capture:     false,
						promotionTo: "",
					},
					"e2-e4": {
						piece:       "P",
						srcSquare:   board.E2,
						dstSquare:   board.E4,
						capture:     false,
						promotionTo: "",
					},
					"e2xe3": {
						piece:       "P",
						srcSquare:   board.E2,
						dstSquare:   board.E3,
						capture:     true,
						promotionTo: "",
					},
					"e7e8Q": {
						srcSquare:   board.E7,
						piece:       "P",
						dstSquare:   board.E8,
						capture:     false,
						promotionTo: "Q",
					},
					"E2E4": {
						piece:       "P",
						srcSquare:   board.E2,
						dstSquare:   board.E4,
						capture:     false,
						promotionTo: "",
					},
					"E2-E4": {
						piece:       "P",
						srcSquare:   board.E2,
						dstSquare:   board.E4,
						capture:     false,
						promotionTo: "",
					},
					"E2XE3": {
						srcSquare:   board.E2,
						piece:       "P",
						dstSquare:   board.E3,
						capture:     true,
						promotionTo: "",
					},
					"E7E8q": {
						piece:       "P",
						srcSquare:   board.E7,
						dstSquare:   board.E8,
						capture:     false,
						promotionTo: "Q",
					},
					"ke3-f5": {
						piece:       "K",
						srcSquare:   board.E3,
						dstSquare:   board.F5,
						capture:     false,
						promotionTo: "",
					},
					"qH3xe5": {
						piece:       "Q",
						srcSquare:   board.H3,
						dstSquare:   board.E5,
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
			Convey("It should return an error if the move is invalid", func() {
				testCases := map[string]MoveError{
					"e2e7": {
						fen:  "",
						move: "e2e7",
						err:  "Pawns do not move like that",
					},
				}
				for move, err := range testCases {
					m, mError := NewMove(move)
					So(*mError, ShouldResemble, err)
					So(m, ShouldEqual, nil)
				}
			})
		})
	})
}

// nolint:funlen // Convey testing is verbose
func TestValidMove(t *testing.T) {
	Convey("With a ValidMove() function", t, func() {
		Convey("It should accept a piece, source and destination arguments and return a bool.", func() {
			testCases := []struct {
				piece          string
				src            board.Square
				dst            board.Square
				expectedResult bool
			}{
				{
					"P",
					board.A2,
					board.A3,
					true,
				},
				{
					"P",
					board.E7,
					board.E5,
					true,
				},
				{
					"P",
					board.E7,
					board.E2,
					false,
				},
				{
					"P",
					board.E3,
					board.D4,
					true,
				},
				{
					"N",
					board.E3,
					board.F5,
					true,
				},
				{
					"N",
					board.E3,
					board.D4,
					false,
				},
				{
					"B",
					board.E3,
					board.D4,
					true,
				},
				{
					"B",
					board.E3,
					board.D3,
					false,
				},
				{
					"R",
					board.E3,
					board.D3,
					true,
				},
				{
					"R",
					board.E3,
					board.D4,
					false,
				},
				{
					"K",
					board.E3,
					board.D4,
					true,
				},
				{
					"K",
					board.E3,
					board.D5,
					false,
				},
				{
					"Q",
					board.E3,
					board.D4,
					true,
				},
				{
					"Q",
					board.E3,
					board.E4,
					true,
				},
				{
					"Q",
					board.E3,
					board.D1,
					false,
				},
			}
			for _, tc := range testCases {
				result := ValidMove(tc.piece, tc.src, tc.dst)
				So(result, ShouldEqual, tc.expectedResult)
			}
		})
	})
}

func TestKnightMoves(t *testing.T) {
	Convey("With a KnightMoves() func", t, func() {
		Convey("It should accept a square and return list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {
					board.C2, board.B3,
				},
				board.A8: {
					board.C7, board.B6,
				},
				board.H1: {
					board.G3, board.F2,
				},
				board.H8: {
					board.G6, board.F7,
				},
				board.E5: {
					board.D3, board.D7, board.F3, board.F7, board.G4, board.G6, board.C4, board.C6,
				},
			}
			for src, expectedDsts := range testCases {
				moves := KnightMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestOrthaganolRankMoves(t *testing.T) {
	Convey("With an OrthaganolRankMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {board.B1, board.C1, board.D1, board.E1, board.F1, board.G1, board.H1},
				board.H1: {board.A1, board.B1, board.C1, board.D1, board.E1, board.F1, board.G1},
				board.C1: {board.A1, board.B1, board.D1, board.E1, board.F1, board.G1, board.H1},
				board.A8: {board.B8, board.C8, board.D8, board.E8, board.F8, board.G8, board.H8},
				board.H8: {board.A8, board.B8, board.C8, board.D8, board.E8, board.F8, board.G8},
				board.C8: {board.A8, board.B8, board.D8, board.E8, board.F8, board.G8, board.H8},
				board.A6: {board.B6, board.C6, board.D6, board.E6, board.F6, board.G6, board.H6},
				board.H6: {board.A6, board.B6, board.C6, board.D6, board.E6, board.F6, board.G6},
				board.C6: {board.A6, board.B6, board.D6, board.E6, board.F6, board.G6, board.H6},
			}
			for src, expectedDsts := range testCases {
				moves := OrthaganolRankMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestOrthaganolFileMoves(t *testing.T) {
	Convey("With an OrthaganolFileMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {board.A2, board.A3, board.A4, board.A5, board.A6, board.A7, board.A8},
				board.A8: {board.A1, board.A2, board.A3, board.A4, board.A5, board.A6, board.A7},
				board.A5: {board.A1, board.A2, board.A3, board.A4, board.A6, board.A7, board.A8},
				board.H1: {board.H2, board.H3, board.H4, board.H5, board.H6, board.H7, board.H8},
				board.H8: {board.H1, board.H2, board.H3, board.H4, board.H5, board.H6, board.H7},
				board.H5: {board.H1, board.H2, board.H3, board.H4, board.H6, board.H7, board.H8},
				board.E1: {board.E2, board.E3, board.E4, board.E5, board.E6, board.E7, board.E8},
				board.E8: {board.E1, board.E2, board.E3, board.E4, board.E5, board.E6, board.E7},
				board.E5: {board.E1, board.E2, board.E3, board.E4, board.E6, board.E7, board.E8},
			}
			for src, expectedDsts := range testCases {
				moves := OrthaganolFileMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestDiagonalMoves(t *testing.T) {
	Convey("Given a DiagonalMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {board.B2, board.C3, board.D4, board.E5, board.F6, board.G7, board.H8},
				board.H8: {board.A1, board.B2, board.C3, board.D4, board.E5, board.F6, board.G7},
				board.A8: {board.B7, board.C6, board.D5, board.E4, board.F3, board.G2, board.H1},
				board.H1: {board.A8, board.B7, board.C6, board.D5, board.E4, board.F3, board.G2},
				board.E5: {
					board.B8, board.C7, board.D6, board.F4, board.G3, board.H2, board.A1,
					board.B2, board.C3, board.D4, board.F6, board.G7, board.H8,
				},
			}
			for src, expectedDsts := range testCases {
				moves := DiagonalMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestKingMoves(t *testing.T) {
	Convey("Given a KingMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {board.A2, board.B2, board.B1},
				board.H8: {board.H7, board.G7, board.G8},
				board.A8: {board.A7, board.B7, board.B8},
				board.H1: {board.H2, board.G2, board.G1},
				board.H2: {board.H1, board.H3, board.G1, board.G2, board.G3},
				board.A7: {board.A8, board.A6, board.B8, board.B7, board.B6},
				board.C1: {board.B1, board.B2, board.C2, board.D1, board.D2},
				board.E8: {board.D7, board.D8, board.E7, board.F7, board.F8},
				board.E5: {board.D4, board.D5, board.D6, board.E4, board.E6, board.F4, board.F5, board.F6},
			}
			for src, expectedDsts := range testCases {
				moves := KingMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestPawnMoves(t *testing.T) {
	Convey("Given a WhitePawnMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A2: {board.A3, board.A4, board.B3},
				board.H2: {board.H3, board.H4, board.G3},
				board.C2: {board.B3, board.C3, board.C4, board.D3},
				board.D3: {board.C4, board.D4, board.E4},
			}
			for src, expectedDsts := range testCases {
				moves := WhitePawnMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
	Convey("Given a BlackPawnMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A7: {board.A6, board.A5, board.B6},
				board.H7: {board.H6, board.H5, board.G6},
				board.D3: {board.C2, board.D2, board.E2},
				board.C7: {board.B6, board.C6, board.C5, board.D6},
			}
			for src, expectedDsts := range testCases {
				moves := BlackPawnMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}
