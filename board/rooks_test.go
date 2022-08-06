package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRooks(t *testing.T) {
	Convey("Given a rook struct", t, func() {
		rook := board.Rooks{}
		Convey("It should have a Positions field", func() {
			So(rook.Positions, ShouldHaveSameTypeAs, &board.BitBoard{})
			var testSide board.Side
			So(rook.Colour, ShouldHaveSameTypeAs, testSide)
		})
		Convey("It should implement the Piece interface", func() {
			bRook := board.NewRooks(board.Black)
			wRook := board.NewRooks(board.White)
			Convey("By returning the appropriate string representation of the piece", func() {
				So(bRook.String(), ShouldEqual, "r")
				So(wRook.String(), ShouldEqual, "R")
			})
			Convey("By returning validating valid moves", func() {
				So(bRook.ValidMove(board.A1, board.A2), ShouldBeTrue)
				So(bRook.ValidMove(board.A1, board.B1), ShouldBeTrue)
				So(bRook.ValidMove(board.A1, board.B2), ShouldBeFalse)
				So(bRook.ValidMove(board.B1, board.H2), ShouldBeFalse)
			})
			So(rook, ShouldImplement, (*board.Piece)(nil))
		})
	})
	Convey("Given a NewRook() function", t, func() {
		Convey("It should accept a side", func() {
			rook := board.NewRooks(board.Black)
			So(*rook, ShouldResemble, board.Rooks{board.NewBitboard(), board.Black})
		})
		Convey("It should Optionally accept any number of Squares to initialise the position", func() {
			rook := board.NewRooks(board.Black, board.D8, board.A1)
			So(*rook, ShouldResemble, board.Rooks{board.NewBitboard(board.D8, board.A1), board.Black})
		})
	})
}
