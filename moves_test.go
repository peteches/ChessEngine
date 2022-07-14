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
