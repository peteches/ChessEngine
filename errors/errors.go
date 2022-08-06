package errors

import "fmt"

type PiecePositionError struct {
	Fen      string
	ErrPiece rune
}

func (e *PiecePositionError) Error() string {
	return fmt.Sprintf("Invalid piece found (%c) in Fen (%s). "+
		"Should be one of [rnbqkpRNBQKP/12345678]", e.ErrPiece, e.Fen)
}

type SideToMoveError struct {
	Fen     string
	ErrSide string
}

func (e *SideToMoveError) Error() string {
	return fmt.Sprintf("Invalid side to move found (%s) in Fen (%s). "+
		"Should be one of [bw]", e.ErrSide, e.Fen)
}

type CastlingRightsError struct {
	Fen     string
	ErrChar rune
}

func (e *CastlingRightsError) Error() string {
	return fmt.Sprintf("Invalid castling rights found (%c) in Fen (%s). "+
		"Should only contain characters 'kqKQ-'", e.ErrChar, e.Fen)
}

type EnPassantTargetError struct {
	Fen       string
	ErrTarget string
}

func (e *EnPassantTargetError) Error() string {
	return fmt.Sprintf("Invalid en passant target square (%s) in Fen (%s). "+
		"Should be one of "+
		"[A3,B3,C3,D3,E3,F3,G3,H3,A6,B6,C6,D6,E6,F6,G6,H6]",
		e.ErrTarget, e.Fen)
}

type HalfMoveClockError struct {
	Fen           string
	Err           string
	HalfMoveClock string
}

func (e *HalfMoveClockError) Error() string {
	return fmt.Sprintf("Invalid half move clock (%s) in Fen (%s). "+
		"Must be a positive number. "+
		"Received Error while parsing Atoi: %s",
		e.HalfMoveClock, e.Fen, e.Err)
}

type FullMoveCounterError struct {
	Fen             string
	Err             string
	FullMoveCounter string
}

func (e *FullMoveCounterError) Error() string {
	return fmt.Sprintf("Invalid full move counter (%s) in Fen (%s). "+
		"Must be a positive number. "+
		"Received Error while parsing Atoi: %s",
		e.FullMoveCounter, e.Fen, e.Err)
}

type MoveError struct {
	Fen  string
	Err  string
	Move string
}

func (e *MoveError) Error() string {
	return fmt.Sprintf("Invalid move (%s) in position %s.\n%s",
		e.Move, e.Fen, e.Err)
}

type InvalidFenstringError struct {
	Fen string
	Err string
}

func (e *InvalidFenstringError) Error() string {
	return fmt.Sprintf("Invalid Fenstring: %s", e.Err)
}
