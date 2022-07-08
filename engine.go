package main

import (
	"context"
	"fmt"
	"strings"
)

func engine(ctx context.Context) (chan<- string, <-chan string, <-chan string) {
	toEng := make(chan string)
	frmEng := make(chan string)
	debug := make(chan string)

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
					case "position":
						{
							go handlePosition(ctx, words[1:])
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
