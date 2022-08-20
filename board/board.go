package board

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/peteches/ChessEngine/errors"
	"github.com/rs/zerolog/log"
)

type BitBoard struct {
	Board uint64
}

func NewBitboard(initPositions ...Square) *BitBoard {
	bitBoard := BitBoard{
		Board: 0,
	}
	for _, x := range initPositions {
		bitBoard.FlipBit(x)
	}

	return &bitBoard
}

func (bb *BitBoard) FlipBit(bit Square) {
	bb.Board ^= uint64(bit)
}

func (bb *BitBoard) Squares() []Square {
	squares := []Square{}

	for _, square := range AllSquares {
		if (bb.Board & uint64(square)) > 0 {
			squares = append(squares, square)
		}
	}

	return squares
}

func (bb *BitBoard) Occupied(sqr Square) bool {
	return (bb.Board & uint64(sqr)) > 0
}

type Side uint8

const (
	White Side = iota
	Black
)

type Board struct {
	WhiteKing *King
	BlackKing *King

	WhiteQueens *Queens
	BlackQueens *Queens

	WhiteBishops *Bishops
	BlackBishops *Bishops

	WhiteKnights *Knights
	BlackKnights *Knights

	WhiteRooks *Rooks
	BlackRooks *Rooks

	WhitePawns *Pawns
	BlackPawns *Pawns
}

func NewBoard() *Board {
	return &Board{
		WhiteKing:    NewKing(White),
		BlackKing:    NewKing(Black),
		WhiteQueens:  NewQueens(White),
		BlackQueens:  NewQueens(Black),
		WhiteBishops: NewBishops(White),
		BlackBishops: NewBishops(Black),
		WhiteKnights: NewKnights(White),
		BlackKnights: NewKnights(Black),
		WhiteRooks:   NewRooks(White),
		BlackRooks:   NewRooks(Black),
		WhitePawns:   NewPawns(White),
		BlackPawns:   NewPawns(Black),
	}
}

// nolint: cyclop // can't be simplified any further
func (b *Board) OccupiedBy(sqr Square) string {
	switch {
	case b.WhiteKing.BitBoard.Occupied(sqr):
		return b.WhiteKing.String()
	case b.WhiteQueens.Positions().Occupied(sqr):
		return b.WhiteQueens.String()
	case b.WhiteBishops.BitBoard.Occupied(sqr):
		return b.WhiteBishops.String()
	case b.WhiteKnights.BitBoard.Occupied(sqr):
		return b.WhiteKnights.String()
	case b.WhiteRooks.BitBoard.Occupied(sqr):
		return b.WhiteRooks.String()
	case b.WhitePawns.BitBoard.Occupied(sqr):
		return b.WhitePawns.String()
	case b.BlackKing.BitBoard.Occupied(sqr):
		return b.BlackKing.String()
	case b.BlackQueens.Positions().Occupied(sqr):
		return b.BlackQueens.String()
	case b.BlackBishops.BitBoard.Occupied(sqr):
		return b.BlackBishops.String()
	case b.BlackKnights.BitBoard.Occupied(sqr):
		return b.BlackKnights.String()
	case b.BlackRooks.BitBoard.Occupied(sqr):
		return b.BlackRooks.String()
	case b.BlackPawns.BitBoard.Occupied(sqr):
		return b.BlackPawns.String()
	default:
		return ""
	}
}

func (b *Board) Occupied(sqr Square) bool {
	occupiedBy := b.OccupiedBy(sqr)

	return len(occupiedBy) > 0
}

func (b *Board) String() string {
	fen := ""
	unoccupied := 0

	for idx, sqr := range AllSquares {
		if b.Occupied(sqr) {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			fen += b.OccupiedBy(sqr)
		} else {
			unoccupied++
		}

		if (idx+1)%8 == 0 {
			if unoccupied > 0 {
				fen += strconv.Itoa(unoccupied)
				unoccupied = 0
			}

			if (idx + 1) < TotalSquares {
				fen += "/"
			}
		}
	}

	return fen
}

// nolint:funlen,cyclop // TODO look are refactoring this
func (b *Board) SetPieces(pieces string) *errors.PiecePositionError {
	// reset all positions
	b.BlackRooks.BitBoard.Board &= uint64(0)
	b.BlackKnights.BitBoard.Board &= uint64(0)
	b.BlackBishops.BitBoard.Board &= uint64(0)
	b.BlackQueens.Positions().Board &= uint64(0)
	b.BlackKing.BitBoard.Board &= uint64(0)
	b.BlackPawns.BitBoard.Board &= uint64(0)
	b.WhiteRooks.BitBoard.Board &= uint64(0)
	b.WhiteKnights.BitBoard.Board &= uint64(0)
	b.WhiteBishops.BitBoard.Board &= uint64(0)
	b.WhiteQueens.Positions().Board &= uint64(0)
	b.WhiteKing.BitBoard.Board &= uint64(0)
	b.WhitePawns.BitBoard.Board &= uint64(0)

	offset := 0

	// fen strings are odd. They move from A8-H8,A7-H7 etc.
	// I want them to go A1-H1,A2-H2 etc. This then makes more sense when
	// using the Constants A1-H8 where A1 is the smallest and H8 is the
	// largest. Using fenstrings as is A8 is smaller than A1, which isn't
	// very intuitive. This does mean I need to manipulate the fen string
	// slightly to allow more intuitive BitBoards
	pieceRanks := strings.Split(pieces, "/")
	piecesSensibleOrder := ""

	for i := 7; i >= 0; i-- {
		piecesSensibleOrder += pieceRanks[i]
	}

	// this case statement pushes over the limit for funlen
	// but is largely unavoidable. complexity in this function is as simple as can be made.
	for index, pos := range piecesSensibleOrder {
		switch pos {
		case 'r':
			{
				b.BlackRooks.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'n':
			{
				b.BlackKnights.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'b':
			{
				b.BlackBishops.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'q':
			{
				b.BlackQueens.Positions().FlipBit(Square(1 << (index + offset)))
			}
		case 'k':
			{
				b.BlackKing.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'p':
			{
				b.BlackPawns.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'R':
			{
				b.WhiteRooks.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'N':
			{
				b.WhiteKnights.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'B':
			{
				b.WhiteBishops.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'Q':
			{
				b.WhiteQueens.Positions().FlipBit(Square(1 << (index + offset)))
			}
		case 'K':
			{
				b.WhiteKing.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case 'P':
			{
				b.WhitePawns.BitBoard.FlipBit(Square(1 << (index + offset)))
			}
		case '1', '2', '3', '4', '5', '6', '7', '8':
			{
				offset += (int(pos-'0') - 1)
			}
		default:
			{
				return &errors.PiecePositionError{
					ErrPiece: pos,
				}
			}
		}
	}

	return nil
}

type parsedMove struct {
	matches   []string
	piece     string
	src       Square
	dst       Square
	capture   bool
	promotion string
}

func parseMove(lanMove string) *parsedMove {
	r := "(?i)^(?P<piece>[NBRQK])?(?P<src>[A-H][1-8])(?P<capture>[-X])?(?P<dst>[A-H][1-8])(?P<promotionTo>[NBRQ])?$"
	moveRegex := regexp.MustCompile(r)
	matches := moveRegex.FindStringSubmatch(lanMove)

	pieceIndex := moveRegex.SubexpIndex("piece")
	srcIndex := moveRegex.SubexpIndex("src")
	capIndex := moveRegex.SubexpIndex("capture")
	dstIndex := moveRegex.SubexpIndex("dst")
	promoIndex := moveRegex.SubexpIndex("promotionTo")

	log.Debug().
		Int("pieceIndex", pieceIndex).
		Int("srcIndex", srcIndex).
		Int("capIndex", capIndex).
		Int("dstIndex", dstIndex).
		Int("promoIndex", promoIndex).
		Str("piece", matches[pieceIndex]).
		Str("src", matches[srcIndex]).
		Str("cap", matches[capIndex]).
		Str("dst", matches[dstIndex]).
		Str("promo", matches[promoIndex]).
		Interface("RegexMatches", matches).
		Msg("regex indexes")

	return &parsedMove{
		matches:   matches,
		piece:     matches[pieceIndex],
		src:       BoardMatrixStoI[strings.ToUpper(matches[srcIndex])],
		dst:       BoardMatrixStoI[strings.ToUpper(matches[dstIndex])],
		capture:   matches[capIndex] != "",
		promotion: matches[promoIndex],
	}
}

func (b *Board) clearPath(src, dst Square) bool {
	log.Debug().
		Str("src", src.String()).
		Str("dst", dst.String()).
		Msg("Checking if there are any pieces between src and dst")

	inbetweenSquares := SquaresBetween(src, dst)
	if len(inbetweenSquares) == 0 && !SquaresAdjacent(src, dst) {
		return false
	}

	for _, inbetweenSqr := range inbetweenSquares {
		if b.Occupied(inbetweenSqr) {
			return false
		}
	}

	return true
}

func (b *Board) rookCheck(checkingSide Side, kingSqr Square) bool {
	var rooks *Rooks

	switch checkingSide {
	case Black:
		rooks = b.BlackRooks
	case White:
		rooks = b.WhiteRooks
	}

	for _, rookSqr := range rooks.BitBoard.Squares() {
		log.Debug().Msgf("Checking if rook on %s is checking king", rookSqr.String())

		if b.clearPath(kingSqr, rookSqr) {
			log.Debug().
				Str("Fen", b.String()).
				Str("CheckingPiece", rooks.String()).
				Str("CheckingSquare", rookSqr.String()).
				Msgf("Yes it can")

			return true
		}
	}

	return false
}

func (b *Board) knightCheck(checkingSide Side, kingSqr Square) bool {
	var knights *Knights

	switch checkingSide {
	case Black:
		knights = b.BlackKnights
	case White:
		knights = b.WhiteKnights
	}

	for _, knightMove := range KnightMoves(kingSqr) {
		for _, knightSquare := range knights.BitBoard.Squares() {
			if knightMove == knightSquare {
				return true
			}
		}
	}

	return false
}

func (b *Board) bishopCheck(checkingSide Side, kingSqr Square) bool {
	var bishops *Bishops

	switch checkingSide {
	case Black:
		bishops = b.BlackBishops
	case White:
		bishops = b.WhiteBishops
	}

	for _, bishopSqr := range bishops.BitBoard.Squares() {
		log.Debug().Msgf("Checking if bishop on %s is checking king", bishopSqr.String())

		if b.clearPath(kingSqr, bishopSqr) {
			log.Debug().
				Str("Fen", b.String()).
				Str("CheckingPiece", bishops.String()).
				Str("CheckingSquare", bishopSqr.String()).
				Msgf("Yes it can")

			return true
		}
	}

	return false
}

func (b *Board) queenCheck(checkingSide Side, kingSqr Square) bool {
	var queens *Queens

	switch checkingSide {
	case Black:
		queens = b.BlackQueens
	case White:
		queens = b.WhiteQueens
	}

	for _, queenSqr := range queens.Positions().Squares() {
		log.Debug().Msgf("Checking if queen on %s is checking king", queenSqr.String())

		if b.clearPath(kingSqr, queenSqr) {
			log.Debug().
				Str("Fen", b.String()).
				Str("CheckingPiece", queens.String()).
				Str("CheckingSquare", queenSqr.String()).
				Msgf("Yes it can")

			return true
		}
	}

	return false
}

func (b *Board) pawnCheck(checkingSide Side, kingSqr Square) bool {
	var pawns *Pawns

	var pawnCapFunc func(Square) []Square

	switch checkingSide {
	case Black:
		pawns = b.BlackPawns
		pawnCapFunc = BlackPawnCaptureMoves
	case White:
		pawns = b.WhitePawns
		pawnCapFunc = WhitePawnCaptureMoves
	}

	for _, pawnSqr := range pawns.BitBoard.Squares() {
		log.Debug().Msgf("Checking if pawn on %s is checking king", pawnSqr.String())

		for _, captureSqr := range pawnCapFunc(pawnSqr) {
			if captureSqr == kingSqr {
				return true
			}
		}
	}

	return false
}

func (b *Board) IsInCheck(side Side) bool {
	var checkedKing *King

	var checkingSide Side

	switch side {
	case Black:
		checkedKing = b.BlackKing
		checkingSide = White
	case White:
		checkedKing = b.WhiteKing
		checkingSide = Black
	}

	checkedKingSquare := checkedKing.BitBoard.Squares()[0] // only one king

	switch {
	case b.rookCheck(checkingSide, checkedKingSquare):
		log.Debug().Msg("Rooks can see Enemy King")

		return true
	case b.knightCheck(checkingSide, checkedKingSquare):
		log.Debug().Msg("Knights can see Enemy King")

		return true
	case b.bishopCheck(checkingSide, checkedKingSquare):
		log.Debug().Msg("Bishops can see Enemy King")

		return true
	case b.queenCheck(checkingSide, checkedKingSquare):
		log.Debug().Msg("Queens can see Enemy King")

		return true
	case b.pawnCheck(checkingSide, checkedKingSquare):
		log.Debug().Msg("Pawns can see Enemy King")

		return true
	default:
		log.Debug().Msg("No Pieces can see enemy king.")

		return false
	}
}

// nolint:funlen,cyclop // can't make this any smaller
func (b *Board) MakeMove(side Side, lanMove string) (*Board, *errors.MoveError) {
	parsedMove := parseMove(lanMove)

	var movingPiece Piece

	newBoard := NewBoard()
	// ignore errors we know this position is legal
	_ = newBoard.SetPieces(b.String())

	switch strings.ToUpper(parsedMove.piece) {
	case "", "P":
		switch side {
		case Black:
			movingPiece = newBoard.BlackPawns
		case White:
			movingPiece = newBoard.WhitePawns
		}
	case "R":
		switch side {
		case Black:
			movingPiece = newBoard.BlackRooks
		case White:
			movingPiece = newBoard.WhiteRooks
		}
	case "N":
		switch side {
		case Black:
			movingPiece = newBoard.BlackKnights
		case White:
			movingPiece = newBoard.WhiteKnights
		}
	case "B":
		switch side {
		case Black:
			movingPiece = newBoard.BlackBishops
		case White:
			movingPiece = newBoard.WhiteBishops
		}
	case "Q":
		switch side {
		case Black:
			movingPiece = newBoard.BlackQueens
		case White:
			movingPiece = newBoard.WhiteQueens
		}
	case "K":
		switch side {
		case Black:
			movingPiece = newBoard.BlackKing
		case White:
			movingPiece = newBoard.WhiteKing
		}
	}

	if !movingPiece.ValidMove(parsedMove.src, parsedMove.dst) {
		log.Debug().
			Interface("move", parsedMove).
			Msg("Move is invalid")

		return nil, &errors.MoveError{
			Fen:  b.String(),
			Err:  "Invalid move.",
			Move: lanMove,
		}
	}

	if newBoard.OccupiedBy(parsedMove.src) != movingPiece.String() {
		return nil, &errors.MoveError{
			Fen:  b.String(),
			Err:  "Illegal move, the src square does not contain the expected piece.",
			Move: lanMove,
		}
	}

	if newBoard.Occupied(parsedMove.dst) {
		if unicode.IsUpper(rune(newBoard.OccupiedBy(parsedMove.src)[0])) && unicode.IsUpper(rune(newBoard.OccupiedBy(parsedMove.dst)[0])) {
			return nil, &errors.MoveError{
				Fen:  b.String(),
				Err:  "Illegal move, you cannot capture your own pieces.",
				Move: lanMove,
			}
		}
	}

	for _, sqr := range SquaresBetween(parsedMove.src, parsedMove.dst) {
		if newBoard.Occupied(sqr) {
			return nil, &errors.MoveError{
				Fen:  b.String(),
				Err:  "Illegal move, there is an intervening piece.",
				Move: lanMove,
			}
		}
	}

	// actually move the piece
	movingPiece.Positions().FlipBit(parsedMove.src)
	movingPiece.Positions().FlipBit(parsedMove.dst)

	if newBoard.IsInCheck(side) {
		return nil, &errors.MoveError{
			Fen:  b.String(),
			Err:  "Illegal move, you cannot end your turn in check.",
			Move: lanMove,
		}
	}

	return nil, nil
}
