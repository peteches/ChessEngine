package board

/*
Piece is the interface that all pieces on the board must adhere to.

It should be able to return a string representation of the piece,
as would be used in a Fen String

It should be able to return potential Squares to move to.
*/
type Piece interface {
	String() string
	ValidMove(Square, Square) bool
	Positions() *BitBoard
}
