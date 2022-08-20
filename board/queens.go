package board

import (
	"github.com/rs/zerolog/log"
)

type Queens struct {
	BitBoard *BitBoard
	Colour   Side
}

func NewQueens(colour Side, startingPosition ...Square) *Queens {
	queens := Queens{
		BitBoard: NewBitboard(startingPosition...),
		Colour:   colour,
	}

	return &queens
}

func (q *Queens) String() string {
	switch q.Colour {
	case Black:
		return "q"
	case White:
		return "Q"
	default:
		return ""
	}
}

func (q *Queens) ValidMove(src, dst Square) bool {
	log.Debug().
		Str("source", src.String()).
		Str("destination", dst.String()).
		Msg("Validating Queen move")

	for _, validDst := range DiagonalMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", src.String()).
				Str("destination", dst.String()).
				Msg("Dest is a valid diagonal Move")

			return true
		}
	}

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

func (q *Queens) Positions() *BitBoard {
	return q.BitBoard
}
