package main

// nolint:cyclop,gocognit // can't really be simplified
func KnightMoves(src Square) []Square {
	moves := []Square{}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > firstRank && src.File() > BFile {
		moves = append(moves, src<<6)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > firstRank && src.File() < GFile {
		moves = append(moves, src<<10)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > secondRank && src.File() > AFile {
		moves = append(moves, src<<15)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() > secondRank && src.File() < HFile {
		moves = append(moves, src<<17)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < seventhRank && src.File() > AFile {
		moves = append(moves, src>>17)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < seventhRank && src.File() < HFile {
		moves = append(moves, src>>15)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < eighthRank && src.File() > BFile {
		moves = append(moves, src>>10)
	}

	// nolint:gomnd // not sure how to name these as constants
	if src.Rank() < eighthRank && src.File() < GFile {
		moves = append(moves, src>>6)
	}

	return moves
}
