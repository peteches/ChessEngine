package board

const TotalSquares = 64

const (
	FirstRank uint8 = iota + 1
	SecondRank
	ThirdRank
	FourthRank
	FifthRank
	SixthRank
	SeventhRank
	EighthRank
)

const (
	AFile uint8 = iota + 1
	BFile
	CFile
	DFile
	EFile
	FFile
	GFile
	HFile
)

type Square uint64

func (s *Square) File() uint8 {
	switch *s {
	case A1, A2, A3, A4, A5, A6, A7, A8:
		return AFile
	case B1, B2, B3, B4, B5, B6, B7, B8:
		return BFile
	case C1, C2, C3, C4, C5, C6, C7, C8:
		return CFile
	case D1, D2, D3, D4, D5, D6, D7, D8:
		return DFile
	case E1, E2, E3, E4, E5, E6, E7, E8:
		return EFile
	case F1, F2, F3, F4, F5, F6, F7, F8:
		return FFile
	case G1, G2, G3, G4, G5, G6, G7, G8:
		return GFile
	case H1, H2, H3, H4, H5, H6, H7, H8:
		return HFile
	default:
		return 0
	}
}

func (s *Square) Rank() uint8 {
	switch *s {
	case A1, B1, C1, D1, E1, F1, H1, G1:
		return FirstRank
	case A2, B2, C2, D2, E2, F2, H2, G2:
		return SecondRank
	case A3, B3, C3, D3, E3, F3, H3, G3:
		return ThirdRank
	case A4, B4, C4, D4, E4, F4, H4, G4:
		return FourthRank
	case A5, B5, C5, D5, E5, F5, H5, G5:
		return FifthRank
	case A6, B6, C6, D6, E6, F6, H6, G6:
		return SixthRank
	case A7, B7, C7, D7, E7, F7, H7, G7:
		return SeventhRank
	case A8, B8, C8, D8, E8, F8, H8, G8:
		return EighthRank
	default:
		return 0
	}
}

//nolint:varnamelen // these are board coordinates. longer names do not make sense
const (
	A1 Square = 1
	B1 Square = 1 << iota
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// AllSquares is a list of all squares on the board.
// nolint:gochecknoglobals // this is a pseudo const
var AllSquares = [64]Square{
	A8, B8, C8, D8, E8, F8, G8, H8,
	A7, B7, C7, D7, E7, F7, G7, H7,
	A6, B6, C6, D6, E6, F6, G6, H6,
	A5, B5, C5, D5, E5, F5, G5, H5,
	A4, B4, C4, D4, E4, F4, G4, H4,
	A3, B3, C3, D3, E3, F3, G3, H3,
	A2, B2, C2, D2, E2, F2, G2, H2,
	A1, B1, C1, D1, E1, F1, G1, H1,
}

// BoardMatrixStoI returns the Square relating to a string co-ordinate.
// nolint:gochecknoglobals // this is a pseudo const
var BoardMatrixStoI = map[string]Square{
	"A8": A8,
	"B8": B8,
	"C8": C8,
	"D8": D8,
	"E8": E8,
	"F8": F8,
	"G8": G8,
	"H8": H8,
	"A7": A7,
	"B7": B7,
	"C7": C7,
	"D7": D7,
	"E7": E7,
	"F7": F7,
	"G7": G7,
	"H7": H7,
	"A6": A6,
	"B6": B6,
	"C6": C6,
	"D6": D6,
	"E6": E6,
	"F6": F6,
	"G6": G6,
	"H6": H6,
	"A5": A5,
	"B5": B5,
	"C5": C5,
	"D5": D5,
	"E5": E5,
	"F5": F5,
	"G5": G5,
	"H5": H5,
	"A4": A4,
	"B4": B4,
	"C4": C4,
	"D4": D4,
	"E4": E4,
	"F4": F4,
	"G4": G4,
	"H4": H4,
	"A3": A3,
	"B3": B3,
	"C3": C3,
	"D3": D3,
	"E3": E3,
	"F3": F3,
	"G3": G3,
	"H3": H3,
	"A2": A2,
	"B2": B2,
	"C2": C2,
	"D2": D2,
	"E2": E2,
	"F2": F2,
	"G2": G2,
	"H2": H2,
	"A1": A1,
	"B1": B1,
	"C1": C1,
	"D1": D1,
	"E1": E1,
	"F1": F1,
	"G1": G1,
	"H1": H1,
}

// BoardMatrixItoS returns the string co-ordinate of a given Square
// nolint:gochecknoglobals // this is a pseudo const
var BoardMatrixItoS = map[Square]string{
	A8: "A8",
	B8: "B8",
	C8: "C8",
	D8: "D8",
	E8: "E8",
	F8: "F8",
	G8: "G8",
	H8: "H8",
	A7: "A7",
	B7: "B7",
	C7: "C7",
	D7: "D7",
	E7: "E7",
	F7: "F7",
	G7: "G7",
	H7: "H7",
	A6: "A6",
	B6: "B6",
	C6: "C6",
	D6: "D6",
	E6: "E6",
	F6: "F6",
	G6: "G6",
	H6: "H6",
	A5: "A5",
	B5: "B5",
	C5: "C5",
	D5: "D5",
	E5: "E5",
	F5: "F5",
	G5: "G5",
	H5: "H5",
	A4: "A4",
	B4: "B4",
	C4: "C4",
	D4: "D4",
	E4: "E4",
	F4: "F4",
	G4: "G4",
	H4: "H4",
	A3: "A3",
	B3: "B3",
	C3: "C3",
	D3: "D3",
	E3: "E3",
	F3: "F3",
	G3: "G3",
	H3: "H3",
	A2: "A2",
	B2: "B2",
	C2: "C2",
	D2: "D2",
	E2: "E2",
	F2: "F2",
	G2: "G2",
	H2: "H2",
	A1: "A1",
	B1: "B1",
	C1: "C1",
	D1: "D1",
	E1: "E1",
	F1: "F1",
	G1: "G1",
	H1: "H1",
}

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
