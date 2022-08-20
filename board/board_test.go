package board_test

import (
	"testing"

	"github.com/peteches/ChessEngine/board"
	"github.com/peteches/ChessEngine/errors"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:gochecknoglobals // this is for testing purposes
var validPiecePositions = map[string]board.Board{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR": {
		WhiteKing:    board.NewKing(board.White, board.E1),
		WhiteQueens:  board.NewQueens(board.White, board.D1),
		WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
		WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
		WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
		WhitePawns: board.NewPawns(board.White,
			board.A2, board.B2, board.C2, board.D2, board.E2, board.F2, board.G2, board.H2),
		BlackKing:    board.NewKing(board.Black, board.E8),
		BlackQueens:  board.NewQueens(board.Black, board.D8),
		BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
		BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
		BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
		BlackPawns: board.NewPawns(board.Black,
			board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
	},
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR": {
		WhiteKing:    board.NewKing(board.White, board.E1),
		WhiteQueens:  board.NewQueens(board.White, board.D1),
		WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
		WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
		WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
		WhitePawns: board.NewPawns(board.White,
			board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
		BlackKing:    board.NewKing(board.Black, board.E8),
		BlackQueens:  board.NewQueens(board.Black, board.D8),
		BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
		BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
		BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
		BlackPawns: board.NewPawns(board.Black,
			board.A7, board.B7, board.C7, board.D7, board.E7, board.F7, board.G7, board.H7),
	},
	"rnbqkbnr/ppp1pppp/3p4/8/4P3/8/PPPP1PPP/RNBQKBNR": {
		WhiteKing:    board.NewKing(board.White, board.E1),
		WhiteQueens:  board.NewQueens(board.White, board.D1),
		WhiteBishops: board.NewBishops(board.White, board.C1, board.F1),
		WhiteKnights: board.NewKnights(board.White, board.B1, board.G1),
		WhiteRooks:   board.NewRooks(board.White, board.A1, board.H1),
		WhitePawns: board.NewPawns(board.White,
			board.A2, board.B2, board.C2, board.D2, board.E4, board.F2, board.G2, board.H2),
		BlackKing:    board.NewKing(board.Black, board.E8),
		BlackQueens:  board.NewQueens(board.Black, board.D8),
		BlackBishops: board.NewBishops(board.Black, board.C8, board.F8),
		BlackKnights: board.NewKnights(board.Black, board.B8, board.G8),
		BlackRooks:   board.NewRooks(board.Black, board.A8, board.H8),
		BlackPawns: board.NewPawns(board.Black,
			board.A7, board.B7, board.C7, board.D6, board.E7, board.F7, board.G7, board.H7),
	},
}

//nolint:funlen // convey testing is verbose
func TestConstants(t *testing.T) {
	Convey("Constants Set Correctly", t, func() {
		Convey("Board Squares", func() {
			So(board.A1, ShouldEqual, board.Square(1))
			So(board.B1, ShouldEqual, board.Square(1<<1))
			So(board.C1, ShouldEqual, board.Square(1<<2))
			So(board.D1, ShouldEqual, board.Square(1<<3))
			So(board.E1, ShouldEqual, board.Square(1<<4))
			So(board.F1, ShouldEqual, board.Square(1<<5))
			So(board.G1, ShouldEqual, board.Square(1<<6))
			So(board.H1, ShouldEqual, board.Square(1<<7))
			So(board.A2, ShouldEqual, board.Square(1<<8))
			So(board.B2, ShouldEqual, board.Square(1<<9))
			So(board.C2, ShouldEqual, board.Square(1<<10))
			So(board.D2, ShouldEqual, board.Square(1<<11))
			So(board.E2, ShouldEqual, board.Square(1<<12))
			So(board.F2, ShouldEqual, board.Square(1<<13))
			So(board.G2, ShouldEqual, board.Square(1<<14))
			So(board.H2, ShouldEqual, board.Square(1<<15))
			So(board.A3, ShouldEqual, board.Square(1<<16))
			So(board.B3, ShouldEqual, board.Square(1<<17))
			So(board.C3, ShouldEqual, board.Square(1<<18))
			So(board.D3, ShouldEqual, board.Square(1<<19))
			So(board.E3, ShouldEqual, board.Square(1<<20))
			So(board.F3, ShouldEqual, board.Square(1<<21))
			So(board.G3, ShouldEqual, board.Square(1<<22))
			So(board.H3, ShouldEqual, board.Square(1<<23))
			So(board.A4, ShouldEqual, board.Square(1<<24))
			So(board.B4, ShouldEqual, board.Square(1<<25))
			So(board.C4, ShouldEqual, board.Square(1<<26))
			So(board.D4, ShouldEqual, board.Square(1<<27))
			So(board.E4, ShouldEqual, board.Square(1<<28))
			So(board.F4, ShouldEqual, board.Square(1<<29))
			So(board.G4, ShouldEqual, board.Square(1<<30))
			So(board.H4, ShouldEqual, board.Square(1<<31))
			So(board.A5, ShouldEqual, board.Square(1<<32))
			So(board.B5, ShouldEqual, board.Square(1<<33))
			So(board.C5, ShouldEqual, board.Square(1<<34))
			So(board.D5, ShouldEqual, board.Square(1<<35))
			So(board.E5, ShouldEqual, board.Square(1<<36))
			So(board.F5, ShouldEqual, board.Square(1<<37))
			So(board.G5, ShouldEqual, board.Square(1<<38))
			So(board.H5, ShouldEqual, board.Square(1<<39))
			So(board.A6, ShouldEqual, board.Square(1<<40))
			So(board.B6, ShouldEqual, board.Square(1<<41))
			So(board.C6, ShouldEqual, board.Square(1<<42))
			So(board.D6, ShouldEqual, board.Square(1<<43))
			So(board.E6, ShouldEqual, board.Square(1<<44))
			So(board.F6, ShouldEqual, board.Square(1<<45))
			So(board.G6, ShouldEqual, board.Square(1<<46))
			So(board.H6, ShouldEqual, board.Square(1<<47))
			So(board.A7, ShouldEqual, board.Square(1<<48))
			So(board.B7, ShouldEqual, board.Square(1<<49))
			So(board.C7, ShouldEqual, board.Square(1<<50))
			So(board.D7, ShouldEqual, board.Square(1<<51))
			So(board.E7, ShouldEqual, board.Square(1<<52))
			So(board.F7, ShouldEqual, board.Square(1<<53))
			So(board.G7, ShouldEqual, board.Square(1<<54))
			So(board.H7, ShouldEqual, board.Square(1<<55))
			So(board.A8, ShouldEqual, board.Square(1<<56))
			So(board.B8, ShouldEqual, board.Square(1<<57))
			So(board.C8, ShouldEqual, board.Square(1<<58))
			So(board.D8, ShouldEqual, board.Square(1<<59))
			So(board.E8, ShouldEqual, board.Square(1<<60))
			So(board.F8, ShouldEqual, board.Square(1<<61))
			So(board.G8, ShouldEqual, board.Square(1<<62))
			So(board.H8, ShouldEqual, board.Square(1<<63))
		})
		Convey("Side Types should be set", func() {
			var testSide board.Side
			So(board.White, ShouldHaveSameTypeAs, testSide)
			So(board.Black, ShouldHaveSameTypeAs, testSide)

			So(board.White, ShouldEqual, 0)
			So(board.Black, ShouldEqual, 1)
		})
	})
}

func TestBoardMatricies(t *testing.T) {
	Convey("Given a BoardMatrixStoI func", t, func() {
		Convey("It should map string co-ordinates to Const Squares", func() {
			testCases := map[string]board.Square{
				"A8": board.A8,
			}
			for sqr, expectedSquare := range testCases {
				resultSquare := board.BoardMatrixStoI[sqr]
				So(resultSquare, ShouldEqual, expectedSquare)
			}
		})
	})
	Convey("Given a BoardMatrixItoS func", t, func() {
		Convey("It should map Const Squares to string co-ordinates", func() {
			testCases := map[board.Square]string{
				board.A8: "A8",
			}
			for sqr, expectedSquare := range testCases {
				resultSquare := board.BoardMatrixItoS[sqr]
				So(resultSquare, ShouldEqual, expectedSquare)
			}
		})
	})
}

//nolint:funlen // convey testing is verbose
func TestBitboard(t *testing.T) {
	Convey("Given a board.NewBitboard function", t, func() {
		Convey("With no arguments", func() {
			Convey("It should return an empty board", func() {
				bb := board.NewBitboard()
				So(*bb, ShouldResemble, board.BitBoard{})
				So(bb.Board, ShouldEqual, 0)
			})

			Convey("With args", func() {
				Convey("It should return an initialised board", func() {
					bitBoard := board.NewBitboard(board.A8)
					So(bitBoard.Board, ShouldEqual, board.A8)
					bitBoard = board.NewBitboard(board.E3)
					So(bitBoard.Board, ShouldEqual, board.E3)
					bitBoard = board.NewBitboard(board.A2, board.B2)
					So(bitBoard.Board, ShouldEqual, board.A2+board.B2)
					bitBoard = board.NewBitboard(board.A8, board.H1)
					So(bitBoard.Board, ShouldEqual, board.A8+board.H1)
				})
			})
		})
	})
	Convey("Given an existing BitBoard", t, func() {
		bitboard := board.NewBitboard()
		Convey("Bit manipulation basics", func() {
			So(0^(1<<0), ShouldEqual, 1)
		})
		Convey("When FlipBit method called", func() {
			Convey("It should update its board attribute", func() {
				So(bitboard.Board, ShouldEqual, 0)
				for _, sqr := range board.AllSquares {
					bb := board.NewBitboard()
					So(bb.Board, ShouldEqual, 0)
					bb.FlipBit(sqr)
					So(bb.Board, ShouldEqual, sqr)
				}
			})
		})
		Convey("When board.Squares() method called returns []Square where bits are 1", func() {
			So(bitboard.Squares(), ShouldResemble, []board.Square{})
			bitboard.FlipBit(board.A8)
			So(bitboard.Squares(), ShouldHaveLength, 1)
			So(bitboard.Squares(), ShouldContain, board.A8)
			bitboard.FlipBit(board.H3)
			bitboard.FlipBit(board.H4)
			So(bitboard.Squares(), ShouldHaveLength, 3)
			So(bitboard.Squares(), ShouldContain, board.A8)
			So(bitboard.Squares(), ShouldContain, board.H4)
			So(bitboard.Squares(), ShouldContain, board.H3)
		})

		Convey("When Occupied() method called with square, returns true if square occupied", func() {
			for _, sqr := range board.AllSquares {
				So(bitboard.Occupied(sqr), ShouldEqual, false)
			}
			bitboard.FlipBit(board.B4)
			So(bitboard.Occupied(board.B4), ShouldEqual, true)
			bitboard.FlipBit(board.A8)
			So(bitboard.Occupied(board.B4), ShouldEqual, true)
			So(bitboard.Occupied(board.A8), ShouldEqual, true)
			So(bitboard.Occupied(board.H7), ShouldEqual, false)
			bitboard.FlipBit(board.H7)
			So(bitboard.Occupied(board.H7), ShouldEqual, true)
		})
	})
}

func TestBoard(t *testing.T) {
	Convey("Given a Board struct", t, func() {
		testBoard := board.NewBoard()
		Convey("There should be a NewBoard() function", func() {
			So(*testBoard, ShouldHaveSameTypeAs, board.Board{})
			So(testBoard.WhiteKing, ShouldHaveSameTypeAs, &board.King{})
			So(testBoard.WhiteQueens, ShouldHaveSameTypeAs, &board.Queens{})
			So(testBoard.WhiteBishops, ShouldHaveSameTypeAs, &board.Bishops{})
			So(testBoard.WhiteKnights, ShouldHaveSameTypeAs, &board.Knights{})
			So(testBoard.WhiteRooks, ShouldHaveSameTypeAs, &board.Rooks{})
			So(testBoard.WhitePawns, ShouldHaveSameTypeAs, &board.Pawns{})
			So(testBoard.BlackKing, ShouldHaveSameTypeAs, &board.King{})
			So(testBoard.BlackQueens, ShouldHaveSameTypeAs, &board.Queens{})
			So(testBoard.BlackBishops, ShouldHaveSameTypeAs, &board.Bishops{})
			So(testBoard.BlackKnights, ShouldHaveSameTypeAs, &board.Knights{})
			So(testBoard.BlackRooks, ShouldHaveSameTypeAs, &board.Rooks{})
			So(testBoard.BlackPawns, ShouldHaveSameTypeAs, &board.Pawns{})
			So(testBoard.WhiteKing, ShouldNotBeNil)
			So(testBoard.WhiteQueens, ShouldNotBeNil)
			So(testBoard.WhiteBishops, ShouldNotBeNil)
			So(testBoard.WhiteKnights, ShouldNotBeNil)
			So(testBoard.WhiteRooks, ShouldNotBeNil)
			So(testBoard.WhitePawns, ShouldNotBeNil)
			So(testBoard.BlackKing, ShouldNotBeNil)
			So(testBoard.BlackQueens, ShouldNotBeNil)
			So(testBoard.BlackBishops, ShouldNotBeNil)
			So(testBoard.BlackKnights, ShouldNotBeNil)
			So(testBoard.BlackRooks, ShouldNotBeNil)
			So(testBoard.BlackPawns, ShouldNotBeNil)
		})
		err := testBoard.SetPieces("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
		So(err, ShouldBeNil)
		Convey("It should have an OccupiedBy() Method", func() {
			So(testBoard.OccupiedBy(board.C8), ShouldEqual, "b")
			So(testBoard.OccupiedBy(board.F8), ShouldEqual, "b")
			So(testBoard.OccupiedBy(board.D8), ShouldEqual, "q")
			So(testBoard.OccupiedBy(board.H8), ShouldEqual, "r")
			So(testBoard.OccupiedBy(board.C1), ShouldEqual, "B")
			So(testBoard.OccupiedBy(board.F1), ShouldEqual, "B")
			So(testBoard.OccupiedBy(board.D1), ShouldEqual, "Q")
			So(testBoard.OccupiedBy(board.H1), ShouldEqual, "R")
		})
		Convey("the Occupied() method should return true if any piece is in square", func() {
			for _, sqr := range board.AllSquares {
				switch sqr.Rank() {
				case board.FirstRank, board.SecondRank, board.SeventhRank, board.EighthRank:
					So(testBoard.Occupied(sqr), ShouldEqual, true)
				default:
					So(testBoard.Occupied(sqr), ShouldEqual, false)
				}
			}
		})
		Convey("the String() method returns the string representation of the pieces ala fen notation", func() {
			So(testBoard.String(), ShouldEqual, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
		})
		Convey("The setPieces method should update Pieces bitboards", func() {
			for piecePositions, expectedPiecePosition := range validPiecePositions {
				err := testBoard.SetPieces(piecePositions)
				So(err, ShouldEqual, nil)
				So(*testBoard, ShouldResemble, expectedPiecePosition)
			}
		})
		Convey("The IsInCheck() method should", func() {
			Convey("return true if Side is in check", func() {
				inCheck := []string{
					"k7/8/8/8/8/8/8/R1K5",
					"k7/2N5/8/8/8/8/8/K",
					"k7/8/8/8/8/8/8/K6B",
					"k7/8/8/8/8/8/8/K6B",
					"k7/8/8/8/8/8/8/Q1K5",
					"k7/8/8/8/8/8/8/K6Q",
					"k7/1P6/8/8/8/8/8/K7",
				}
				for _, tc := range inCheck {
					err := testBoard.SetPieces(tc)
					So(err, ShouldBeNil)
					So(testBoard.IsInCheck(board.Black), ShouldBeTrue)
				}
			})
			Convey("return false if Side is not in check", func() {
				notInCheck := []string{
					"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
				}
				for _, tc := range notInCheck {
					err := testBoard.SetPieces(tc)
					So(err, ShouldBeNil)
					So(testBoard.IsInCheck(board.Black), ShouldBeFalse)
				}
			})
		})
		Convey("The MakeMove() method should", func() {
			Convey("return an error if given an invalid move", func() {
				testCases := []struct {
					move string
					side board.Side
					err  *errors.MoveError
				}{
					{
						"e2-e5",
						board.Black,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "e2-e5",
						},
					},
					{
						"e2-e5",
						board.White,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "e2-e5",
						},
					},
					{
						"qe2-h6",
						board.Black,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "qe2-h6",
						},
					},
					{
						"re2-b5",
						board.Black,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "re2-b5",
						},
					},
					{
						"be2-e5",
						board.Black,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "be2-e5",
						},
					},
					{
						"ne2-e5",
						board.Black,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "ne2-e5",
						},
					},
					{
						"ke2-e5",
						board.Black,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Invalid move.",
							Move: "ke2-e5",
						},
					},
				}
				for _, tc := range testCases {
					_, err := testBoard.MakeMove(tc.side, tc.move)
					So(err, ShouldResemble, tc.err)
				}
			})
			Convey("return a MoveError if the move is Valid but Illegal", func() {
				testCases := []struct {
					position string
					move     string
					side     board.Side
					err      *errors.MoveError
				}{
					{
						"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
						"bC1-A3",
						board.White,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Illegal move, there is an intervening piece.",
							Move: "bC1-A3",
						},
					},
					{
						"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
						"qC1-A3",
						board.White,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Illegal move, the src square does not contain the expected piece.",
							Move: "qC1-A3",
						},
					},
					{
						"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
						"kE1-D1",
						board.White,
						&errors.MoveError{
							Fen:  testBoard.String(),
							Err:  "Illegal move, you cannot capture your own pieces.",
							Move: "kE1-D1",
						},
					},
					{
						"rnbqkbnr/ppppRppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
						"nG8-H6",
						board.Black,
						&errors.MoveError{
							Fen:  "rnbqkbnr/ppppRppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
							Err:  "Illegal move, you cannot end your turn in check.",
							Move: "nG8-H6",
						},
					},
				}
				for _, tc := range testCases {
					errPos := testBoard.SetPieces(tc.position)
					So(errPos, ShouldBeNil)
					_, err := testBoard.MakeMove(tc.side, tc.move)
					So(err, ShouldResemble, tc.err)
				}
			})

			SkipConvey("update the board with the relevant move", func() {
				testCases := map[string]string{
					"e2-e4": "rnbqkbnr/pppppppp/8/8/8/4P3/PPPP1PPP/RNBQKBNR",
				}

				for move, resultFen := range testCases {
					So(err, ShouldBeNil)
					_, err := testBoard.MakeMove(board.White, move)
					So(err, ShouldBeNil)
					So(testBoard.String(), ShouldEqual, resultFen)
				}
			})
		})
	})
}
