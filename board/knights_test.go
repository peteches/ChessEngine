package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKnights(t *testing.T) {
	Convey("Given a knight struct", t, func() {
		knight := board.Knights{}
		Convey("It should have a Positions field", func() {
			So(knight.Positions, ShouldHaveSameTypeAs, &board.BitBoard{})
			var testSide board.Side
			So(knight.Colour, ShouldHaveSameTypeAs, testSide)
		})
		Convey("It should implement the Piece interface", func() {
			bKnight := board.NewKnights(board.Black)
			wKnight := board.NewKnights(board.White)
			Convey("By returning the appropriate string representation of the piece", func() {
				So(bKnight.String(), ShouldEqual, "n")
				So(wKnight.String(), ShouldEqual, "N")
			})
			Convey("By returning validating valid moves", func() {
				So(bKnight.ValidMove(board.A1, board.A2), ShouldBeFalse)
				So(bKnight.ValidMove(board.A1, board.B1), ShouldBeFalse)
				So(bKnight.ValidMove(board.A1, board.B3), ShouldBeTrue)
				So(bKnight.ValidMove(board.B1, board.H2), ShouldBeFalse)
			})
			So(knight, ShouldImplement, (*board.Piece)(nil))
		})
	})
	Convey("Given a NewKnight() function", t, func() {
		Convey("It should accept a side", func() {
			knight := board.NewKnights(board.Black)
			So(*knight, ShouldResemble, board.Knights{board.NewBitboard(), board.Black})
		})
		Convey("It should Optionally accept any number of Squares to initialise the position", func() {
			knight := board.NewKnights(board.Black, board.D8, board.A1)
			So(*knight, ShouldResemble, board.Knights{board.NewBitboard(board.D8, board.A1), board.Black})
		})
	})
}
