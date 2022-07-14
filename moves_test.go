package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
