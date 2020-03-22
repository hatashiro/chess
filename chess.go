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
	Row, Col int8
}

func (loc Location) IsValid() bool {
	return loc.Row >= 0 && loc.Row < MAX_RANK &&
		loc.Col >= 0 && loc.Col < MAX_RANK
}

func (loc Location) RelativeTo(other Location) Location {
	return Location{loc.Row - other.Row, loc.Col - other.Col}
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
	Owner Player
	Type  PieceType
	Moved bool
}

func (piece *Piece) IsOwnedBy(player Player) bool {
	return piece.Owner == player
}

func (piece *Piece) Symbol() string {
	symbols := map[Player](map[PieceType]string){
		P1: map[PieceType]string{
			KING: "♚", QUEEN: "♛", ROOK: "♜", BISHOP: "♝", KNIGHT: "♞", PAWN: "♟",
		},
		P2: map[PieceType]string{
			KING: "♔", QUEEN: "♕", ROOK: "♖", BISHOP: "♗", KNIGHT: "♘", PAWN: "♙",
		},
	}

	return symbols[piece.Owner][piece.Type]
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

	if !ok || piece.IsOwnedBy(player) {
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

	state.move(from, to)

	// Process castling.
	if piece.Type == KING && to.RelativeTo(from).Abs().Col > 1 {
		// TODO: castling
	}

	// Check promotion and flip turn only when there's no pawn waiting for it.
	if piece.Type == PAWN && to.Row == piece.Owner.RankedRow(MAX_RANK) {
		state.Promotion = &to
	}

	state.LastUpdated = time.Now().Unix()

	return nil // No error, succeed.
}

func (state *State) move(from Location, to Location) {
	piece := state.Board[from]
	delete(state.Board, from)
	piece.Moved = true
	state.Board[to] = piece
}

func (state *State) flipTurn() {
	state.Turn = -state.Turn
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

	if to == PAWN {
		return errors.New("You cannot promote a pawn to a pawn.")
	}

	state.Board[loc] = Piece{piece.Owner, to, false}

	state.Promotion = nil
	state.flipTurn()

	return nil
}
