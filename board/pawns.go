package board

import (
	"github.com/rs/zerolog/log"
)

type Pawns struct {
	BitBoard *BitBoard
	Colour   Side
}

func NewPawns(colour Side, startingPosition ...Square) *Pawns {
	pawns := Pawns{
		BitBoard: NewBitboard(startingPosition...),
		Colour:   colour,
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

// nolint:cyclop // unable to reduce
func (q *Pawns) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", src.String()).
		Str("destination", dst.String()).
		Msg("Validating Pawn move")

	switch q.Colour {
	case White:
		for _, validDst := range WhitePawnAdvanceMoves(src) {
			if dst == validDst {
				log.Debug().
					Str("source", src.String()).
					Str("destination", dst.String()).
					Msg("Dest is a valid white pawn Move")

				return true
			}
		}

		for _, validDst := range WhitePawnCaptureMoves(src) {
			if dst == validDst {
				log.Debug().
					Str("source", src.String()).
					Str("destination", dst.String()).
					Msg("Dest is a valid white pawn Move")

				return true
			}
		}
	case Black:
		for _, validDst := range BlackPawnCaptureMoves(src) {
			if dst == validDst {
				log.Debug().
					Str("source", src.String()).
					Str("destination", dst.String()).
					Msg("Dest is a valid black pawn Move")

				return true
			}
		}

		for _, validDst := range BlackPawnAdvanceMoves(src) {
			if dst == validDst {
				log.Debug().
					Str("source", src.String()).
					Str("destination", dst.String()).
					Msg("Dest is a valid black pawn Move")

				return true
			}
		}
	}

	return false
}

func (q *Pawns) Positions() *BitBoard {
	return q.BitBoard
}
