package main

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var UCIOK = `id name PetechesChessBot 0.0
id author Pete 'Peteches' McCabe
uciok
`

func TestUciInit(t *testing.T) {
	Convey("Given uciInit function", t, func() {
		Convey("It will write identifying info to the given io.Writer", func() {
			var out bytes.Buffer
			uciInit(&out)
			So(out.String(), ShouldEqual, UCIOK)
		})
	})
}

func TestProcessInput(t *testing.T) {
	Convey("Given a process_input function", t, func() {
		Convey("When given the command 'uci'", func() {
			Convey("It should output identifying information", func() {
				var out bytes.Buffer
				process_input(&out, "uci")
				So(out.String(), ShouldEqual, UCIOK)
			})
		})
		Convey("When given an unrecognised command", func() {
			Convey("It should output notice", func() {
				var out bytes.Buffer
				process_input(&out, "Geoff")
				So(out.String(), ShouldEqual, "Recieved unknown CMD: Geoff\n")
			})
		})
	})
}
