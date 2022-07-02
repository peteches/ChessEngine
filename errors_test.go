package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	Convey("Given a piecePositionError", t, func() {
		ppe := &piecePositionError{
			fen:      "fen",
			errPiece: 'k',
		}
		errMsg := fmt.Sprintf("Invalid piece found (%c) in fen (%s). Should be one of [rnbqkpRNBQKP/12345678]", ppe.errPiece, ppe.fen)
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
		e := &sideToMoveError{
			fen:     "fen",
			errSide: "w",
		}
		errMsg := fmt.Sprintf("Invalid side to move found (%s) in fen (%s). Should be one of [bw]", e.errSide, e.fen)
		Convey("It should implement the error interface", func() {
			So(e, ShouldImplement, (*error)(nil))
			So(e.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(e.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an errSide attribute", func() {
			So(e.errSide, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a enPassantTargetError", t, func() {
		e := &enPassantTargetError{
			fen:       "fen",
			errTarget: "w",
		}
		errMsg := fmt.Sprintf("Invalid en passant target square (%s) in fen (%s). Should be one of [A3,B3,C3,D3,E3,F3,G3,H3,A6,B6,C6,D6,E6,F6,G6,H6]", e.errTarget, e.fen)
		Convey("It should implement the error interface", func() {
			So(e, ShouldImplement, (*error)(nil))
			So(e.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(e.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an errTarget attribute", func() {
			So(e.errTarget, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a halfMoveClockError", t, func() {
		e := &halfMoveClockError{
			fen:           "fen",
			err:           "w",
			halfMoveClock: "1",
		}
		errMsg := fmt.Sprintf("Invalid half move clock (%s) in fen (%s). Must be a positive number. Recieved error while parsing Atoi: %s", e.halfMoveClock, e.fen, e.err)
		Convey("It should implement the error interface", func() {
			So(e, ShouldImplement, (*error)(nil))
			So(e.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(e.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an err attribute", func() {
			So(e.err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a halfMoveClock attribute", func() {
			So(e.halfMoveClock, ShouldHaveSameTypeAs, "h")
		})
	})
	Convey("Given a fullMoveCounterError", t, func() {
		e := &fullMoveCounterError{
			fen:             "fen",
			err:             "w",
			fullMoveCounter: "1",
		}
		errMsg := fmt.Sprintf("Invalid full move counter (%s) in fen (%s). Must be a positive number. Recieved error while parsing Atoi: %s", e.fullMoveCounter, e.fen, e.err)
		Convey("It should implement the error interface", func() {
			So(e, ShouldImplement, (*error)(nil))
			So(e.Error(), ShouldEqual, errMsg)
		})
		Convey("It should have a fen attribute", func() {
			So(e.fen, ShouldHaveSameTypeAs, "")
		})
		Convey("It should have an err attribute", func() {
			So(e.err, ShouldHaveSameTypeAs, "h")
		})
		Convey("It should have a fullMoveCounter attribute", func() {
			So(e.fullMoveCounter, ShouldHaveSameTypeAs, "h")
		})
	})
}
