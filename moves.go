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
