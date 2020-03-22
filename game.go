package chess

import (
	"errors"
)

type Game struct {
	Id string

	Phase

	Sessions map[uint64]Player
	Players  map[Player]string

	State
}

type Phase uint8

const (
	WAITING Phase = iota
	ACTIVE
	DONE
)

func CreateGame(id string) Game {
	return Game{
		Id:       id,
		Sessions: make(map[uint64]Player),
		Players:  make(map[Player]string),
		State:    NewState(),
	}
}

func (game *Game) Reset() error {
	if game.Phase != DONE {
		return errors.New("The game is not yet done.")
	}
	game.Phase = WAITING
	game.Sessions = make(map[uint64]Player)
	game.Players = make(map[Player]string)
	game.State = NewState()
	return nil
}

func (game *Game) Register(player Player, session uint64, name string) error {
	if game.Phase != WAITING {
		return errors.New("The game is already started.")
	}

	_, ok := game.Players[player]
	if ok {
		return errors.New("The player is already registered.")
	}

	_, ok = game.Sessions[session]
	if ok {
		return errors.New("You are already registered.")
	}

	game.Sessions[session] = player
	game.Players[player] = name

	// Check the opponent. If both are ready, start the game.
	_, ok = game.Players[-player]
	if ok {
		game.Phase = ACTIVE
	}

	return nil
}

func (game *Game) Unregister(session uint64) error {
	if game.Phase != WAITING {
		return errors.New("The game is already started.")
	}

	player, ok := game.Sessions[session]
	if !ok {
		return errors.New("You are not registered.")
	}

	delete(game.Players, player)
	delete(game.Sessions, session)

	return nil
}

func (game *Game) assertTurn(session uint64) error {
	player, ok := game.Sessions[session]
	if !ok {
		return errors.New("You are not registered.")
	}
	if game.State.Turn != player {
		return errors.New("It's not your turn.")
	}
	return nil
}

func (game *Game) Move(session uint64, from Location, to Location) error {
	if err := game.assertTurn(session); err != nil {
		return err
	}

	if err := game.State.TryMove(from, to); err != nil {
		return err
	}

	// TODO: Checkmate or win/lose

	return nil
}

func (game *Game) Promote(session uint64, to PieceType) error {
	if err := game.assertTurn(session); err != nil {
		return err
	}

	if err := game.State.TryPromote(to); err != nil {
		return err
	}

	// TODO: Checkmate or win/lose

	return nil
}
