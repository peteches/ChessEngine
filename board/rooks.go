package board

import (
	"github.com/rs/zerolog/log"
)

type Rooks struct {
	BitBoard *BitBoard
	Colour   Side
}

func NewRooks(colour Side, startingPosition ...Square) *Rooks {
	rooks := Rooks{
		BitBoard: NewBitboard(startingPosition...),
		Colour:   colour,
	}

	return &rooks
}

func (q *Rooks) String() string {
	switch q.Colour {
	case Black:
		return "r"
	case White:
		return "R"
	default:
		return ""
	}
}

func (q *Rooks) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", src.String()).
		Str("destination", dst.String()).
		Msg("Validating Queen move")

	for _, validDst := range OrthaganolFileMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", src.String()).
				Str("destination", dst.String()).
				Msg("Dest is a valid orthagonal File Move")

			return true
		}
	}

	for _, validDst := range OrthaganolRankMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", src.String()).
				Str("destination", dst.String()).
				Msg("Dest is a valid orthagonal Rank Move")

			return true
		}
	}

	return false
}

func (q *Rooks) Positions() *BitBoard {
	return q.BitBoard
}
