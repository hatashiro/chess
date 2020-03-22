package chess

import (
	"errors"
	"fmt"
	"time"
)

type Game struct {
	Id string

	Sessions map[uint64]Player
	Players  map[Player]string

	State
}

func CreateGame(id string) Game {
	return Game{
		Id:       id,
		Sessions: make(map[uint64]Player),
		Players:  make(map[Player]string),
		State: State{
			Turn: P1,
			Board: Board{
				Location{0, 0}: piece(P1, ROOK),
				Location{0, 1}: piece(P1, KNIGHT),
				Location{0, 2}: piece(P1, BISHOP),
				Location{0, 3}: piece(P1, QUEEN),
				Location{0, 4}: piece(P1, KING),
				Location{0, 5}: piece(P1, BISHOP),
				Location{0, 6}: piece(P1, KNIGHT),
				Location{0, 7}: piece(P1, ROOK),
				Location{1, 0}: piece(P1, PAWN),
				Location{1, 1}: piece(P1, PAWN),
				Location{1, 2}: piece(P1, PAWN),
				Location{1, 3}: piece(P1, PAWN),
				Location{1, 4}: piece(P1, PAWN),
				Location{1, 5}: piece(P1, PAWN),
				Location{1, 6}: piece(P1, PAWN),
				Location{1, 7}: piece(P1, PAWN),
				Location{6, 0}: piece(P2, PAWN),
				Location{6, 1}: piece(P2, PAWN),
				Location{6, 2}: piece(P2, PAWN),
				Location{6, 3}: piece(P2, PAWN),
				Location{6, 4}: piece(P2, PAWN),
				Location{6, 5}: piece(P2, PAWN),
				Location{6, 6}: piece(P2, PAWN),
				Location{6, 7}: piece(P2, PAWN),
				Location{7, 0}: piece(P2, ROOK),
				Location{7, 1}: piece(P2, KNIGHT),
				Location{7, 2}: piece(P2, BISHOP),
				Location{7, 3}: piece(P2, QUEEN),
				Location{7, 4}: piece(P2, KING),
				Location{7, 5}: piece(P2, BISHOP),
				Location{7, 6}: piece(P2, KNIGHT),
				Location{7, 7}: piece(P2, ROOK),
			},
			LastUpdated: time.Now().Unix(),
		},
	}
}

// Helper functions to create an initial piece.
func piece(owner Player, pieceType PieceType) Piece {
	return Piece{
		Owner: owner,
		Type:  pieceType,
		Moved: false,
	}
}

func (game *Game) Register(player Player, session uint64, name string) error {
	_, ok := game.Players[player]
	if ok {
		return errors.New("The player is already registered")
	}

	_, ok = game.Sessions[session]
	if ok {
		return errors.New("You are already registered")
	}

	game.Sessions[session] = player
	game.Players[player] = name

	return nil
}

func (game *Game) Unregister(session uint64) error {
	player, ok := game.Sessions[session]
	if !ok {
		return errors.New("You are not registered")
	}

	delete(game.Players, player)
	delete(game.Sessions, session)

	return nil
}

func (game *Game) Move(session uint64, from Location, to Location) error {
	player, ok := game.Sessions[session]
	if !ok {
		return errors.New("You are not registered")
	}

	return Move(player, &game.State, from, to)
}

func (game *Game) Print() {
	p1 := game.Players[P1]
	p2 := game.Players[P2]

	fmt.Println("  P1:", p1)
	fmt.Println("  1 2 3 4 5 6 7 8")
	fmt.Println(" ┌─┬─┬─┬─┬─┬─┬─┬─┐")
	var row, col int8
	for row = 0; row < LINE_SIZE; row++ {
		fmt.Printf("%c│", int8('A')+row)
		for col = 0; col < LINE_SIZE; col++ {
			piece, ok := game.Board[Location{row, col}]
			if ok {
				fmt.Print(piece.Symbol())
			} else {
				fmt.Print(" ")
			}
			fmt.Print("│")
		}
		fmt.Println()

		if row < LINE_SIZE-1 {
			fmt.Println(" ├─┼─┼─┼─┼─┼─┼─┼─┤")
		}
	}
	fmt.Println(" └─┴─┴─┴─┴─┴─┴─┴─┘")
	fmt.Println("  P2:", p2)
	fmt.Println()
	fmt.Printf("  %s's turn.\n", game.Players[game.State.Turn])
}
