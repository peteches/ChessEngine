package board

import (
	"github.com/rs/zerolog/log"
)

type Pawns struct {
	Positions *BitBoard
	Colour    Side
}

func NewPawns(colour Side, startingPosition ...Square) *Pawns {
	pawns := Pawns{
		Positions: NewBitboard(startingPosition...),
		Colour:    colour,
	}

	return &pawns
}

func (q *Pawns) String() string {
	switch q.Colour {
	case Black:
		return "p"
	case White:
		return "P"
	default:
		return ""
	}
}

func (q *Pawns) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", BoardMatrixItoS[src]).
		Str("destination", BoardMatrixItoS[dst]).
		Msg("Validating Pawn move")

	switch q.Colour {
	case White:
		for _, validDst := range WhitePawnMoves(src) {
			if dst == validDst {
				log.Debug().
					Str("source", BoardMatrixItoS[src]).
					Str("destination", BoardMatrixItoS[dst]).
					Msg("Dest is a valid white pawn Move")

				return true
			}
		}
	case Black:
		for _, validDst := range BlackPawnMoves(src) {
			if dst == validDst {
				log.Debug().
					Str("source", BoardMatrixItoS[src]).
					Str("destination", BoardMatrixItoS[dst]).
					Msg("Dest is a valid black pawn Move")

				return true
			}
		}
	}

	return false
}
