package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

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
				moves := board.KnightMoves(src)
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
				moves := board.OrthaganolRankMoves(src)
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
				moves := board.OrthaganolFileMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestDiagonalMoves(t *testing.T) {
	Convey("Given a DiagonalMovesA1H8() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {board.B2, board.C3, board.D4, board.E5, board.F6, board.G7, board.H8},
				board.H8: {board.A1, board.B2, board.C3, board.D4, board.E5, board.F6, board.G7},
				board.E2: {board.D1, board.F3, board.G4, board.H5},
				board.E5: {
					board.A1, board.B2, board.C3, board.D4, board.F6, board.G7, board.H8,
				},
				board.B1: {board.C2, board.D3, board.E4, board.F5, board.G6, board.H7},
				board.G1: {board.H2},
				board.B8: {board.A7},
				board.G8: {board.A2, board.B3, board.C4, board.D5, board.E6, board.F7},
			}
			for src, expectedDsts := range testCases {
				moves := board.DiagonalMovesA1H8(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
	Convey("Given a DiagonalMovesA8H1() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A8: {board.B7, board.C6, board.D5, board.E4, board.F3, board.G2, board.H1},
				board.H1: {board.A8, board.B7, board.C6, board.D5, board.E4, board.F3, board.G2},
				board.E2: {board.F1, board.D3, board.C4, board.B5, board.A6},
				board.E5: {board.B8, board.C7, board.D6, board.F4, board.G3, board.H2},
				board.B1: {board.A2},
				board.G1: {board.A7, board.B6, board.C5, board.D4, board.E3, board.F2},
				board.B8: {board.C7, board.D6, board.E5, board.F4, board.G3, board.H2},
				board.G8: {board.H7},
			}
			for src, expectedDsts := range testCases {
				moves := board.DiagonalMovesA8H1(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
	Convey("Given a DiagonalMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A1: {board.B2, board.C3, board.D4, board.E5, board.F6, board.G7, board.H8},
				board.H8: {board.A1, board.B2, board.C3, board.D4, board.E5, board.F6, board.G7},
				board.A8: {board.B7, board.C6, board.D5, board.E4, board.F3, board.G2, board.H1},
				board.H1: {board.A8, board.B7, board.C6, board.D5, board.E4, board.F3, board.G2},
				board.E2: {board.F1, board.D1, board.D3, board.C4, board.B5, board.A6, board.F3, board.G4, board.H5},
				board.E5: {
					board.B8, board.C7, board.D6, board.F4, board.G3, board.H2, board.A1,
					board.B2, board.C3, board.D4, board.F6, board.G7, board.H8,
				},
				board.B1: {board.A2, board.C2, board.D3, board.E4, board.F5, board.G6, board.H7},
				board.G1: {board.A7, board.B6, board.C5, board.D4, board.E3, board.F2, board.H2},
				board.B8: {board.A7, board.C7, board.D6, board.E5, board.F4, board.G3, board.H2},
				board.G8: {board.A2, board.B3, board.C4, board.D5, board.E6, board.F7, board.H7},
			}
			for src, expectedDsts := range testCases {
				moves := board.DiagonalMoves(src)
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
				moves := board.KingMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}

func TestPawnMoves(t *testing.T) {
	Convey("Given a WhitePawnCaptureMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A2: {board.B3},
				board.H2: {board.G3},
				board.C2: {board.B3, board.D3},
				board.D3: {board.C4, board.E4},
			}
			for src, expectedDsts := range testCases {
				moves := board.WhitePawnCaptureMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
	Convey("Given a WhitePawnAdvanceMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A2: {board.A3, board.A4},
				board.H2: {board.H3, board.H4},
				board.C2: {board.C3, board.C4},
				board.D3: {board.D4},
			}
			for src, expectedDsts := range testCases {
				moves := board.WhitePawnAdvanceMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
	Convey("Given a BlackPawnCaptureMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A7: {board.B6},
				board.H7: {board.G6},
				board.D3: {board.C2, board.E2},
				board.C7: {board.B6, board.D6},
			}
			for src, expectedDsts := range testCases {
				moves := board.BlackPawnCaptureMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
	Convey("Given a BlackPawnAdvanceMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[board.Square][]board.Square{
				board.A7: {board.A6, board.A5},
				board.H7: {board.H6, board.H5},
				board.D3: {board.D2},
				board.C7: {board.C6, board.C5},
			}
			for src, expectedDsts := range testCases {
				moves := board.BlackPawnAdvanceMoves(src)
				So(moves, ShouldHaveLength, len(expectedDsts))
				for _, dstSqr := range expectedDsts {
					So(moves, ShouldContain, dstSqr)
				}
			}
		})
	})
}
