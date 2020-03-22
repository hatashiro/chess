package main

import (
	"fmt"

	"github.com/utatti/chess"
)

func clear() {
	fmt.Print("\033[2J")
}

func session(player chess.Player) uint64 {
	return uint64(player) + 1234
}

func Symbol(piece *chess.Piece) string {
	symbols := map[chess.Player](map[chess.PieceType]string){
		chess.P1: map[chess.PieceType]string{
			chess.KING:   "♚",
			chess.QUEEN:  "♛",
			chess.ROOK:   "♜",
			chess.BISHOP: "♝",
			chess.KNIGHT: "♞",
			chess.PAWN:   "♟",
		},
		chess.P2: map[chess.PieceType]string{
			chess.KING:   "♔",
			chess.QUEEN:  "♕",
			chess.ROOK:   "♖",
			chess.BISHOP: "♗",
			chess.KNIGHT: "♘",
			chess.PAWN:   "♙",
		},
	}

	return symbols[piece.Owner][piece.Type]
}

func PrintGame(game *chess.Game) {
	p1 := game.Players[chess.P1]
	p2 := game.Players[chess.P2]

	fmt.Println("  P1:", p1)
	fmt.Println("  A B C D E F G H")
	fmt.Println(" ┌─┬─┬─┬─┬─┬─┬─┬─┐")
	var row, col int8
	for row = 0; row < chess.MAX_RANK; row++ {
		fmt.Printf("%d│", row+1)
		for col = 0; col < chess.MAX_RANK; col++ {
			piece, ok := game.Board[chess.Location{row, col}]
			if ok {
				fmt.Print(Symbol(&piece))
			} else {
				fmt.Print(" ")
			}
			fmt.Print("│")
		}
		fmt.Println()

		if row < chess.MAX_RANK-1 {
			fmt.Println(" ├─┼─┼─┼─┼─┼─┼─┼─┤")
		}
	}
	fmt.Println(" └─┴─┴─┴─┴─┴─┴─┴─┘")
	fmt.Println("  P2:", p2)
	fmt.Println()
	if game.Phase == chess.ACTIVE {
		fmt.Printf("  %s's turn.\n", game.Players[game.State.Turn])
	}
}

func main() {
	game := chess.CreateGame("Chess")

	var err error

	for {
		clear()

		if err != nil {
			fmt.Println("Error:", err, "\n")
		}

		PrintGame(game)
		fmt.Println()

		if game.Phase == chess.WAITING {
			fmt.Print("P1: ")
			var name string
			fmt.Scanln(&name)
			game.Register(chess.P1, session(chess.P1), name)

			fmt.Print("P2: ")
			fmt.Scanln(&name)
			game.Register(chess.P2, session(chess.P2), name)
		} else if game.Phase == chess.ACTIVE {
			if game.State.Promotion == nil {
				// Move
				fmt.Print("Move: ")

				var fromRow, toRow int8
				var fromCol, toCol int8
				fmt.Scanf("%c%d,%c%d", &fromCol, &fromRow, &toCol, &toRow)

				from := chess.Location{fromRow - 1, fromCol - int8('A')}
				to := chess.Location{toRow - 1, toCol - int8('A')}

				err = game.Move(session(game.State.Turn), from, to)
			} else {
				// Promotion
				fmt.Println("Queen: 1, Rook: 2, Bishop: 3, Knight: 4")
				fmt.Print("Promote to: ")

				var to uint8
				fmt.Scanf("%d", to)

				err = game.Promote(session(game.State.Turn), chess.PieceType(to))
			}
		}
	}
}
