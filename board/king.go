//nolint:dupl // boilerplate for similar pieces
package board

import (
	"github.com/rs/zerolog/log"
)

type King struct {
	Positions *BitBoard
	Colour    Side
}

func NewKing(colour Side, startingPosition ...Square) *King {
	kings := King{
		Positions: NewBitboard(startingPosition...),
		Colour:    colour,
	}

	return &kings
}

func (q *King) String() string {
	switch q.Colour {
	case Black:
		return "k"
	case White:
		return "K"
	default:
		return ""
	}
}

func (q *King) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", BoardMatrixItoS[src]).
		Str("destination", BoardMatrixItoS[dst]).
		Msg("Validating King move")

	for _, validDst := range KingMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", BoardMatrixItoS[src]).
				Str("destination", BoardMatrixItoS[dst]).
				Msg("Dest is a valid King Move")

			return true
		}
	}

	return false
}
