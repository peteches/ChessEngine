package board

import (
	"github.com/rs/zerolog/log"
)

//nolint:cyclop,gocognit // can't really be simplified
func KnightMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for Knight Moves")

	moves := []Square{}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > FirstRank && src.File() > BFile {
		moves = append(moves, src>>10)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < EighthRank && src.File() < GFile {
		moves = append(moves, src<<10)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > SecondRank && src.File() < HFile {
		moves = append(moves, src>>15)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < SeventhRank && src.File() > AFile {
		moves = append(moves, src<<15)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < SeventhRank && src.File() < HFile {
		moves = append(moves, src<<17)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > SecondRank && src.File() > AFile {
		moves = append(moves, src>>17)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < EighthRank && src.File() > BFile {
		moves = append(moves, src<<6)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > FirstRank && src.File() < GFile {
		moves = append(moves, src>>6)
	}

	return moves
}

func OrthaganolRankMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for Rank Moves")

	moves := []Square{}

	tgtSquare := src >> 1
	for tgtSquare.Rank() == src.Rank() {
		log.Debug().
			Str("Source", src.String()).
			Str("Square:", tgtSquare.String()).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare >>= 1
	}

	tgtSquare = src << 1
	for tgtSquare.Rank() == src.Rank() {
		log.Debug().
			Str("Source", src.String()).
			Str("Square:", tgtSquare.String()).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare <<= 1
	}

	return moves
}

func OrthaganolFileMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for File Moves")

	moves := []Square{}

	//nolint:gomnd // 8 is the number of squares between Ranks
	tgtSquare := src >> 8
	for tgtSquare.File() == src.File() {
		log.Debug().
			Str("Source", src.String()).
			Str("Square:", tgtSquare.String()).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare >>= 8
	}

	//nolint:gomnd // 8 is the number of squares between Ranks
	tgtSquare = src << 8
	for tgtSquare.File() == src.File() {
		log.Debug().
			Str("Source", src.String()).
			Str("Square:", tgtSquare.String()).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare <<= 8
	}

	return moves
}

// nolint: dupl,cyclop // These are similar but different.
func DiagonalMovesA1H8(src Square) []Square {
	moves := []Square{}

	if src.File() > AFile && src.Rank() > FirstRank {
		//nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 9; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare >>= 9 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", ">>9").
				Str("Source", src.String()).
				Str("Square:", tgtSquare.String()).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	if src.File() < HFile && src.Rank() < EighthRank {
		//nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src << 9; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare <<= 9 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", ">>9").
				Str("Source", src.String()).
				Str("Square:", tgtSquare.String()).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	return moves
}

// nolint: dupl,cyclop // These are similar but different.
func DiagonalMovesA8H1(src Square) []Square {
	moves := []Square{}

	if src.File() < HFile && src.Rank() > FirstRank {
		//nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 7; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare >>= 7 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", ">>7").
				Str("Source", src.String()).
				Str("Square:", tgtSquare.String()).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	if src.File() > AFile && src.Rank() < EighthRank {
		//nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src << 7; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare <<= 7 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", "<<7").
				Str("Source", src.String()).
				Str("Square:", tgtSquare.String()).
				Uint64("SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	return moves
}

//nolint:funlen,gocognit,cyclop // Can't think how to simplify this yet
func DiagonalMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for Diagonal Moves")

	moves := []Square{}

	moves = append(moves, DiagonalMovesA8H1(src)...)
	moves = append(moves, DiagonalMovesA1H8(src)...)

	return moves
}

//nolint:funlen // don't think this can be simplified any more
func KingMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for King Moves")

	moves := []Square{}

	//nolint:gomnd // these are movement bit shifts
	switch {
	case src == A1:
		moves = append(moves, src<<9)
		moves = append(moves, src<<8)
		moves = append(moves, src<<1)

	case src == A8:
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
		moves = append(moves, src<<1)
	case src == H1:
		moves = append(moves, src<<8)
		moves = append(moves, src<<7)
		moves = append(moves, src>>1)
	case src == H8:
		moves = append(moves, src>>9)
		moves = append(moves, src>>8)
		moves = append(moves, src>>1)
	case src.Rank() == FirstRank:
		moves = append(moves, src>>1)
		moves = append(moves, src<<1)
		moves = append(moves, src<<9)
		moves = append(moves, src<<8)
		moves = append(moves, src<<7)
	case src.Rank() == EighthRank:
		moves = append(moves, src>>1)
		moves = append(moves, src<<1)
		moves = append(moves, src>>9)
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
	case src.File() == AFile:
		moves = append(moves, src<<8)
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
		moves = append(moves, src<<1)
		moves = append(moves, src<<9)
	case src.File() == HFile:
		moves = append(moves, src<<8)
		moves = append(moves, src>>8)
		moves = append(moves, src<<7)
		moves = append(moves, src>>1)
		moves = append(moves, src>>9)

	default:
		moves = append(moves, src<<9)
		moves = append(moves, src<<8)
		moves = append(moves, src<<7)
		moves = append(moves, src<<1)
		moves = append(moves, src>>9)
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
		moves = append(moves, src>>1)
	}

	return moves
}

func WhitePawnAdvanceMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for White Pawn Moves")

	moves := []Square{}

	//nolint:gomnd // these are movement bit shifts
	moves = append(moves, src<<8)

	//nolint:gomnd // these are movement bit shifts
	if src.Rank() == SecondRank {
		moves = append(moves, src<<16)
	}

	return moves
}

func WhitePawnCaptureMoves(src Square) []Square {
	moves := []Square{}

	//nolint:gomnd // these are movement bit shifts
	if src.File() > AFile {
		moves = append(moves, src<<7)
	}

	//nolint:gomnd // these are movement bit shifts
	if src.File() < HFile {
		moves = append(moves, src<<9)
	}

	return moves
}

func BlackPawnCaptureMoves(src Square) []Square {
	moves := []Square{}

	//nolint:gomnd // these are movement bit shifts
	if src.File() > AFile {
		moves = append(moves, src>>9)
	}

	//nolint:gomnd // these are movement bit shifts
	if src.File() < HFile {
		moves = append(moves, src>>7)
	}

	return moves
}

func BlackPawnAdvanceMoves(src Square) []Square {
	log.Debug().Str("Source", src.String()).
		Msg("Checking Destinations for Black Pawn Moves")

	moves := []Square{}

	//nolint:gomnd // these are movement bit shifts
	moves = append(moves, src>>8)

	//nolint:gomnd // these are movement bit shifts
	if src.Rank() == SeventhRank {
		moves = append(moves, src>>16)
	}

	return moves
}
