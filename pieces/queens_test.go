package pieces_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/pieces"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQueens(t *testing.T) {
	Convey("Given a queen struct", t, func() {
		queen := pieces.Queens{}
		Convey("It should have a Positions field", func() {
			So(queen.Positions, ShouldHaveSameTypeAs, &board.BitBoard{})
			var testSide board.Side
			So(queen.Colour, ShouldHaveSameTypeAs, testSide)
		})
		Convey("It should implement the Piece interface", func() {
			bQueen := pieces.NewQueen(board.Black)
			wQueen := pieces.NewQueen(board.White)
			Convey("By returning the appropriate string representation of the piece", func() {
				So(bQueen.String(), ShouldEqual, "q")
				So(wQueen.String(), ShouldEqual, "Q")
			})
			Convey("By returning validating valid moves", func() {
				So(bQueen.ValidMove(board.A1, board.A2), ShouldBeTrue)
				So(bQueen.ValidMove(board.A1, board.B1), ShouldBeTrue)
				So(bQueen.ValidMove(board.A1, board.B2), ShouldBeTrue)
				So(bQueen.ValidMove(board.B1, board.H2), ShouldBeFalse)
			})
			So(queen, ShouldImplement, (*pieces.Piece)(nil))
		})
	})
	Convey("Given a NewQueen() function", t, func() {
		Convey("It should accept a side", func() {
			queen := pieces.NewQueen(board.Black)
			So(*queen, ShouldResemble, pieces.Queens{board.NewBitboard(), board.Black})
		})
		Convey("It should Optionally accept any number of Squares to initialise the position", func() {
			queen := pieces.NewQueen(board.Black, board.D8, board.A1)
			So(*queen, ShouldResemble, pieces.Queens{board.NewBitboard(board.D8, board.A1), board.Black})
		})
	})
}
