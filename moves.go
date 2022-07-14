package main

import (
	"github.com/rs/zerolog/log"
)

// nolint:cyclop,gocognit // can't really be simplified
func KnightMoves(src Square) []Square {
	moves := []Square{}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > firstRank && src.File() > BFile {
		moves = append(moves, src<<6)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > firstRank && src.File() < GFile {
		moves = append(moves, src<<10)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > secondRank && src.File() > AFile {
		moves = append(moves, src<<15)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > secondRank && src.File() < HFile {
		moves = append(moves, src<<17)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < seventhRank && src.File() > AFile {
		moves = append(moves, src>>17)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < seventhRank && src.File() < HFile {
		moves = append(moves, src>>15)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < eighthRank && src.File() > BFile {
		moves = append(moves, src>>10)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < eighthRank && src.File() < GFile {
		moves = append(moves, src>>6)
	}

	return moves
}

func OrthaganolRankMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for Rank Moves")

	moves := []Square{}

	tgtSquare := src >> 1
	for tgtSquare.Rank() == src.Rank() {
		log.Debug().
			Str("Source", boardMatrixItoS[src]).
			Str("Square:", boardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare >>= 1
	}

	tgtSquare = src << 1
	for tgtSquare.Rank() == src.Rank() {
		log.Debug().
			Str("Source", boardMatrixItoS[src]).
			Str("Square:", boardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare <<= 1
	}

	return moves
}

func OrthaganolFileMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for File Moves")

	moves := []Square{}

	// nolint:gomnd // 8 is the number of squares between Ranks
	tgtSquare := src >> 8
	for tgtSquare.File() == src.File() {
		log.Debug().
			Str("Source", boardMatrixItoS[src]).
			Str("Square:", boardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare >>= 8
	}

	// nolint:gomnd // 8 is the number of squares between Ranks
	tgtSquare = src << 8
	for tgtSquare.File() == src.File() {
		log.Debug().
			Str("Source", boardMatrixItoS[src]).
			Str("Square:", boardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare <<= 8
	}

	return moves
}

// nolint:funlen,gocognit,cyclop // Can't think how to simplify this yet
func DiagonalMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for Diagonal Moves")

	moves := []Square{}

	if src.File() < HFile && src.Rank() < eighthRank {
		// nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 7; tgtSquare <= H1 && tgtSquare >= A8; tgtSquare >>= 7 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("Source", boardMatrixItoS[src]).
				Str("Square:", boardMatrixItoS[tgtSquare]).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.File() == HFile || tgtSquare.Rank() == eighthRank {
				break
			}
		}
	}

	if src.File() > AFile && src.Rank() > firstRank {
		// nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src << 7; tgtSquare <= H1 && tgtSquare >= A8; tgtSquare <<= 7 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("Source", boardMatrixItoS[src]).
				Str("Square:", boardMatrixItoS[tgtSquare]).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.File() == HFile || tgtSquare.Rank() == firstRank {
				break
			}
		}
	}

	if src.File() > AFile && src.Rank() < eighthRank {
		// nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 9; tgtSquare <= H1 && tgtSquare >= A8; tgtSquare >>= 9 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("Source", boardMatrixItoS[src]).
				Str("Square:", boardMatrixItoS[tgtSquare]).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.File() == AFile || tgtSquare.Rank() == eighthRank {
				break
			}
		}
	}

	if src.File() < HFile && src.Rank() > firstRank {
		// nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src << 9; tgtSquare <= H1 && tgtSquare >= A8; tgtSquare <<= 9 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("Source", boardMatrixItoS[src]).
				Str("Square:", boardMatrixItoS[tgtSquare]).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.File() == AFile || tgtSquare.Rank() == firstRank {
				break
			}
		}
	}

	return moves
}
