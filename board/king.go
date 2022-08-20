//nolint:dupl // boilerplate for similar pieces
package board

import (
	"github.com/rs/zerolog/log"
)

type King struct {
	BitBoard *BitBoard
	Colour   Side
}

func NewKing(colour Side, startingPosition ...Square) *King {
	kings := King{
		BitBoard: NewBitboard(startingPosition...),
		Colour:   colour,
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
		Str("source", src.String()).
		Str("destination", dst.String()).
		Msg("Validating King move")

	for _, validDst := range KingMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", src.String()).
				Str("destination", dst.String()).
				Msg("Dest is a valid King Move")

			return true
		}
	}

	return false
}

func (q *King) Positions() *BitBoard {
	return q.BitBoard
}
