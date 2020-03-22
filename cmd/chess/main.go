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
				fmt.Print(piece.Symbol())
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
	fmt.Printf("  %s's turn.\n", game.Players[game.State.Turn])
}

func main() {
	game := chess.CreateGame("Chess")

	game.Register(chess.P1, session(chess.P1), "Yui")
	game.Register(chess.P2, session(chess.P2), "Mio")

	var err error

	for {
		clear()

		if err != nil {
			fmt.Println("Error:", err, "\n")
		}

		PrintGame(&game)

		fmt.Println()
		fmt.Print("Move: ")

		var fromRow, toRow int8
		var fromCol, toCol int8
		fmt.Scanf("%c%d,%c%d", &fromCol, &fromRow, &toCol, &toRow)

		from := chess.Location{fromRow - 1, fromCol - int8('A')}
		to := chess.Location{toRow - 1, toCol - int8('A')}

		err = game.Move(session(game.State.Turn), from, to)
	}
}
