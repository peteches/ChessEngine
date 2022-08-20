package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:funlen,maintidx // Convey testing is verbose.
func TestSquare(t *testing.T) {
	Convey("Constants Set Correctly", t, func() {
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
	Convey("Given a board.Square type", t, func() {
		Convey("It should have a String() method that reveal the co-ordinates of the square", func() {
			So(board.A1.String(), ShouldEqual, "A1")
			So(board.B1.String(), ShouldEqual, "B1")
			So(board.C1.String(), ShouldEqual, "C1")
			So(board.D1.String(), ShouldEqual, "D1")
			So(board.E1.String(), ShouldEqual, "E1")
			So(board.F1.String(), ShouldEqual, "F1")
			So(board.G1.String(), ShouldEqual, "G1")
			So(board.H1.String(), ShouldEqual, "H1")
			So(board.A2.String(), ShouldEqual, "A2")
			So(board.B2.String(), ShouldEqual, "B2")
			So(board.C2.String(), ShouldEqual, "C2")
			So(board.D2.String(), ShouldEqual, "D2")
			So(board.E2.String(), ShouldEqual, "E2")
			So(board.F2.String(), ShouldEqual, "F2")
			So(board.G2.String(), ShouldEqual, "G2")
			So(board.H2.String(), ShouldEqual, "H2")
			So(board.A3.String(), ShouldEqual, "A3")
			So(board.B3.String(), ShouldEqual, "B3")
			So(board.C3.String(), ShouldEqual, "C3")
			So(board.D3.String(), ShouldEqual, "D3")
			So(board.E3.String(), ShouldEqual, "E3")
			So(board.F3.String(), ShouldEqual, "F3")
			So(board.G3.String(), ShouldEqual, "G3")
			So(board.H3.String(), ShouldEqual, "H3")
			So(board.A4.String(), ShouldEqual, "A4")
			So(board.B4.String(), ShouldEqual, "B4")
			So(board.C4.String(), ShouldEqual, "C4")
			So(board.D4.String(), ShouldEqual, "D4")
			So(board.E4.String(), ShouldEqual, "E4")
			So(board.F4.String(), ShouldEqual, "F4")
			So(board.G4.String(), ShouldEqual, "G4")
			So(board.H4.String(), ShouldEqual, "H4")
			So(board.A5.String(), ShouldEqual, "A5")
			So(board.B5.String(), ShouldEqual, "B5")
			So(board.C5.String(), ShouldEqual, "C5")
			So(board.D5.String(), ShouldEqual, "D5")
			So(board.E5.String(), ShouldEqual, "E5")
			So(board.F5.String(), ShouldEqual, "F5")
			So(board.G5.String(), ShouldEqual, "G5")
			So(board.H5.String(), ShouldEqual, "H5")
			So(board.A6.String(), ShouldEqual, "A6")
			So(board.B6.String(), ShouldEqual, "B6")
			So(board.C6.String(), ShouldEqual, "C6")
			So(board.D6.String(), ShouldEqual, "D6")
			So(board.E6.String(), ShouldEqual, "E6")
			So(board.F6.String(), ShouldEqual, "F6")
			So(board.G6.String(), ShouldEqual, "G6")
			So(board.H6.String(), ShouldEqual, "H6")
			So(board.A7.String(), ShouldEqual, "A7")
			So(board.B7.String(), ShouldEqual, "B7")
			So(board.C7.String(), ShouldEqual, "C7")
			So(board.D7.String(), ShouldEqual, "D7")
			So(board.E7.String(), ShouldEqual, "E7")
			So(board.F7.String(), ShouldEqual, "F7")
			So(board.G7.String(), ShouldEqual, "G7")
			So(board.H7.String(), ShouldEqual, "H7")
			So(board.A8.String(), ShouldEqual, "A8")
			So(board.B8.String(), ShouldEqual, "B8")
			So(board.C8.String(), ShouldEqual, "C8")
			So(board.D8.String(), ShouldEqual, "D8")
			So(board.E8.String(), ShouldEqual, "E8")
			So(board.F8.String(), ShouldEqual, "F8")
			So(board.G8.String(), ShouldEqual, "G8")
			So(board.H8.String(), ShouldEqual, "H8")
		})
		Convey("It should have a File() method that reveals which file the square is in", func() {
			testCases := map[board.Square]uint8{
				board.A1: 1, board.A2: 1, board.A3: 1, board.A4: 1, board.A5: 1, board.A6: 1, board.A7: 1, board.A8: 1,
				board.B1: 2, board.B2: 2, board.B3: 2, board.B4: 2, board.B5: 2, board.B6: 2, board.B7: 2, board.B8: 2,
				board.C1: 3, board.C2: 3, board.C3: 3, board.C4: 3, board.C5: 3, board.C6: 3, board.C7: 3, board.C8: 3,
				board.D1: 4, board.D2: 4, board.D3: 4, board.D4: 4, board.D5: 4, board.D6: 4, board.D7: 4, board.D8: 4,
				board.E1: 5, board.E2: 5, board.E3: 5, board.E4: 5, board.E5: 5, board.E6: 5, board.E7: 5, board.E8: 5,
				board.F1: 6, board.F2: 6, board.F3: 6, board.F4: 6, board.F5: 6, board.F6: 6, board.F7: 6, board.F8: 6,
				board.G1: 7, board.G2: 7, board.G3: 7, board.G4: 7, board.G5: 7, board.G6: 7, board.G7: 7, board.G8: 7,
				board.H1: 8, board.H2: 8, board.H3: 8, board.H4: 8, board.H5: 8, board.H6: 8, board.H7: 8, board.H8: 8,
			}
			for sqr, expectedFile := range testCases {
				So(sqr.File(), ShouldEqual, expectedFile)
				So(sqr.File(), ShouldBeLessThan, expectedFile+1)
				So(sqr.File(), ShouldBeGreaterThan, expectedFile-1)
			}
		})
		Convey("It should have a Rank() method that reveals which rank the square is in", func() {
			testCases := map[board.Square]uint8{
				board.A1: 1, board.A2: 2, board.A3: 3, board.A4: 4, board.A5: 5, board.A6: 6, board.A7: 7, board.A8: 8,
				board.B1: 1, board.B2: 2, board.B3: 3, board.B4: 4, board.B5: 5, board.B6: 6, board.B7: 7, board.B8: 8,
				board.C1: 1, board.C2: 2, board.C3: 3, board.C4: 4, board.C5: 5, board.C6: 6, board.C7: 7, board.C8: 8,
				board.D1: 1, board.D2: 2, board.D3: 3, board.D4: 4, board.D5: 5, board.D6: 6, board.D7: 7, board.D8: 8,
				board.E1: 1, board.E2: 2, board.E3: 3, board.E4: 4, board.E5: 5, board.E6: 6, board.E7: 7, board.E8: 8,
				board.F1: 1, board.F2: 2, board.F3: 3, board.F4: 4, board.F5: 5, board.F6: 6, board.F7: 7, board.F8: 8,
				board.G1: 1, board.G2: 2, board.G3: 3, board.G4: 4, board.G5: 5, board.G6: 6, board.G7: 7, board.G8: 8,
				board.H1: 1, board.H2: 2, board.H3: 3, board.H4: 4, board.H5: 5, board.H6: 6, board.H7: 7, board.H8: 8,
			}
			for sqr, expectedRank := range testCases {
				So(sqr.Rank(), ShouldEqual, expectedRank)
				So(sqr.Rank(), ShouldBeLessThan, expectedRank+1)
				So(sqr.Rank(), ShouldBeGreaterThan, expectedRank-1)
			}
		})
		Convey("It should have a OnEdge() method that returns true if the square is on the edge of the board", func() {
			testCases := map[board.Square]bool{
				board.A1: true,
				board.B1: true,
				board.C1: true,
				board.D1: true,
				board.E1: true,
				board.F1: true,
				board.G1: true,
				board.H1: true,
				board.A2: true,
				board.B2: false,
				board.C2: false,
				board.D2: false,
				board.E2: false,
				board.F2: false,
				board.G2: false,
				board.H2: true,
				board.A3: true,
				board.B3: false,
				board.C3: false,
				board.D3: false,
				board.E3: false,
				board.F3: false,
				board.G3: false,
				board.H3: true,
				board.A4: true,
				board.B4: false,
				board.C4: false,
				board.D4: false,
				board.E4: false,
				board.F4: false,
				board.G4: false,
				board.H4: true,
				board.A5: true,
				board.B5: false,
				board.C5: false,
				board.D5: false,
				board.E5: false,
				board.F5: false,
				board.G5: false,
				board.H5: true,
				board.A6: true,
				board.B6: false,
				board.C6: false,
				board.D6: false,
				board.E6: false,
				board.F6: false,
				board.G6: false,
				board.H6: true,
				board.A7: true,
				board.B7: false,
				board.C7: false,
				board.D7: false,
				board.E7: false,
				board.F7: false,
				board.G7: false,
				board.H7: true,
				board.A8: true,
				board.B8: true,
				board.C8: true,
				board.D8: true,
				board.E8: true,
				board.F8: true,
				board.G8: true,
				board.H8: true,
			}
			for sqr, expectedResult := range testCases {
				So(sqr.OnEdge(), ShouldEqual, expectedResult)
			}
		})
	})
}

func TestSquaresAdjacent(t *testing.T) {
	Convey("Given a SquaresAdjacent() Function", t, func() {
		Convey("It should return true if src and dst are 1 square away", func() {
			testCasesAdj := [][2]board.Square{
				{board.E5, board.E6},
				{board.E5, board.E4},
				{board.E5, board.D4},
				{board.E5, board.D5},
				{board.E5, board.D6},
				{board.E5, board.F4},
				{board.E5, board.F5},
				{board.E5, board.F6},
			}
			for _, tc := range testCasesAdj {
				So(board.SquaresAdjacent(tc[0], tc[1]), ShouldBeTrue)
			}
		})
		Convey("It should return false if src and dst are more than 1 square away", func() {
			testCasesAdj := [][2]board.Square{
				{board.A1, board.A3},
			}
			for _, tc := range testCasesAdj {
				So(board.SquaresAdjacent(tc[0], tc[1]), ShouldBeFalse)
			}
		})
	})
}

func TestSquaresBetween(t *testing.T) {
	Convey("Given a SquaresBetween() Function", t, func() {
		Convey("It should return an empty list if the squares are not on the same diagonal, file or rank", func() {
			testCases := [][2]board.Square{
				{board.A1, board.B1},
			}
			for _, x := range testCases {
				So(board.SquaresBetween(x[0], x[1]), ShouldBeEmpty)
			}
		})
		Convey("It should return a list of Squares between src and dst if they are on the same diagonal, rank or file", func() {
			Convey("This should be an empty list if the squares are adjacant or do not share a diagonal, rank or file", func() {
				testCases := [][2]board.Square{
					{board.A1, board.B2},
					{board.A1, board.A2},
					{board.A1, board.B1},
					{board.E8, board.A1},
				}
				for _, x := range testCases {
					So(board.SquaresBetween(x[0], x[1]), ShouldBeEmpty)
				}
			})
			Convey("This should not be an empty list if the squares are not adjacant", func() {
				testCases := map[[2]board.Square][]board.Square{
					{board.A1, board.A8}: {board.A2, board.A3, board.A4, board.A5, board.A6, board.A7},
					{board.A8, board.A1}: {board.A2, board.A3, board.A4, board.A5, board.A6, board.A7},
					{board.A1, board.H1}: {board.B1, board.C1, board.D1, board.E1, board.F1, board.G1},
					{board.H1, board.A1}: {board.B1, board.C1, board.D1, board.E1, board.F1, board.G1},
					{board.B2, board.H8}: {board.C3, board.D4, board.E5, board.F6, board.G7},
					{board.E2, board.C4}: {board.D3},
				}
				for x, expectedSquares := range testCases {
					result := board.SquaresBetween(x[0], x[1])
					So(result, ShouldHaveLength, len(expectedSquares))
					for _, sqr := range expectedSquares {
						So(result, ShouldContain, sqr)
					}
				}
			})
		})
	})
}
