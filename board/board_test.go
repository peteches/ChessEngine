package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:funlen // convey testing is verbose
func TestConstants(t *testing.T) {
	Convey("Constants Set Correctly", t, func() {
		Convey("Board Squares", func() {
			So(board.A1, ShouldEqual, board.Square(1))
			So(board.B1, ShouldEqual, board.Square(1<<1))
			So(board.C1, ShouldEqual, board.Square(1<<2))
			So(board.D1, ShouldEqual, board.Square(1<<3))
			So(board.E1, ShouldEqual, board.Square(1<<4))
			So(board.F1, ShouldEqual, board.Square(1<<5))
			So(board.G1, ShouldEqual, board.Square(1<<6))
			So(board.H1, ShouldEqual, board.Square(1<<7))
			So(board.A2, ShouldEqual, board.Square(1<<8))
			So(board.B2, ShouldEqual, board.Square(1<<9))
			So(board.C2, ShouldEqual, board.Square(1<<10))
			So(board.D2, ShouldEqual, board.Square(1<<11))
			So(board.E2, ShouldEqual, board.Square(1<<12))
			So(board.F2, ShouldEqual, board.Square(1<<13))
			So(board.G2, ShouldEqual, board.Square(1<<14))
			So(board.H2, ShouldEqual, board.Square(1<<15))
			So(board.A3, ShouldEqual, board.Square(1<<16))
			So(board.B3, ShouldEqual, board.Square(1<<17))
			So(board.C3, ShouldEqual, board.Square(1<<18))
			So(board.D3, ShouldEqual, board.Square(1<<19))
			So(board.E3, ShouldEqual, board.Square(1<<20))
			So(board.F3, ShouldEqual, board.Square(1<<21))
			So(board.G3, ShouldEqual, board.Square(1<<22))
			So(board.H3, ShouldEqual, board.Square(1<<23))
			So(board.A4, ShouldEqual, board.Square(1<<24))
			So(board.B4, ShouldEqual, board.Square(1<<25))
			So(board.C4, ShouldEqual, board.Square(1<<26))
			So(board.D4, ShouldEqual, board.Square(1<<27))
			So(board.E4, ShouldEqual, board.Square(1<<28))
			So(board.F4, ShouldEqual, board.Square(1<<29))
			So(board.G4, ShouldEqual, board.Square(1<<30))
			So(board.H4, ShouldEqual, board.Square(1<<31))
			So(board.A5, ShouldEqual, board.Square(1<<32))
			So(board.B5, ShouldEqual, board.Square(1<<33))
			So(board.C5, ShouldEqual, board.Square(1<<34))
			So(board.D5, ShouldEqual, board.Square(1<<35))
			So(board.E5, ShouldEqual, board.Square(1<<36))
			So(board.F5, ShouldEqual, board.Square(1<<37))
			So(board.G5, ShouldEqual, board.Square(1<<38))
			So(board.H5, ShouldEqual, board.Square(1<<39))
			So(board.A6, ShouldEqual, board.Square(1<<40))
			So(board.B6, ShouldEqual, board.Square(1<<41))
			So(board.C6, ShouldEqual, board.Square(1<<42))
			So(board.D6, ShouldEqual, board.Square(1<<43))
			So(board.E6, ShouldEqual, board.Square(1<<44))
			So(board.F6, ShouldEqual, board.Square(1<<45))
			So(board.G6, ShouldEqual, board.Square(1<<46))
			So(board.H6, ShouldEqual, board.Square(1<<47))
			So(board.A7, ShouldEqual, board.Square(1<<48))
			So(board.B7, ShouldEqual, board.Square(1<<49))
			So(board.C7, ShouldEqual, board.Square(1<<50))
			So(board.D7, ShouldEqual, board.Square(1<<51))
			So(board.E7, ShouldEqual, board.Square(1<<52))
			So(board.F7, ShouldEqual, board.Square(1<<53))
			So(board.G7, ShouldEqual, board.Square(1<<54))
			So(board.H7, ShouldEqual, board.Square(1<<55))
			So(board.A8, ShouldEqual, board.Square(1<<56))
			So(board.B8, ShouldEqual, board.Square(1<<57))
			So(board.C8, ShouldEqual, board.Square(1<<58))
			So(board.D8, ShouldEqual, board.Square(1<<59))
			So(board.E8, ShouldEqual, board.Square(1<<60))
			So(board.F8, ShouldEqual, board.Square(1<<61))
			So(board.G8, ShouldEqual, board.Square(1<<62))
			So(board.H8, ShouldEqual, board.Square(1<<63))
		})
		Convey("Side Types should be set", func() {
			var testSide board.Side
			So(board.White, ShouldHaveSameTypeAs, testSide)
			So(board.Black, ShouldHaveSameTypeAs, testSide)

			So(board.White, ShouldEqual, 0)
			So(board.Black, ShouldEqual, 1)
		})
	})
}

func TestBoardMatricies(t *testing.T) {
	Convey("Given a BoardMatrixStoI func", t, func() {
		Convey("It should map string co-ordinates to Const Squares", func() {
			testCases := map[string]board.Square{
				"A8": board.A8,
			}
			for sqr, expectedSquare := range testCases {
				resultSquare := board.BoardMatrixStoI[sqr]
				So(resultSquare, ShouldEqual, expectedSquare)
			}
		})
	})
	Convey("Given a BoardMatrixItoS func", t, func() {
		Convey("It should map Const Squares to string co-ordinates", func() {
			testCases := map[board.Square]string{
				board.A8: "A8",
			}
			for sqr, expectedSquare := range testCases {
				resultSquare := board.BoardMatrixItoS[sqr]
				So(resultSquare, ShouldEqual, expectedSquare)
			}
		})
	})
}

//nolint:funlen // convey testing is verbose
func TestBitboard(t *testing.T) {
	Convey("Given a board.NewBitboard function", t, func() {
		Convey("With no arguments", func() {
			Convey("It should return an empty board", func() {
				bb := board.NewBitboard()
				So(*bb, ShouldResemble, board.BitBoard{})
				So(bb.Board, ShouldEqual, 0)
			})

			Convey("With args", func() {
				Convey("It should return an initialised board", func() {
					bitBoard := board.NewBitboard(board.A8)
					So(bitBoard.Board, ShouldEqual, board.A8)
					bitBoard = board.NewBitboard(board.E3)
					So(bitBoard.Board, ShouldEqual, board.E3)
					bitBoard = board.NewBitboard(board.A2, board.B2)
					So(bitBoard.Board, ShouldEqual, board.A2+board.B2)
					bitBoard = board.NewBitboard(board.A8, board.H1)
					So(bitBoard.Board, ShouldEqual, board.A8+board.H1)
				})
			})
		})
	})
	Convey("Given an existing BitBoard", t, func() {
		bitboard := board.NewBitboard()
		Convey("Bit manipulation basics", func() {
			So(0^(1<<0), ShouldEqual, 1)
		})
		Convey("When FlipBit method called", func() {
			Convey("It should update its board attribute", func() {
				So(bitboard.Board, ShouldEqual, 0)
				for _, sqr := range board.AllSquares {
					bb := board.NewBitboard()
					So(bb.Board, ShouldEqual, 0)
					bb.FlipBit(sqr)
					So(bb.Board, ShouldEqual, sqr)
				}
			})
		})
		Convey("When board.Squares() method called returns []Square where bits are 1", func() {
			So(bitboard.Squares(), ShouldResemble, []board.Square{})
			bitboard.FlipBit(board.A8)
			So(bitboard.Squares(), ShouldHaveLength, 1)
			So(bitboard.Squares(), ShouldContain, board.A8)
			bitboard.FlipBit(board.H3)
			bitboard.FlipBit(board.H4)
			So(bitboard.Squares(), ShouldHaveLength, 3)
			So(bitboard.Squares(), ShouldContain, board.A8)
			So(bitboard.Squares(), ShouldContain, board.H4)
			So(bitboard.Squares(), ShouldContain, board.H3)
		})

		Convey("When Occupied() method called with square, returns true if square occupied", func() {
			for _, sqr := range board.AllSquares {
				So(bitboard.Occupied(sqr), ShouldEqual, false)
			}
			bitboard.FlipBit(board.B4)
			So(bitboard.Occupied(board.B4), ShouldEqual, true)
			bitboard.FlipBit(board.A8)
			So(bitboard.Occupied(board.B4), ShouldEqual, true)
			So(bitboard.Occupied(board.A8), ShouldEqual, true)
			So(bitboard.Occupied(board.H7), ShouldEqual, false)
			bitboard.FlipBit(board.H7)
			So(bitboard.Occupied(board.H7), ShouldEqual, true)
		})
	})
}
