package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// nolint:funlen // convey testing is verbose
func TestErrors(t *testing.T) {
	Convey("Given a piecePositionError", t, func() {
		ppe := &PiecePositionError{
			fen:      "fen",
			errPiece: 'k',
		}
		errMsg := fmt.Sprintf("Invalid piece found (%c) in fen (%s). "+
			"Should be one of [rnbqkpRNBQKP/12345678]",
			ppe.errPiece, ppe.fen)
		Convey("It should implement the error interface", func() {
			So(ppe, ShouldImplement, (*error)(nil))
			So(ppe.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(ppe.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an errPiece attribute", func() {
			So(ppe.errPiece, ShouldHaveSameTypeAs, 'h')
		})
	})
	Convey("Given a sideToMoveError", t, func() {
		err := &SideToMoveError{
			fen:     "fen",
			errSide: "w",
		}
		errMsg := fmt.Sprintf("Invalid side to move found (%s) in fen (%s). Should be one of [bw]", err.errSide, err.fen)
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an errSide attribute", func() {
			So(err.errSide, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a CastlingRightsError", t, func() {
		err := &CastlingRightsError{
			fen:     "fen",
			errChar: 'z',
		}
		errMsg := fmt.Sprintf("Invalid castling rights found (%c) in fen (%s). "+
			"Should only contain characters 'kqKQ-'", err.errChar, err.fen)
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an errSide attribute", func() {
			So(err.errChar, ShouldHaveSameTypeAs, 'z')
		})
	})
	Convey("Given a enPassantTargetError", t, func() {
		err := &EnPassantTargetError{
			fen:       "fen",
			errTarget: "w",
		}
		errMsg := fmt.Sprintf("Invalid en passant target square (%s) in fen (%s). "+
			"Should be one of [A3,B3,C3,D3,E3,F3,G3,H3,A6,B6,C6,D6,E6,F6,G6,H6]", err.errTarget, err.fen)
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an errTarget attribute", func() {
			So(err.errTarget, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a halfMoveClockError", t, func() {
		err := &HalfMoveClockError{
			fen:           "fen",
			err:           "w",
			halfMoveClock: "1",
		}
		errMsg := fmt.Sprintf("Invalid half move clock (%s) in fen (%s). "+
			"Must be a positive number. Received error while parsing Atoi: %s",
			err.halfMoveClock, err.fen, err.err)
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an err attribute", func() {
			So(err.err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a halfMoveClock attribute", func() {
			So(err.halfMoveClock, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a fullMoveCounterError", t, func() {
		err := &FullMoveCounterError{
			fen:             "fen",
			err:             "w",
			fullMoveCounter: "1",
		}
		errMsg := fmt.Sprintf("Invalid full move counter (%s) in fen (%s). "+
			"Must be a positive number. "+
			"Received error while parsing Atoi: %s",
			err.fullMoveCounter, err.fen, err.err)
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an err attribute", func() {
			So(err.err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a fullMoveCounter attribute", func() {
			So(err.fullMoveCounter, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a MoveError", t, func() {
		err := &MoveError{
			fen:  "fen",
			err:  "w",
			move: "e2e4",
		}
		errMsg := fmt.Sprintf("Invalid move (%s) in position %s.\n%s",
			err.move, err.fen, err.err)
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an err attribute", func() {
			So(err.err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a Move attribute", func() {
			So(err.move, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given an InvalidFenstringError", t, func() {
		err := &InvalidFenstringError{
			fen: "fen",
			err: "w",
		}
		errMsg := "Invalid Fenstring: w"
		Convey("It should implement the error interface", func() {
			So(err, ShouldImplement, (*error)(nil))
			So(err.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(err.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an err attribute", func() {
			So(err.err, ShouldHaveSameTypeAs, "h")
		})
	})
}
