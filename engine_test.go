package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/peteches/ChessEngine/board"
	. "github.com/smartystreets/goconvey/convey"
)

// nolint:funlen,gocognit,cyclop // Convey testing is verbose
func TestEngine(t *testing.T) {
	Convey("Given an Engine Function", t, func() {
		ctx := context.Background()
		Convey("It should accept a context", func() {
			engine(context.Background())
		})
		Convey("It should return three string channels two readonly one writeonly", func() {
			toEng, frmEng, debug := engine(context.Background())
			So(toEng, ShouldHaveSameTypeAs, make(chan<- string))
			So(frmEng, ShouldHaveSameTypeAs, make(<-chan string))
			So(debug, ShouldHaveSameTypeAs, make(<-chan string))
		})

		Convey("It should read from the string channel and ", func() {
			Convey("When the string read is 'uci'", func() {
				Convey("It should output identifying information to the returned channel", func() {
					ctx, ctxCancel := context.WithCancel(ctx)
					toEng, frmEng, _ := engine(ctx)
					toEng <- "uci"
					out := ""
					for x := range frmEng {
						out += x
						if x == "uciok\n" {
							break
						}
					}
					ctxCancel()
					So(out, ShouldEqual, uciOkMsg)
				})
			})
		})
		Convey("When given the printPosition command over the channel", func() {
			Convey("The engine should out put a debug line with the current position", func() {
				ctx, ctxCancel := context.WithCancel(ctx)
				toEng, _, debug := engine(ctx)
				toEng <- "printPosition"
				for x := range debug {
					So(x, ShouldEqual, "info string 8/8/8/8/8/8/8/8 w - - 0 1")

					break
				}
				ctxCancel()
			})
		})
		Convey("When given the command 'position'", func() {
			Convey("with the startpos argument", func() {
				Convey("It should initialise a new Position with pieces in their starting positions.", func() {
					ctx, ctxCancel := context.WithCancel(ctx)
					toEng, _, debug := engine(ctx)
					toEng <- "position startpos"
					toEng <- "printPosition"
					for x := range debug {
						So(x, ShouldEqual, "info string rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

						break
					}
					ctxCancel()
				})
			})
			Convey("with a fen string ", func() {
				SkipConvey("should initialise that position", func() {
					for fen := range validFenstrings {
						ctx, ctxCancel := context.WithCancel(ctx)
						toEng, _, debug := engine(ctx)
						toEng <- fmt.Sprintf("position %s", fen)
						toEng <- "printPosition"
						for x := range debug {
							So(x, ShouldEqual, fmt.Sprintf("info string %s", fen))

							break
						}
						ctxCancel()
					}
				})
				validFenstringsWithMoves := map[string]Position{
					"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": {
						Pieces: &PiecePositions{
							WhiteKing:   board.NewBitboard(board.E1),
							WhiteQueen:  board.NewBitboard(board.D1),
							WhiteBishop: board.NewBitboard(board.C1, board.F1),
							WhiteKnight: board.NewBitboard(board.B1, board.G1),
							WhiteRook:   board.NewBitboard(board.A1, board.H1),
							WhitePawn:   board.NewBitboard(board.A2, board.B2, board.C2, board.D2, board.E2, board.F2, board.G2, board.H2),
							BlackKing:   board.NewBitboard(board.E8),
							BlackQueen:  board.NewBitboard(board.D8),
							BlackBishop: board.NewBitboard(board.C8, board.F8),
							BlackKnight: board.NewBitboard(board.B8, board.G8),
							BlackRook:   board.NewBitboard(board.A8, board.H8),
							BlackPawn:   board.NewBitboard(board.A7, board.B7, board.C7, board.D4, board.E7, board.F7, board.G7, board.H7),
						},
						SideToMove: WHITE,
						CastlingRights: 0 ^ (WhiteKingSideAllowed |
							WhiteQueenSideAllowed |
							BlackKingSideAllowed |
							BlackQueenSideAllowed),
						EnPassantTarget: 0,
						HalfmoveClock:   0,
						FullMoveCounter: 1,
					},
				}
				SkipConvey("with moves Should initialise the position and make the relevant moves", func() {
					for fen, finalPosition := range validFenstringsWithMoves {
						ctx, ctxCancel := context.WithCancel(ctx)
						toEng, _, debug := engine(ctx)
						toEng <- fmt.Sprintf("position %s", fen)
						toEng <- "printPosition"
						for x := range debug {
							So(x, ShouldEqual, fmt.Sprintf("info string %s", finalPosition.String()))

							break
						}
						ctxCancel()
					}
				})
			})
			Convey("When an invalid fen is supplied an error message is returned on the frmEng channel", func() {
				for fen, errMsg := range invalidFenstrings {
					ctx, ctxCancel := context.WithCancel(ctx)
					toEng, frmEng, _ := engine(ctx)
					toEng <- fmt.Sprintf("position %s", fen)
					for x := range frmEng {
						So(x, ShouldEqual, fmt.Sprintf("info string Error setting position: %s", errMsg))

						break
					}
					ctxCancel()
				}
			})
		})
		Convey("When given an unrecognised command", func() {
			Convey("It should output notice", func() {
				ctx, ctxCancel := context.WithCancel(ctx)
				toEng, frmEng, _ := engine(ctx)
				toEng <- "Geoff"
				out := ""
				x := ""
				for x != "Received unknown CMD: Geoff\n" {
					x = <-frmEng
					out += x
				}
				ctxCancel()
				So(out, ShouldEqual, "Received unknown CMD: Geoff\n")
			})
		})
	})
}
