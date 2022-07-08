package main

import (
	"context"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// nolint:gochecknoglobals // for testing this is fine
var uciOkMsg = `id name PetechesChessBot 0.0
id author Pete 'Peteches' McCabe
uciok
`

func TestUciInit(t *testing.T) {
	ctx := context.Background()

	Convey("Given a handleUci function", t, func() {
		Convey("It should accept a context", func() {
			So(func() {
				handleUci(ctx)
			}, ShouldNotPanic)
		})
		Convey("It will return a readonly channel ", func() {
			outChan := handleUci(ctx)
			So(outChan, ShouldHaveSameTypeAs, make(<-chan string))
		})
		Convey("It will write UCI info to the returned channel", func() {
			outChan := handleUci(ctx)
			out := ""
			for x := range outChan {
				out += x
			}
			So(out, ShouldEqual, uciOkMsg)
		})
	})
}

func TestScanForCommands(t *testing.T) {
	ctx := context.Background()

	Convey("Given a scanForCommands function", t, func() {
		Convey("It should accept a context an io.Reader and return a readonly string channel", func() {
			ctx, ctxCancel := context.WithCancel(ctx)
			in := strings.NewReader("hello\n")
			So(func() { scanForCommands(ctx, in) }, ShouldNotPanic)
			ctxCancel()
			ctx, ctxCancel = context.WithCancel(ctx)
			chn := scanForCommands(ctx, in)
			So(chn, ShouldHaveSameTypeAs, make(<-chan string))
			ctxCancel()
		})
		Convey("It should read from the io.Reader", func() {
			Convey("And quit when the quit command is passed in", func() {
				in := strings.NewReader("quit\nnotread\n")
				chn := scanForCommands(ctx, in)
				msg := ""
				for i := range chn {
					msg += i
				}
				So(msg, ShouldEqual, "")
			})
			Convey("And send any other messages to the returned channel", func() {
				in := strings.NewReader("bob\nquit")
				chn := scanForCommands(ctx, in)
				msg := ""
				for i := range chn {
					msg += i
				}
				So(msg, ShouldEqual, "bob")
			})
		})
	})
}
