package main

import "context"

func engine(ctx context.Context) (chan<- string, <-chan string) {
	toEng := make(chan string)
	frmEng := make(chan string)

	return toEng, frmEng
}
