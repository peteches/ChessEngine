package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPawnss(t *testing.T) {
	Convey("Given a pawns struct", t, func() {
		pawns := board.Pawns{}
		Convey("It should have a Positions field", func() {
			So(pawns.Positions, ShouldHaveSameTypeAs, &board.BitBoard{})
			var testSide board.Side
			So(pawns.Colour, ShouldHaveSameTypeAs, testSide)
		})
		Convey("It should implement the Piece interface", func() {
			bPawns := board.NewPawns(board.Black)
			wPawns := board.NewPawns(board.White)
			Convey("By returning the appropriate string representation of the piece", func() {
				So(bPawns.String(), ShouldEqual, "p")
				So(wPawns.String(), ShouldEqual, "P")
			})
			Convey("By returning validating valid moves", func() {
				So(wPawns.ValidMove(board.A1, board.B1), ShouldBeFalse)
				So(wPawns.ValidMove(board.B2, board.H2), ShouldBeFalse)
				So(wPawns.ValidMove(board.B2, board.B5), ShouldBeFalse)
				So(bPawns.ValidMove(board.A1, board.B1), ShouldBeFalse)
				So(bPawns.ValidMove(board.A2, board.B3), ShouldBeFalse)
				So(bPawns.ValidMove(board.B2, board.H2), ShouldBeFalse)
				So(bPawns.ValidMove(board.B2, board.C3), ShouldBeFalse)
				So(bPawns.ValidMove(board.B2, board.B4), ShouldBeFalse)
				So(bPawns.ValidMove(board.B2, board.B5), ShouldBeFalse)

				So(wPawns.ValidMove(board.E2, board.E3), ShouldBeTrue)
				So(wPawns.ValidMove(board.E2, board.E4), ShouldBeTrue)
				So(wPawns.ValidMove(board.E2, board.D3), ShouldBeTrue)
				So(wPawns.ValidMove(board.E2, board.F3), ShouldBeTrue)

				So(bPawns.ValidMove(board.E7, board.E6), ShouldBeTrue)
				So(bPawns.ValidMove(board.E7, board.E5), ShouldBeTrue)
				So(bPawns.ValidMove(board.E7, board.F6), ShouldBeTrue)
				So(bPawns.ValidMove(board.E7, board.D6), ShouldBeTrue)
			})
			So(pawns, ShouldImplement, (*board.Piece)(nil))
		})
	})
	Convey("Given a NewPawns() function", t, func() {
		Convey("It should accept a side", func() {
			pawns := board.NewPawns(board.Black)
			So(*pawns, ShouldResemble, board.Pawns{board.NewBitboard(), board.Black})
		})
		Convey("It should Optionally accept any number of Squares to initialise the position", func() {
			pawns := board.NewPawns(board.Black, board.D8, board.A1)
			So(*pawns, ShouldResemble, board.Pawns{board.NewBitboard(board.D8, board.A1), board.Black})
		})
	})
}
