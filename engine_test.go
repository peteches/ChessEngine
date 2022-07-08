package main

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEngine(t *testing.T) {
	Convey("Given an Engine Function", t, func() {
		ctx := context.Background()
		Convey("It should accept a context", func() {
			engine(context.Background())
		})
		Convey("It should return three string channels two readonly one writeonly", func() {
			toEng, frmEng, debug := engine(context.Background())
			So(toEng, ShouldHaveSameTypeAs, make(chan<- string))
			So(frmEng, ShouldHaveSameTypeAs, make(<-chan string))
			So(debug, ShouldHaveSameTypeAs, make(<-chan string))
		})

		Convey("It should read from the string channel and ", func() {
			Convey("When the string read is 'uci'", func() {
				Convey("It should output identifying information to the returned channel", func() {
					ctx, ctxCancel := context.WithCancel(ctx)
					toEng, frmEng, _ := engine(ctx)
					toEng <- "uci"
					out := ""
					x := ""
					for x != "uciok\n" {
						x = <-frmEng
						out += x
					}
					ctxCancel()
					So(out, ShouldEqual, uciOkMsg)
				})
			})
		})
		SkipConvey("When given the command 'position'", func() {
			Convey("with the startpos argument", func() {
				Convey("It should initialise a new Position with pieces in their starting positions.", func() {
				})
			})
		})
		Convey("When given the command 'debugPosition'", func() {
			Convey("The engine will write the current position to frmEng channel as FEN", func() {
			})
			Convey("When given an unrecognised command", func() {
				Convey("It should output notice", func() {
					ctx, ctxCancel := context.WithCancel(ctx)
					toEng, frmEng, _ := engine(ctx)
					toEng <- "Geoff"
					out := ""
					x := ""
					for x != "Received unknown CMD: Geoff\n" {
						x = <-frmEng
						out += x
					}
					ctxCancel()
					So(out, ShouldEqual, "Received unknown CMD: Geoff\n")
				})
			})
		})
	})
}
