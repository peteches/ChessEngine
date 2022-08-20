package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKings(t *testing.T) {
	Convey("Given a king struct", t, func() {
		king := board.King{}
		Convey("It should have a Colour field", func() {
			var testSide board.Side
			So(king.Colour, ShouldHaveSameTypeAs, testSide)
		})
		Convey("It should implement the Piece interface", func() {
			bKing := board.NewKing(board.Black)
			wKing := board.NewKing(board.White)
			Convey("By returning the appropriate string representation of the piece", func() {
				So(bKing.String(), ShouldEqual, "k")
				So(wKing.String(), ShouldEqual, "K")
			})
			Convey("By returning validating valid moves", func() {
				So(bKing.ValidMove(board.A1, board.A2), ShouldBeTrue)
				So(bKing.ValidMove(board.A1, board.A3), ShouldBeFalse)
				So(bKing.ValidMove(board.A1, board.B1), ShouldBeTrue)
				So(bKing.ValidMove(board.A1, board.B2), ShouldBeTrue)
				So(bKing.ValidMove(board.B1, board.H2), ShouldBeFalse)
			})
			Convey("By returning a pointer to it's internal BitBoard", func() {
				So(king.Positions(), ShouldHaveSameTypeAs, &board.BitBoard{})
			})
			So(king, ShouldImplement, (*board.Piece)(nil))
		})
	})
	Convey("Given a NewKing() function", t, func() {
		Convey("It should accept a side", func() {
			king := board.NewKing(board.Black)
			So(*king, ShouldResemble, board.King{board.NewBitboard(), board.Black})
		})
		Convey("It should Optionally accept any number of Squares to initialise the position", func() {
			king := board.NewKing(board.Black, board.D8, board.A1)
			So(*king, ShouldResemble, board.King{board.NewBitboard(board.D8, board.A1), board.Black})
		})
	})
}
