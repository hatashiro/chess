package main

import (
  "fmt"

  chess "github.com/utatti/chess"
)

func clear() {
  fmt.Print("\033[2J")
}

func session(player chess.Player) uint64 {
  return uint64(player) + 1234
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

    game.Print()

    fmt.Println()
    fmt.Print("Move: ")

    var fromRow, toRow int8
    var fromCol, toCol int8
    fmt.Scanf("%c%d,%c%d", &fromRow, &fromCol, &toRow, &toCol)

    from := chess.Location{fromRow - int8('A'), fromCol - 1}
    to := chess.Location{toRow - int8('A'), toCol - 1}

    err = game.Move(session(game.State.Turn), from, to)
  }
}
