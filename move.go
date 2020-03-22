package chess

func filter(locations []Location, fn func(Location) bool) []Location {
	var result []Location
	for _, location := range locations {
		if fn(location) {
			result = append(result, location)
		}
	}
	return result
}

func MovableLocationsFromKing(board Board, from Location) []Location {
	piece := board[from]
	owner := piece.Owner

	result := []Location{
		from.Relative(-1, -1),
		from.Relative(-1, 0),
		from.Relative(-1, +1),
		from.Relative(0, -1),
		from.Relative(0, +1),
		from.Relative(+1, -1),
		from.Relative(+1, 0),
		from.Relative(+1, +1),
	}

	result = filter(result, func(loc Location) bool {
		if !loc.IsValid() {
			return false
		}
		other, ok := board[loc]
		return !ok || !other.IsOwnedBy(owner)
	})

	// Castling
	if !piece.Moved {
		// Left
		rook, ok := board[owner.RankedLocation(1, 0)]
		if ok && !rook.Moved && rook.IsOwnedBy(owner) {
			possible := true
			var col int8
			for col = 1; col < from.Col; col++ {
				if _, ok := board[owner.RankedLocation(1, col)]; ok {
					possible = false
					break
				}
			}

			if possible {
				result = append(result, from.Relative(0, -2))
			}
		}
		// Right
		rook, ok = board[owner.RankedLocation(1, 7)]
		if ok && !rook.Moved && rook.IsOwnedBy(owner) {
			possible := true
			var col int8
			for col = 6; col > from.Col; col-- {
				if _, ok := board[owner.RankedLocation(1, col)]; ok {
					possible = false
					break
				}
			}

			if possible {
				result = append(result, from.Relative(0, +2))
			}
		}
	}

	return result
}

func MovableLocationsFromQueen(board Board, from Location) []Location {
	return append(MovableLocationsFromRook(board, from),
		MovableLocationsFromBishop(board, from)...)
}

func appendUntilMeet(
	result []Location,
	owner Player,
	board Board,
	loc Location,
	relRow int8,
	relCol int8,
) []Location {
	newLoc := loc.Relative(relRow, relCol)
	if !newLoc.IsValid() {
		return result
	}

	piece, ok := board[newLoc]
	if ok {
		if piece.IsOwnedBy(owner) {
			return result
		} else {
			return append(result, newLoc)
		}
	} else {
		return appendUntilMeet(
			append(result, newLoc),
			owner, board, newLoc, relRow, relCol,
		)
	}
}

func MovableLocationsFromRook(board Board, from Location) []Location {
	piece := board[from]

	result := []Location{}
	result = appendUntilMeet(result, piece.Owner, board, from, -1, 0)
	result = appendUntilMeet(result, piece.Owner, board, from, +1, 0)
	result = appendUntilMeet(result, piece.Owner, board, from, 0, -1)
	result = appendUntilMeet(result, piece.Owner, board, from, 0, +1)
	return result
}

func MovableLocationsFromBishop(board Board, from Location) []Location {
	piece := board[from]

	result := []Location{}
	result = appendUntilMeet(result, piece.Owner, board, from, -1, -1)
	result = appendUntilMeet(result, piece.Owner, board, from, -1, +1)
	result = appendUntilMeet(result, piece.Owner, board, from, +1, -1)
	result = appendUntilMeet(result, piece.Owner, board, from, +1, +1)
	return result
}

func MovableLocationsFromKnight(board Board, from Location) []Location {
	piece := board[from]

	locations := []Location{
		from.Relative(+1, +2),
		from.Relative(+1, -2),
		from.Relative(-1, +2),
		from.Relative(-1, -2),
		from.Relative(+2, +1),
		from.Relative(+2, -1),
		from.Relative(-2, +1),
		from.Relative(-2, -1),
	}

	return filter(locations, func(loc Location) bool {
		if !loc.IsValid() {
			return false
		}
		other, ok := board[loc]
		return !ok || !other.IsOwnedBy(piece.Owner)
	})
}

func MovableLocationsFromPawn(board Board, from Location) []Location {
	piece := board[from]

	result := []Location{}

	// P1 = +1, P2 = -1
	movableLoc := from.Relative(int8(piece.Owner), 0)
	if movableLoc.IsValid() {
		if _, ok := board[movableLoc]; !ok {
			result = append(result, movableLoc)
		}
	}
	if !piece.Moved {
		movableLoc = movableLoc.Relative(int8(piece.Owner), 0)
		if movableLoc.IsValid() {
			if _, ok := board[movableLoc]; !ok {
				result = append(result, movableLoc)
			}
		}
	}

	attackableLocs := []Location{
		movableLoc.Relative(0, -1),
		movableLoc.Relative(0, +1),
	}

	result = append(
		result,
		filter(attackableLocs, func(loc Location) bool {
			if !loc.IsValid() {
				return false
			}
			other, ok := board[loc]
			return ok && !other.IsOwnedBy(piece.Owner)
		})...,
	)

	return result
}
