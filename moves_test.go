package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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

func TestKnightMoves(t *testing.T) {
	Convey("With a KnightMoves() func", t, func() {
		Convey("It should accept a square and return list of valid destination moves", func() {
			testCases := map[Square][]Square{
				A1: {
					C2, B3,
				},
				A8: {
					C7, B6,
				},
				H1: {
					G3, F2,
				},
				H8: {
					G6, F7,
				},
				E5: {
					D3, D7, F3, F7, G4, G6, C4, C6,
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

// nolint:dupl // this could be merged with Rank Moves but prefer clarity of
// separation.
func TestOrthaganolRankMoves(t *testing.T) {
	Convey("With an OrthaganolRankMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[Square][]Square{
				A1: {B1, C1, D1, E1, F1, G1, H1},
				H1: {A1, B1, C1, D1, E1, F1, G1},
				C1: {A1, B1, D1, E1, F1, G1, H1},
				A8: {B8, C8, D8, E8, F8, G8, H8},
				H8: {A8, B8, C8, D8, E8, F8, G8},
				C8: {A8, B8, D8, E8, F8, G8, H8},
				A6: {B6, C6, D6, E6, F6, G6, H6},
				H6: {A6, B6, C6, D6, E6, F6, G6},
				C6: {A6, B6, D6, E6, F6, G6, H6},
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

// nolint:dupl // this could be merged with Rank Moves but prefer clarity of
// separation.
func TestOrthaganolFileMoves(t *testing.T) {
	Convey("With an OrthaganolFileMoves() func", t, func() {
		Convey("It should accept a square and return a list of valid destination moves", func() {
			testCases := map[Square][]Square{
				A1: {A2, A3, A4, A5, A6, A7, A8},
				A8: {A1, A2, A3, A4, A5, A6, A7},
				A5: {A1, A2, A3, A4, A6, A7, A8},
				H1: {H2, H3, H4, H5, H6, H7, H8},
				H8: {H1, H2, H3, H4, H5, H6, H7},
				H5: {H1, H2, H3, H4, H6, H7, H8},
				E1: {E2, E3, E4, E5, E6, E7, E8},
				E8: {E1, E2, E3, E4, E5, E6, E7},
				E5: {E1, E2, E3, E4, E6, E7, E8},
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
			testCases := map[Square][]Square{
				A1: {B2, C3, D4, E5, F6, G7, H8},
				H8: {A1, B2, C3, D4, E5, F6, G7},
				A8: {B7, C6, D5, E4, F3, G2, H1},
				H1: {A8, B7, C6, D5, E4, F3, G2},
				E5: {B8, C7, D6, F4, G3, H2, A1, B2, C3, D4, F6, G7, H8},
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
			testCases := map[Square][]Square{
				A1: {A2, B2, B1},
				H8: {H7, G7, G8},
				A8: {A7, B7, B8},
				H1: {H2, G2, G1},
				H2: {H1, H3, G1, G2, G3},
				A7: {A8, A6, B8, B7, B6},
				C1: {B1, B2, C2, D1, D2},
				E8: {D7, D8, E7, F7, F8},
				E5: {D4, D5, D6, E4, E6, F4, F5, F6},
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
			testCases := map[Square][]Square{
				A2: {A3, A4, B3},
				H2: {H3, H4, G3},
				C2: {B3, C3, C4, D3},
				D3: {C4, D4, E4},
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
			testCases := map[Square][]Square{
				A7: {A6, A5, B6},
				H7: {H6, H5, G6},
				D3: {C2, D2, E2},
				C7: {B6, C6, C5, D6},
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
