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

	locations := []Location{
		Location{from.Row - 1, from.Col - 1},
		Location{from.Row - 1, from.Col},
		Location{from.Row - 1, from.Col + 1},
		Location{from.Row, from.Col - 1},
		Location{from.Row, from.Col + 1},
		Location{from.Row + 1, from.Col - 1},
		Location{from.Row + 1, from.Col},
		Location{from.Row + 1, from.Col + 1},
	}

	return filter(locations, func(location Location) bool {
		other, ok := board[location]
		return !ok || other.IsOwnedBy(piece.Owner)
	})
}

func MovableLocationsFromQueen(board Board, from Location) []Location {
	return append(MovableLocationsFromRook(board, from),
		MovableLocationsFromBishop(board, from)...)
}

func MovableLocationsFromRook(board Board, from Location) []Location {
	// piece := board[from]
	return []Location{}
}

func MovableLocationsFromBishop(board Board, from Location) []Location {
	// piece := board[from]
	return []Location{}
}

func MovableLocationsFromKnight(board Board, from Location) []Location {
	piece := board[from]

	locations := []Location{
		Location{from.Row + 1, from.Col + 2},
		Location{from.Row + 1, from.Col - 2},
		Location{from.Row - 1, from.Col + 2},
		Location{from.Row - 1, from.Col - 2},
		Location{from.Row + 2, from.Col + 1},
		Location{from.Row + 2, from.Col - 1},
		Location{from.Row - 2, from.Col + 1},
		Location{from.Row - 2, from.Col - 1},
	}

	return filter(locations, func(location Location) bool {
		other, ok := board[location]
		return !ok || other.IsOwnedBy(piece.Owner)
	})
}

func MovableLocationsFromPawn(board Board, from Location) []Location {
	// piece := board[from]
	return []Location{}
}
