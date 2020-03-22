package chess

type State struct {
	Turn Player
	Board
	LastUpdated int64
}

type Player int8

const (
	P1 Player = 1
	P2        = -1
)

const LINE_SIZE = 8

type Location struct {
	Row, Col int8
}

func (loc *Location) IsValid() bool {
	return loc.Row >= 0 && loc.Row < LINE_SIZE && loc.Col >= 0 && loc.Col < LINE_SIZE
}

type Board map[Location]Piece

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

func (piece *Piece) Symbol() string {
	switch piece.Owner {
	case P1:
		switch piece.Type {
		case KING:
			return "♚"
		case QUEEN:
			return "♛"
		case ROOK:
			return "♜"
		case BISHOP:
			return "♝"
		case KNIGHT:
			return "♞"
		case PAWN:
			return "♟"
		}
	case P2:
		switch piece.Type {
		case KING:
			return "♔"
		case QUEEN:
			return "♕"
		case ROOK:
			return "♖"
		case BISHOP:
			return "♗"
		case KNIGHT:
			return "♘"
		case PAWN:
			return "♙"
		}
	}
	return ""
}
