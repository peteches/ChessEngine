//nolint:dupl // boilerplate for similar pieces
package board

import (
	"github.com/rs/zerolog/log"
)

type Bishops struct {
	BitBoard *BitBoard
	Colour   Side
}

func NewBishops(colour Side, startingPosition ...Square) *Bishops {
	bishops := Bishops{
		BitBoard: NewBitboard(startingPosition...),
		Colour:   colour,
	}

	return &bishops
}

func (q *Bishops) String() string {
	switch q.Colour {
	case Black:
		return "b"
	case White:
		return "B"
	default:
		return ""
	}
}

func (q *Bishops) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", src.String()).
		Str("destination", dst.String()).
		Msg("Validating Bishop move")

	for _, validDst := range DiagonalMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", src.String()).
				Str("destination", dst.String()).
				Msg("Dest is a valid diagonal Move")

			return true
		}
	}

	return false
}

func (q *Bishops) Positions() *BitBoard {
	return q.BitBoard
}
