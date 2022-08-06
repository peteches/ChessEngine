package errors_test

import (
	"fmt"
	"testing"

	"github.com/peteches/ChessEngine/errors"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:funlen // convey testing is verbose
func TestErrors(t *testing.T) {
	Convey("Given a piecePositionError", t, func() {
		ppe := &errors.PiecePositionError{
			Fen:      "Fen",
			ErrPiece: 'k',
		}
		ErrMsg := fmt.Sprintf("Invalid piece found (%c) in Fen (%s). "+
			"Should be one of [rnbqkpRNBQKP/12345678]",
			ppe.ErrPiece, ppe.Fen)
		Convey("It should implement the Error interface", func() {
			So(ppe, ShouldImplement, (*error)(nil))
			So(ppe.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(ppe.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an ErrPiece attribute", func() {
			So(ppe.ErrPiece, ShouldHaveSameTypeAs, 'h')
		})
	})
	Convey("Given a sideToMoveError", t, func() {
		Err := &errors.SideToMoveError{
			Fen:     "Fen",
			ErrSide: "w",
		}
		ErrMsg := fmt.Sprintf("Invalid side to move found (%s) in Fen (%s). Should be one of [bw]", Err.ErrSide, Err.Fen)
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an ErrSide attribute", func() {
			So(Err.ErrSide, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a CastlingRightsError", t, func() {
		Err := &errors.CastlingRightsError{
			Fen:     "Fen",
			ErrChar: 'z',
		}
		ErrMsg := fmt.Sprintf("Invalid castling rights found (%c) in Fen (%s). "+
			"Should only contain characters 'kqKQ-'", Err.ErrChar, Err.Fen)
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an ErrSide attribute", func() {
			So(Err.ErrChar, ShouldHaveSameTypeAs, 'z')
		})
	})
	Convey("Given a enPassantTargetError", t, func() {
		Err := &errors.EnPassantTargetError{
			Fen:       "Fen",
			ErrTarget: "w",
		}
		ErrMsg := fmt.Sprintf("Invalid en passant target square (%s) in Fen (%s). "+
			"Should be one of [A3,B3,C3,D3,E3,F3,G3,H3,A6,B6,C6,D6,E6,F6,G6,H6]", Err.ErrTarget, Err.Fen)
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an ErrTarget attribute", func() {
			So(Err.ErrTarget, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a halfMoveClockError", t, func() {
		Err := &errors.HalfMoveClockError{
			Fen:           "Fen",
			Err:           "w",
			HalfMoveClock: "1",
		}
		ErrMsg := fmt.Sprintf("Invalid half move clock (%s) in Fen (%s). "+
			"Must be a positive number. Received Error while parsing Atoi: %s",
			Err.HalfMoveClock, Err.Fen, Err.Err)
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an Err attribute", func() {
			So(Err.Err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a halfMoveClock attribute", func() {
			So(Err.HalfMoveClock, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a fullMoveCounterError", t, func() {
		Err := &errors.FullMoveCounterError{
			Fen:             "Fen",
			Err:             "w",
			FullMoveCounter: "1",
		}
		ErrMsg := fmt.Sprintf("Invalid full move counter (%s) in Fen (%s). "+
			"Must be a positive number. "+
			"Received Error while parsing Atoi: %s",
			Err.FullMoveCounter, Err.Fen, Err.Err)
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an Err attribute", func() {
			So(Err.Err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a fullMoveCounter attribute", func() {
			So(Err.FullMoveCounter, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a MoveError", t, func() {
		Err := &errors.MoveError{
			Fen:  "Fen",
			Err:  "w",
			Move: "e2e4",
		}
		ErrMsg := fmt.Sprintf("Invalid move (%s) in position %s.\n%s",
			Err.Move, Err.Fen, Err.Err)
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an Err attribute", func() {
			So(Err.Err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a Move attribute", func() {
			So(Err.Move, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given an InvalidFenstringError", t, func() {
		Err := &errors.InvalidFenstringError{
			Fen: "Fen",
			Err: "w",
		}
		ErrMsg := "Invalid Fenstring: w"
		Convey("It should implement the Error interface", func() {
			So(Err, ShouldImplement, (*error)(nil))
			So(Err.Error(), ShouldEqual, ErrMsg)
		})
		Convey("It should have a Fen attribute", func() {
			So(Err.Fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an Err attribute", func() {
			So(Err.Err, ShouldHaveSameTypeAs, "h")
		})
	})
}
