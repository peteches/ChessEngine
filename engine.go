package main

import (
	"context"
	"fmt"
	"strings"
)

// nolint:funlen,gocognit,cyclop // not sure how to simplify this yet
func engine(ctx context.Context) (chan<- string, <-chan string, <-chan string) {
	toEng := make(chan string)
	frmEng := make(chan string)
	debug := make(chan string)

	var err error

	position := NewPosition()
	startingFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	go func() {
		defer close(frmEng)
		defer close(toEng)
		defer close(debug)

		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			case cmd := <-toEng:
				{
					words := strings.Split(cmd, " ")
					switch words[0] {
					case "uci":
						{
							go func() {
								o := handleUci(ctx)
								for x := range o {
									frmEng <- x
								}
							}()
						}
					case "printPosition":
						{
							debug <- fmt.Sprintf("info string %s", position.String())
						}
					case "position":
						{
							if words[1] == "startpos" {
								err = position.SetPositionFromFen(startingFen)
								if err != nil {
									frmEng <- fmt.Sprintf("info string Error setting position: %s", err)
								}
							} else {
								if len(words) < numFenElements+1 {
									frmEng <- "info string Error setting position: Invalid Fenstring: Missing Fen elements"

									continue
								}
								fen := strings.Join(words[1:numFenElements+1], " ")
								err = position.SetPositionFromFen(fen)
								if err != nil {
									frmEng <- fmt.Sprintf("info string Error setting position: %s", err)
								}
							}
						}
					default:
						{
							frmEng <- fmt.Sprintf("Received unknown CMD: %s\n", cmd)
						}
					}
				}
			}
		}
	}()

	return toEng, frmEng, debug
}
