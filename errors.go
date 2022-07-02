package main

import "fmt"

type piecePositionError struct {
	fen      string
	errPiece rune
}

func (e *piecePositionError) Error() string {
	return fmt.Sprintf("Invalid piece found (%c) in fen (%s). Should be one of [rnbqkpRNBQKP/12345678]", e.errPiece, e.fen)
}

type sideToMoveError struct {
	fen     string
	errSide string
}

func (e *sideToMoveError) Error() string {
	return fmt.Sprintf("Invalid side to move found (%s) in fen (%s). Should be one of [bw]", e.errSide, e.fen)
}

type enPassantTargetError struct {
	fen       string
	errTarget string
}

func (e *enPassantTargetError) Error() string {
	return fmt.Sprintf("Invalid en passant target square (%s) in fen (%s). Should be one of [A3,B3,C3,D3,E3,F3,G3,H3,A6,B6,C6,D6,E6,F6,G6,H6]", e.errTarget, e.fen)
}

type halfMoveClockError struct {
	fen           string
	err           string
	halfMoveClock string
}

func (e *halfMoveClockError) Error() string {
	return fmt.Sprintf("Invalid half move clock (%s) in fen (%s). Must be a positive number. Recieved error while parsing Atoi: %s", e.halfMoveClock, e.fen, e.err)
}

type fullMoveCounterError struct {
	fen             string
	err             string
	fullMoveCounter string
}

func (e *fullMoveCounterError) Error() string {
	return fmt.Sprintf("Invalid full move counter (%s) in fen (%s). Must be a positive number. Recieved error while parsing Atoi: %s", e.fullMoveCounter, e.fen, e.err)
}
