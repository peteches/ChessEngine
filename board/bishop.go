//nolint:dupl // boilerplate for similar pieces
package board

import (
	"github.com/rs/zerolog/log"
)

type Bishops struct {
	Positions *BitBoard
	Colour    Side
}

func NewBishops(colour Side, startingPosition ...Square) *Bishops {
	bishops := Bishops{
		Positions: NewBitboard(startingPosition...),
		Colour:    colour,
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
		Str("source", BoardMatrixItoS[src]).
		Str("destination", BoardMatrixItoS[dst]).
		Msg("Validating Bishop move")

	for _, validDst := range DiagonalMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", BoardMatrixItoS[src]).
				Str("destination", BoardMatrixItoS[dst]).
				Msg("Dest is a valid diagonal Move")

			return true
		}
	}

	return false
}
