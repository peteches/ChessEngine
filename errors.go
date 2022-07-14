package main

import "fmt"

type PiecePositionError struct {
	fen      string
	errPiece rune
}

func (e *PiecePositionError) Error() string {
	return fmt.Sprintf("Invalid piece found (%c) in fen (%s). "+
		"Should be one of [rnbqkpRNBQKP/12345678]", e.errPiece, e.fen)
}

type SideToMoveError struct {
	fen     string
	errSide string
}

func (e *SideToMoveError) Error() string {
	return fmt.Sprintf("Invalid side to move found (%s) in fen (%s). "+
		"Should be one of [bw]", e.errSide, e.fen)
}

type CastlingRightsError struct {
	fen     string
	errChar rune
}

func (e *CastlingRightsError) Error() string {
	return fmt.Sprintf("Invalid castling rights found (%c) in fen (%s). "+
		"Should only contain characters 'kqKQ-'", e.errChar, e.fen)
}

type EnPassantTargetError struct {
	fen       string
	errTarget string
}

func (e *EnPassantTargetError) Error() string {
	return fmt.Sprintf("Invalid en passant target square (%s) in fen (%s). "+
		"Should be one of "+
		"[A3,B3,C3,D3,E3,F3,G3,H3,A6,B6,C6,D6,E6,F6,G6,H6]",
		e.errTarget, e.fen)
}

type HalfMoveClockError struct {
	fen           string
	err           string
	halfMoveClock string
}

func (e *HalfMoveClockError) Error() string {
	return fmt.Sprintf("Invalid half move clock (%s) in fen (%s). "+
		"Must be a positive number. "+
		"Received error while parsing Atoi: %s",
		e.halfMoveClock, e.fen, e.err)
}

type FullMoveCounterError struct {
	fen             string
	err             string
	fullMoveCounter string
}

func (e *FullMoveCounterError) Error() string {
	return fmt.Sprintf("Invalid full move counter (%s) in fen (%s). "+
		"Must be a positive number. "+
		"Received error while parsing Atoi: %s",
		e.fullMoveCounter, e.fen, e.err)
}

type MoveError struct {
	fen  string
	err  string
	move string
}

func (e *MoveError) Error() string {
	return fmt.Sprintf("Invalid move (%s) in position %s.\n%s",
		e.move, e.fen, e.err)
}

type InvalidFenstringError struct {
	fen string
	err string
}

func (e *InvalidFenstringError) Error() string {
	return fmt.Sprintf("Invalid Fenstring: %s", e.err)
}
