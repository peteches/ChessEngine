package pieces

import (
	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/moves"
	"github.com/rs/zerolog/log"
)

type Queens struct {
	Positions *board.BitBoard
	Colour    board.Side
}

func NewQueen(colour board.Side, startingPosition ...board.Square) *Queens {
	queens := Queens{
		Positions: board.NewBitboard(startingPosition...),
		Colour:    colour,
	}

	return &queens
}

func (q *Queens) String() string {
	switch q.Colour {
	case board.Black:
		return "q"
	case board.White:
		return "Q"
	}

	return "Unknown"
}

func (q *Queens) ValidMove(src, dst board.Square) bool {
	log.Debug().
		Str("source", board.BoardMatrixItoS[src]).
		Str("destination", board.BoardMatrixItoS[dst]).
		Msg("Validating Queen move")

	for _, validDst := range moves.DiagonalMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", board.BoardMatrixItoS[src]).
				Str("destination", board.BoardMatrixItoS[dst]).
				Msg("Dest is a valid diagonal Move")

			return true
		}
	}

	for _, validDst := range moves.OrthaganolFileMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", board.BoardMatrixItoS[src]).
				Str("destination", board.BoardMatrixItoS[dst]).
				Msg("Dest is a valid orthagonal File Move")

			return true
		}
	}

	for _, validDst := range moves.OrthaganolRankMoves(src) {
		if dst == validDst {
			log.Debug().
				Str("source", board.BoardMatrixItoS[src]).
				Str("destination", board.BoardMatrixItoS[dst]).
				Msg("Dest is a valid orthagonal Rank Move")

			return true
		}
	}

	return false
}
