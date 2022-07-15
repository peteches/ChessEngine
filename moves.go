package main

import (
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

type Move struct {
	piece       string
	srcSquare   Square
	dstSquare   Square
	capture     bool
	promotionTo string
}

func NewMove(lanMove string) (*Move, *MoveError) {
	r := "(?i)^(?P<piece>[NBRQK])?(?P<src>[A-H][1-8])(?P<capture>[-X])?(?P<dst>[A-H][1-8])(?P<promotionTo>[NBRQ])?$"
	moveRegex := regexp.MustCompile(r)
	matches := moveRegex.FindStringSubmatch(lanMove)
	pieceIndex := moveRegex.SubexpIndex("piece")
	srcIndex := moveRegex.SubexpIndex("src")
	capIndex := moveRegex.SubexpIndex("capture")
	dstIndex := moveRegex.SubexpIndex("dst")
	promoIndex := moveRegex.SubexpIndex("promotionTo")

	src := boardMatrixStoI[strings.ToUpper(matches[srcIndex])]
	dst := boardMatrixStoI[strings.ToUpper(matches[dstIndex])]

	var piece string

	switch strings.ToUpper(matches[pieceIndex]) {
	case "":
		piece = "P"
	default:
		piece = strings.ToUpper(matches[pieceIndex])
	}

	return &Move{
		piece:       piece,
		srcSquare:   src,
		dstSquare:   dst,
		capture:     strings.ToUpper(matches[capIndex]) == "X",
		promotionTo: strings.ToUpper(matches[promoIndex]),
	}, nil
}

// nolint:cyclop,gocognit // can't really be simplified
func KnightMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for Knight Moves")

	moves := []Square{}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > firstRank && src.File() > BFile {
		moves = append(moves, src>>10)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < eighthRank && src.File() < GFile {
		moves = append(moves, src<<10)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > secondRank && src.File() < HFile {
		moves = append(moves, src>>15)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < seventhRank && src.File() > AFile {
		moves = append(moves, src<<15)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < seventhRank && src.File() < HFile {
		moves = append(moves, src<<17)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > secondRank && src.File() > AFile {
		moves = append(moves, src>>17)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < eighthRank && src.File() > BFile {
		moves = append(moves, src<<6)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > firstRank && src.File() < GFile {
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

	if src.File() < HFile && src.Rank() > firstRank {
		// nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 7; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare >>= 7 {
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
		// nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src << 7; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare <<= 7 {
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
		// nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 9; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare >>= 9 {
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

	if src.File() < HFile && src.Rank() < eighthRank {
		// nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src << 9; tgtSquare <= H8 && tgtSquare >= A1; tgtSquare <<= 9 {
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

	return moves
}

// nolint:funlen // don't think this can be simplified any more
func KingMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for King Moves")

	moves := []Square{}

	// nolint:gomnd // these are movement bit shifts
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
	case src.Rank() == firstRank:
		moves = append(moves, src>>1)
		moves = append(moves, src<<1)
		moves = append(moves, src<<9)
		moves = append(moves, src<<8)
		moves = append(moves, src<<7)
	case src.Rank() == eighthRank:
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

func WhitePawnMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for White Pawn Moves")

	moves := []Square{}

	// nolint:gomnd // these are movement bit shifts
	moves = append(moves, src<<8)

	// nolint:gomnd // these are movement bit shifts
	if src.File() > AFile {
		moves = append(moves, src<<7)
	}

	// nolint:gomnd // these are movement bit shifts
	if src.File() < HFile {
		moves = append(moves, src<<9)
	}

	// nolint:gomnd // these are movement bit shifts
	if src.Rank() == secondRank {
		moves = append(moves, src<<16)
	}

	return moves
}

func BlackPawnMoves(src Square) []Square {
	log.Debug().Str("Source", boardMatrixItoS[src]).
		Msg("Checking Destinations for Black Pawn Moves")

	moves := []Square{}

	// nolint:gomnd // these are movement bit shifts
	moves = append(moves, src>>8)

	// nolint:gomnd // these are movement bit shifts
	if src.File() > AFile {
		moves = append(moves, src>>9)
	}

	// nolint:gomnd // these are movement bit shifts
	if src.File() < HFile {
		moves = append(moves, src>>7)
	}

	// nolint:gomnd // these are movement bit shifts
	if src.Rank() == seventhRank {
		moves = append(moves, src>>16)
	}

	return moves
}
