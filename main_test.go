package main

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

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

//nolint:funlen // Convey testing is verbose
func TestProcessInput(t *testing.T) {
	ctx := context.Background()

	Convey("Given a process_input function", t, func() {
		Convey("It should accept a context and a  reader string channel", func() {
			Convey("And quit when the context is cancelled.", func() {
				cmdChn := make(chan string)
				ctx, ctxCancel := context.WithCancel(ctx)
				So(func() { processInput(ctx, cmdChn); time.Sleep(1 * time.Second); ctxCancel() }, ShouldNotPanic)
			})
			Convey("And return a readonly channel", func() {
				cmdChn := make(chan string)
				ctx, ctxCancel := context.WithCancel(ctx)
				outChan := processInput(ctx, cmdChn)
				ctxCancel()

				So(outChan, ShouldHaveSameTypeAs, make(<-chan string))
			})
		})

		Convey("It should read from the string channel and ", func() {
			Convey("When the string read is 'uci'", func() {
				Convey("It should output identifying information to the returned channel", func() {
					cmdChn := make(chan string)
					ctx, ctxCancel := context.WithCancel(ctx)
					outChan := processInput(ctx, cmdChn)
					cmdChn <- "uci"
					out := ""
					x := ""
					for x != "uciok\n" {
						x = <-outChan
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
					cmdChn := make(chan string)
					var out bytes.Buffer
					ctx, ctxCancel := context.WithCancel(ctx)
					processInput(ctx, cmdChn)
					cmdChn <- "position startpos"
					ctxCancel()
					So(out.String(), ShouldEqual, "Initialising Position")
				})
			})
		})
		Convey("When given an unrecognised command", func() {
			Convey("It should output notice", func() {
				cmdChn := make(chan string)
				ctx, ctxCancel := context.WithCancel(ctx)
				outChan := processInput(ctx, cmdChn)
				cmdChn <- "Geoff"
				out := ""
				x := ""
				for x != "Received unknown CMD: Geoff\n" {
					x = <-outChan
					out += x
				}
				ctxCancel()
				So(out, ShouldEqual, "Received unknown CMD: Geoff\n")
			})
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
