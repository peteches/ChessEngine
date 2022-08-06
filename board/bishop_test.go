package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBishops(t *testing.T) {
	Convey("Given a bishop struct", t, func() {
		bishop := board.Bishops{}
		Convey("It should have a Positions field", func() {
			So(bishop.Positions, ShouldHaveSameTypeAs, &board.BitBoard{})
			var testSide board.Side
			So(bishop.Colour, ShouldHaveSameTypeAs, testSide)
		})
		Convey("It should implement the Piece interface", func() {
			bBishop := board.NewBishop(board.Black)
			wBishop := board.NewBishop(board.White)
			Convey("By returning the appropriate string representation of the piece", func() {
				So(bBishop.String(), ShouldEqual, "b")
				So(wBishop.String(), ShouldEqual, "B")
			})
			Convey("By returning validating valid moves", func() {
				So(bBishop.ValidMove(board.A1, board.A2), ShouldBeFalse)
				So(bBishop.ValidMove(board.A1, board.B1), ShouldBeFalse)
				So(bBishop.ValidMove(board.A1, board.B2), ShouldBeTrue)
				So(bBishop.ValidMove(board.B1, board.H2), ShouldBeFalse)
			})
			So(bishop, ShouldImplement, (*board.Piece)(nil))
		})
	})
	Convey("Given a NewBishop() function", t, func() {
		Convey("It should accept a side", func() {
			bishop := board.NewBishop(board.Black)
			So(*bishop, ShouldResemble, board.Bishops{board.NewBitboard(), board.Black})
		})
		Convey("It should Optionally accept any number of Squares to initialise the position", func() {
			bishop := board.NewBishop(board.Black, board.D8, board.A1)
			So(*bishop, ShouldResemble, board.Bishops{board.NewBitboard(board.D8, board.A1), board.Black})
		})
	})
}
