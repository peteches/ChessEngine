//nolint:dupl // boilerplate for similar pieces
package board

import (
	"github.com/rs/zerolog/log"
)

type Knights struct {
	BitBoard *BitBoard
	Colour   Side
}

func NewKnights(colour Side, startingPosition ...Square) *Knights {
	knights := Knights{
		BitBoard: NewBitboard(startingPosition...),
		Colour:   colour,
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
		Str("source", src.String()).
		Str("destination", dst.String()).
		Msg("Validating Knight move")

	for _, validDst := range KnightMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", src.String()).
				Str("destination", dst.String()).
				Msg("Dest is a valid Knight Move")

			return true
		}
	}

	return false
}

func (q *Knights) Positions() *BitBoard {
	return q.BitBoard
}
