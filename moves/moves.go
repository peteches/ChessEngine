package moves

import (
	"regexp"
	"strings"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/errors"
	"github.com/rs/zerolog/log"
)

type Move struct {
	Piece       string
	SrcSquare   board.Square
	DstSquare   board.Square
	Capture     bool
	PromotionTo string
}

func NewMove(lanMove string) (*Move, *errors.MoveError) {
	log.Debug().Str("move", lanMove).Msg("Creating new Move struct")

	r := "(?i)^(?P<piece>[NBRQK])?(?P<src>[A-H][1-8])(?P<capture>[-X])?(?P<dst>[A-H][1-8])(?P<promotionTo>[NBRQ])?$"
	moveRegex := regexp.MustCompile(r)
	matches := moveRegex.FindStringSubmatch(lanMove)

	pieceIndex := moveRegex.SubexpIndex("piece")
	srcIndex := moveRegex.SubexpIndex("src")
	capIndex := moveRegex.SubexpIndex("capture")
	dstIndex := moveRegex.SubexpIndex("dst")
	promoIndex := moveRegex.SubexpIndex("promotionTo")

	log.Debug().
		Int("pieceIndex", pieceIndex).
		Int("srcIndex", srcIndex).
		Int("capIndex", capIndex).
		Int("dstIndex", dstIndex).
		Int("promoIndex", promoIndex).
		Str("piece", matches[pieceIndex]).
		Str("src", matches[srcIndex]).
		Str("cap", matches[capIndex]).
		Str("dst", matches[dstIndex]).
		Str("promo", matches[promoIndex]).
		Interface("RegexMatches", matches).
		Msg("regex indexes")

	src := board.BoardMatrixStoI[strings.ToUpper(matches[srcIndex])]
	dst := board.BoardMatrixStoI[strings.ToUpper(matches[dstIndex])]

	var piece string

	switch strings.ToUpper(matches[pieceIndex]) {
	case "":
		piece = "P"
		if !ValidMove(piece, src, dst) {
			return nil, &errors.MoveError{
				Move: lanMove,
				Err:  "Pawns do not move like that",
			}
		}
	default:
		piece = strings.ToUpper(matches[pieceIndex])
	}

	return &Move{
		Piece:       piece,
		SrcSquare:   src,
		DstSquare:   dst,
		Capture:     strings.ToUpper(matches[capIndex]) == "X",
		PromotionTo: strings.ToUpper(matches[promoIndex]),
	}, nil
}

/*
 ValidMove will check a move to ensure it could be valid in principal. This means it will check that the
 Source and Destination squares are valid for the piece in question. It does not check to see if the move is legal.

 So for instance moving a pawn from e3-e7 is invalid because pawns can only move 1 square forward once they are off
 their starting position e2-e4 is a valid pawn move, however if the king is in check or another piece is on the e3
 or e4 squares already the move would not be legal, but would still be considered valid.

 This function is mean to quickly check moves are valid before doing a more expensive check on the legallity of a move.
*/
//nolint:gocognit,cyclop // not sure if this can actually be simplified any.
func ValidMove(piece string, src, dst board.Square) bool {
	switch piece {
	case "P":
		for _, validDst := range WhitePawnMoves(src) {
			if dst == validDst {
				return true
			}
		}

		for _, validDst := range BlackPawnMoves(src) {
			if dst == validDst {
				return true
			}
		}
	case "N":
		for _, validDst := range KnightMoves(src) {
			if dst == validDst {
				return true
			}
		}
	case "B":
		for _, validDst := range DiagonalMoves(src) {
			if dst == validDst {
				return true
			}
		}
	case "R":
		for _, validDst := range OrthaganolFileMoves(src) {
			if dst == validDst {
				return true
			}
		}

		for _, validDst := range OrthaganolRankMoves(src) {
			if dst == validDst {
				return true
			}
		}
	case "K":
		for _, validDst := range KingMoves(src) {
			if dst == validDst {
				return true
			}
		}
	case "Q":
	}

	return false
}

//nolint:cyclop,gocognit // can't really be simplified
func KnightMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for Knight Moves")

	moves := []board.Square{}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > board.FirstRank && src.File() > board.BFile {
		moves = append(moves, src>>10)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < board.EighthRank && src.File() < board.GFile {
		moves = append(moves, src<<10)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > board.SecondRank && src.File() < board.HFile {
		moves = append(moves, src>>15)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < board.SeventhRank && src.File() > board.AFile {
		moves = append(moves, src<<15)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < board.SeventhRank && src.File() < board.HFile {
		moves = append(moves, src<<17)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > board.SecondRank && src.File() > board.AFile {
		moves = append(moves, src>>17)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() < board.EighthRank && src.File() > board.BFile {
		moves = append(moves, src<<6)
	}

	//nolint:gomnd // not sure how to name these as constants
	if src.Rank() > board.FirstRank && src.File() < board.GFile {
		moves = append(moves, src>>6)
	}

	return moves
}

func OrthaganolRankMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for Rank Moves")

	moves := []board.Square{}

	tgtSquare := src >> 1
	for tgtSquare.Rank() == src.Rank() {
		log.Debug().
			Str("Source", board.BoardMatrixItoS[src]).
			Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare >>= 1
	}

	tgtSquare = src << 1
	for tgtSquare.Rank() == src.Rank() {
		log.Debug().
			Str("Source", board.BoardMatrixItoS[src]).
			Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare <<= 1
	}

	return moves
}

func OrthaganolFileMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for File Moves")

	moves := []board.Square{}

	//nolint:gomnd // 8 is the number of squares between Ranks
	tgtSquare := src >> 8
	for tgtSquare.File() == src.File() {
		log.Debug().
			Str("Source", board.BoardMatrixItoS[src]).
			Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare >>= 8
	}

	//nolint:gomnd // 8 is the number of squares between Ranks
	tgtSquare = src << 8
	for tgtSquare.File() == src.File() {
		log.Debug().
			Str("Source", board.BoardMatrixItoS[src]).
			Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
			Msg("Adding square")

		moves = append(moves, tgtSquare)
		tgtSquare <<= 8
	}

	return moves
}

//nolint:funlen,gocognit,cyclop // Can't think how to simplify this yet
func DiagonalMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for Diagonal Moves")

	moves := []board.Square{}

	if src.File() < board.HFile && src.Rank() > board.FirstRank {
		//nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 7; tgtSquare <= board.H8 && tgtSquare >= board.A1; tgtSquare >>= 7 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", ">>7").
				Str("Source", board.BoardMatrixItoS[src]).
				Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
				Uint64("board.SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	if src.File() > board.AFile && src.Rank() < board.EighthRank {
		//nolint:gomnd // 7 is one of the two numbers required to move diagonally
		for tgtSquare := src << 7; tgtSquare <= board.H8 && tgtSquare >= board.A1; tgtSquare <<= 7 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", "<<7").
				Str("Source", board.BoardMatrixItoS[src]).
				Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
				Uint64("board.SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	if src.File() > board.AFile && src.Rank() > board.FirstRank {
		//nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src >> 9; tgtSquare <= board.H8 && tgtSquare >= board.A1; tgtSquare >>= 9 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", ">>9").
				Str("Source", board.BoardMatrixItoS[src]).
				Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
				Uint64("board.SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	if src.File() < board.HFile && src.Rank() < board.EighthRank {
		//nolint:gomnd // 9 is one of the two numbers required to move diagonally
		for tgtSquare := src << 9; tgtSquare <= board.H8 && tgtSquare >= board.A1; tgtSquare <<= 9 {
			moves = append(moves, tgtSquare)

			log.Debug().
				Str("BitShift", ">>9").
				Str("Source", board.BoardMatrixItoS[src]).
				Str("board.Square:", board.BoardMatrixItoS[tgtSquare]).
				Uint64("board.SquareUint64", uint64(tgtSquare)).
				Msg("Adding square")

			if tgtSquare.OnEdge() {
				break
			}
		}
	}

	return moves
}

//nolint:funlen // don't think this can be simplified any more
func KingMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for King Moves")

	moves := []board.Square{}

	//nolint:gomnd // these are movement bit shifts
	switch {
	case src == board.A1:
		moves = append(moves, src<<9)
		moves = append(moves, src<<8)
		moves = append(moves, src<<1)

	case src == board.A8:
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
		moves = append(moves, src<<1)
	case src == board.H1:
		moves = append(moves, src<<8)
		moves = append(moves, src<<7)
		moves = append(moves, src>>1)
	case src == board.H8:
		moves = append(moves, src>>9)
		moves = append(moves, src>>8)
		moves = append(moves, src>>1)
	case src.Rank() == board.FirstRank:
		moves = append(moves, src>>1)
		moves = append(moves, src<<1)
		moves = append(moves, src<<9)
		moves = append(moves, src<<8)
		moves = append(moves, src<<7)
	case src.Rank() == board.EighthRank:
		moves = append(moves, src>>1)
		moves = append(moves, src<<1)
		moves = append(moves, src>>9)
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
	case src.File() == board.AFile:
		moves = append(moves, src<<8)
		moves = append(moves, src>>8)
		moves = append(moves, src>>7)
		moves = append(moves, src<<1)
		moves = append(moves, src<<9)
	case src.File() == board.HFile:
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

func WhitePawnMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for White Pawn Moves")

	moves := []board.Square{}

	//nolint:gomnd // these are movement bit shifts
	moves = append(moves, src<<8)

	//nolint:gomnd // these are movement bit shifts
	if src.File() > board.AFile {
		moves = append(moves, src<<7)
	}

	//nolint:gomnd // these are movement bit shifts
	if src.File() < board.HFile {
		moves = append(moves, src<<9)
	}

	//nolint:gomnd // these are movement bit shifts
	if src.Rank() == board.SecondRank {
		moves = append(moves, src<<16)
	}

	return moves
}

func BlackPawnMoves(src board.Square) []board.Square {
	log.Debug().Str("Source", board.BoardMatrixItoS[src]).
		Msg("Checking Destinations for Black Pawn Moves")

	moves := []board.Square{}

	//nolint:gomnd // these are movement bit shifts
	moves = append(moves, src>>8)

	//nolint:gomnd // these are movement bit shifts
	if src.File() > board.AFile {
		moves = append(moves, src>>9)
	}

	//nolint:gomnd // these are movement bit shifts
	if src.File() < board.HFile {
		moves = append(moves, src>>7)
	}

	//nolint:gomnd // these are movement bit shifts
	if src.Rank() == board.SeventhRank {
		moves = append(moves, src>>16)
	}

	return moves
}
