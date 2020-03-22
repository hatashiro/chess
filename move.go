package chess

import (
	"errors"
	"time"
)

func Move(player Player, state *State, from Location, to Location) error {
	if state.Turn != player {
		return errors.New("It's not the player's turn")
	}

	board := state.Board

	fromPiece, ok := board[from]

	if !ok || fromPiece.Owner != player {
		return errors.New("There is no player's piece in the location")
	}

	// TODO: Checkmate and castling

	found := false
	for _, movable := range getMovableLocations(board, from) {
		if movable == to {
			found = true
		}
	}

	if !found {
		return errors.New("Wrong move")
	}

	// Move the piece.
	delete(board, from)
	board[to] = fromPiece
	fromPiece.Moved = true

	state.Turn = -state.Turn
	state.LastUpdated = time.Now().Unix()

	return nil // No error, succeed.
}

func getMovableLocations(board Board, from Location) []Location {
	piece := board[from]
	owner := piece.Owner

	var locations []Location

	appendIfMovable := func(row int8, col int8) {
		newLocation := Location{from.Row + row, from.Col + col}
		if newLocation.IsValid() {
			piece, ok := board[newLocation]
			if !ok || piece.Owner != owner {
				locations = append(locations, newLocation)
			}
		}
	}

	switch piece.Type {
	case KING:
		appendIfMovable(-1, -1)
		appendIfMovable(-1, 0)
		appendIfMovable(-1, +1)
		appendIfMovable(0, -1)
		appendIfMovable(0, +1)
		appendIfMovable(+1, -1)
		appendIfMovable(+1, 0)
		appendIfMovable(+1, +1)
	case QUEEN:
		var i int8
		for i = 0; i < LINE_SIZE; i++ {
			appendIfMovable(-i, -i)
			appendIfMovable(-i, 0)
			appendIfMovable(-i, +i)
			appendIfMovable(0, -i)
			appendIfMovable(0, +i)
			appendIfMovable(+i, -i)
			appendIfMovable(+i, 0)
			appendIfMovable(+i, +i)
		}
	case ROOK:
		var i int8
		for i = 0; i < LINE_SIZE; i++ {
			appendIfMovable(-i, 0)
			appendIfMovable(+i, 0)
			appendIfMovable(0, -i)
			appendIfMovable(0, +i)
		}
	case BISHOP:
		var i int8
		for i = 0; i < LINE_SIZE; i++ {
			appendIfMovable(-i, -i)
			appendIfMovable(-i, +i)
			appendIfMovable(+i, -i)
			appendIfMovable(+i, +i)
		}
	case KNIGHT:
		appendIfMovable(1, 2)
		appendIfMovable(-1, 2)
		appendIfMovable(1, -2)
		appendIfMovable(-1, -2)
		appendIfMovable(2, 1)
		appendIfMovable(-2, 1)
		appendIfMovable(2, -1)
		appendIfMovable(-2, -1)
	case PAWN:
		movable := Location{from.Row + int8(owner), from.Col}
		if movable.IsValid() {
			if _, ok := board[movable]; !ok {
				locations = append(locations, movable)
			}
			leftAttackable := Location{movable.Row, movable.Col - 1}
			if piece, ok := board[leftAttackable]; ok && piece.Owner != owner {
				locations = append(locations, leftAttackable)
			}
			rightAttackable := Location{movable.Row, movable.Col - 1}
			if piece, ok := board[rightAttackable]; ok && piece.Owner != owner {
				locations = append(locations, rightAttackable)
			}
		}
	}

	return locations
}
