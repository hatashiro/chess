package chess

import (
	"errors"
	"time"
)

type Player int8

const (
	P1 Player = 1
	P2        = -1
)

func (player Player) RankedRow(rank int8) int8 {
	if player == P1 {
		return rank - 1
	} else {
		return MAX_RANK - rank
	}
}

func (player Player) RankedLocation(rank int8, col int8) Location {
	return Location{player.RankedRow(rank), col}
}

const MAX_RANK = 8

type Location struct {
	Row int8 `json:"row"`
	Col int8 `json:"col"`
}

func (loc Location) IsValid() bool {
	return loc.Row >= 0 && loc.Row < MAX_RANK &&
		loc.Col >= 0 && loc.Col < MAX_RANK
}

func (loc Location) Relative(row int8, col int8) Location {
	return Location{loc.Row + row, loc.Col + col}
}

func (loc Location) RelativeTo(other Location) Location {
	return loc.Relative(-other.Row, -other.Col)
}

func (loc Location) Abs() Location {
	var row, col int8
	if loc.Row >= 0 {
		row = loc.Row
	} else {
		row = -loc.Row
	}
	if loc.Col >= 0 {
		col = loc.Col
	} else {
		col = -loc.Col
	}
	return Location{row, col}
}

func (loc Location) Int8() int8 {
	return loc.Row*MAX_RANK + loc.Col
}

func LocationFromInt8(i int8) Location {
	return Location{i / MAX_RANK, i % MAX_RANK}
}

type Board map[Location]Piece

func NewBoard() Board {
	return Board{
		Location{0, 0}: Piece{P1, ROOK, false},
		Location{0, 1}: Piece{P1, KNIGHT, false},
		Location{0, 2}: Piece{P1, BISHOP, false},
		Location{0, 3}: Piece{P1, QUEEN, false},
		Location{0, 4}: Piece{P1, KING, false},
		Location{0, 5}: Piece{P1, BISHOP, false},
		Location{0, 6}: Piece{P1, KNIGHT, false},
		Location{0, 7}: Piece{P1, ROOK, false},
		Location{1, 0}: Piece{P1, PAWN, false},
		Location{1, 1}: Piece{P1, PAWN, false},
		Location{1, 2}: Piece{P1, PAWN, false},
		Location{1, 3}: Piece{P1, PAWN, false},
		Location{1, 4}: Piece{P1, PAWN, false},
		Location{1, 5}: Piece{P1, PAWN, false},
		Location{1, 6}: Piece{P1, PAWN, false},
		Location{1, 7}: Piece{P1, PAWN, false},
		Location{6, 0}: Piece{P2, PAWN, false},
		Location{6, 1}: Piece{P2, PAWN, false},
		Location{6, 2}: Piece{P2, PAWN, false},
		Location{6, 3}: Piece{P2, PAWN, false},
		Location{6, 4}: Piece{P2, PAWN, false},
		Location{6, 5}: Piece{P2, PAWN, false},
		Location{6, 6}: Piece{P2, PAWN, false},
		Location{6, 7}: Piece{P2, PAWN, false},
		Location{7, 0}: Piece{P2, ROOK, false},
		Location{7, 1}: Piece{P2, KNIGHT, false},
		Location{7, 2}: Piece{P2, BISHOP, false},
		Location{7, 3}: Piece{P2, QUEEN, false},
		Location{7, 4}: Piece{P2, KING, false},
		Location{7, 5}: Piece{P2, BISHOP, false},
		Location{7, 6}: Piece{P2, KNIGHT, false},
		Location{7, 7}: Piece{P2, ROOK, false},
	}
}

func (board Board) Copy() Board {
	result := make(Board)
	for key, value := range board {
		result[key] = value
	}
	return result
}

func (board Board) movableLocations(from Location) []Location {
	var locations []Location

	piece, ok := board[from]
	if !ok {
		return locations
	}

	switch piece.Type {
	case KING:
		return MovableLocationsFromKing(board, from)
	case QUEEN:
		return MovableLocationsFromQueen(board, from)
	case ROOK:
		return MovableLocationsFromRook(board, from)
	case BISHOP:
		return MovableLocationsFromBishop(board, from)
	case KNIGHT:
		return MovableLocationsFromKnight(board, from)
	case PAWN:
		return MovableLocationsFromPawn(board, from)
	}

	return locations
}

func (board Board) move(from Location, to Location) {
	piece := board[from]
	delete(board, from)
	piece.Moved = true
	board[to] = piece
}

func (board Board) isCastling(from Location, to Location) bool {
	piece := board[from]
	return piece.Type == KING && to.RelativeTo(from).Abs().Col == 2
}

func (board Board) Json() interface{} {
	res := make(map[int8]Piece)
	for loc, piece := range board {
		res[loc.Int8()] = piece
	}
	return res
}

type PieceType uint8

const (
	_ PieceType = iota
	KING
	QUEEN
	ROOK
	BISHOP
	KNIGHT
	PAWN
)

type Piece struct {
	Owner Player    `json:"owner"`
	Type  PieceType `json:"type"`
	Moved bool      `json:"moved"`
}

func (piece *Piece) IsOwnedBy(player Player) bool {
	return piece.Owner == player
}

type State struct {
	Turn Player
	Board
	Promotion   *Location
	LastUpdated int64
}

func NewState() State {
	return State{
		Turn:        P1,
		Board:       NewBoard(),
		LastUpdated: time.Now().Unix(),
	}
}

func (state *State) TryMove(from Location, to Location) error {
	if state.Promotion != nil {
		return errors.New("There is a pawn waiting for promotion.")
	}

	player := state.Turn
	board := state.Board

	piece, ok := board[from]

	if !ok || !piece.IsOwnedBy(player) {
		return errors.New("There is no player's piece in the location.")
	}

	isMovable := false
	for _, movableLoc := range board.movableLocations(from) {
		if movableLoc == to {
			isMovable = true
		}
	}

	if !isMovable {
		return errors.New("It's a wrong move.")
	}

	boardBeforeMove := board.Copy()

	if board.isCastling(from, to) {
		if state.IsChecked() {
			return errors.New("You cannot castle when checked.")
		}

		board.move(from, to)
		if to.RelativeTo(from).Col > 0 {
			// Right
			board.move(piece.Owner.RankedLocation(1, 7), to.Relative(0, -1))
		} else {
			// Left
			board.move(piece.Owner.RankedLocation(1, 0), to.Relative(0, +1))
		}
	} else {
		board.move(from, to)
	}

	if state.IsChecked() {
		// Rollback.
		state.Board = boardBeforeMove
		return errors.New("It's a wrong move. You'll be checked")
	}

	// Check promotion and flip turn only when there's no pawn waiting for it.
	if piece.Type == PAWN && to.Row == piece.Owner.RankedRow(MAX_RANK) {
		state.Promotion = &to
	} else {
		state.flipTurn()
	}

	state.updateTimestamp()

	return nil // No error, succeed.
}

func (state *State) flipTurn() {
	state.Turn = -state.Turn
}

func (state *State) updateTimestamp() {
	state.LastUpdated = time.Now().Unix()
}

func (state *State) TryPromote(to PieceType) error {
	if state.Promotion == nil {
		return errors.New("There is no pawn waiting for promotion.")
	}

	loc := *state.Promotion
	piece, ok := state.Board[loc]

	if !ok || piece.Type != PAWN || !piece.IsOwnedBy(state.Turn) {
		return errors.New("It's a wrong piece for promotion.")
	}

	if to == PAWN || to == KING {
		return errors.New("You cannot promote a pawn to that.")
	}

	state.Board[loc] = Piece{piece.Owner, to, false}

	state.Promotion = nil
	state.flipTurn()
	state.updateTimestamp()

	return nil
}

func (state *State) IsChecked() bool {
	player := state.Turn
	board := state.Board

	var king Location
	var opponents []Location

	for loc, piece := range board {
		if piece.IsOwnedBy(player) {
			if piece.Type == KING {
				king = loc
			}
		} else {
			opponents = append(opponents, loc)
		}
	}

	for _, from := range opponents {
		for _, to := range board.movableLocations(from) {
			if to == king {
				return true
			}
		}
	}

	return false
}

func (state *State) IsCheckmated() bool {
	if !state.IsChecked() {
		return false
	}

	var allies []Location
	for loc, piece := range state.Board {
		if piece.IsOwnedBy(state.Turn) {
			allies = append(allies, loc)
		}
	}

	originalBoard := state.Board.Copy()
	for _, from := range allies {
		for _, to := range originalBoard.movableLocations(from) {
			if state.Board.isCastling(from, to) {
				continue
			}

			state.Board.move(from, to)
			checked := state.IsChecked()
			state.Board = originalBoard.Copy()
			if !checked {
				return false
			}
		}
	}
	return true
}

func (state *State) Json() interface{} {
	return &struct {
		Turn        Player      `json:"turn"`
		Board       interface{} `json:"board"`
		Promotion   *Location   `json:"promotion"`
		LastUpdated int64       `json:"lastUpdated"`
	}{
		state.Turn,
		state.Board.Json(),
		state.Promotion,
		state.LastUpdated,
	}
}
