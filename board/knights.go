//nolint:dupl // boilerplate for similar pieces
package board

import (
	"github.com/rs/zerolog/log"
)

type Knights struct {
	Positions *BitBoard
	Colour    Side
}

func NewKnight(colour Side, startingPosition ...Square) *Knights {
	knights := Knights{
		Positions: NewBitboard(startingPosition...),
		Colour:    colour,
	}

	return &knights
}

func (q *Knights) String() string {
	switch q.Colour {
	case Black:
		return "n"
	case White:
		return "N"
	default:
		return ""
	}
}

func (q *Knights) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", BoardMatrixItoS[src]).
		Str("destination", BoardMatrixItoS[dst]).
		Msg("Validating Knight move")

	for _, validDst := range KnightMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", BoardMatrixItoS[src]).
				Str("destination", BoardMatrixItoS[dst]).
				Msg("Dest is a valid Knight Move")

			return true
		}
	}

	return false
}
