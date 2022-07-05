package main

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEngine(t *testing.T) {
	Convey("Given an Engine Function", t, func() {
		Convey("It should accept a context", func() {
			engine(context.Background())
		})
		Convey("It should return two string channels one readonly one writeonly", func() {
			toEng, frmEng := engine(context.Background())
			So(toEng, ShouldHaveSameTypeAs, make(chan<- string))
			So(frmEng, ShouldHaveSameTypeAs, make(<-chan string))
		})
	})
}
