package chess

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

var games map[string]*Game = make(map[string]*Game)

func asset(path string) string {
	gopath := os.Getenv("GOPATH")
	return filepath.Join(gopath, "src/github.com/utatti/chess", path)
}

func NewServer() *echo.Echo {
	e := echo.New()

	e.Debug = len(os.Getenv("DEBUG")) > 0

	e.Static("/static", asset("web/static"))

	e.File("/", asset("web/index.html"))
	e.File("/:id", asset("web/game.html"))

	e.GET("/game/:id", getGame)

	e.POST("/game/:id/register", registerPlayer)
	e.POST("/game/:id/unregister", unregisterPlayer)
	e.POST("/game/:id/move", move)
	e.POST("/game/:id/promote", promote)
	e.POST("/game/:id/reset", reset)

	return e
}

func getGame(c echo.Context) error {
	gameId := c.Param("id")

	game, ok := games[gameId]
	if !ok {
		game = CreateGame(gameId)
		games[gameId] = game
	}

	return c.JSON(200, game.Json())
}

func registerPlayer(c echo.Context) error {
	gameId := c.Param("id")

	game, ok := games[gameId]
	if !ok {
		return errors.New("No game found.")
	}

	body := new(struct {
		Player  Player `json:"player"`
		Session uint64 `json:"session"`
		Name    string `json:"name"`
	})

	if err := c.Bind(body); err != nil {
		return err
	}

	if err := game.Register(body.Player, body.Session, body.Name); err != nil {
		return err
	}

	return c.String(200, "ok")
}

func unregisterPlayer(c echo.Context) error {
	gameId := c.Param("id")

	game, ok := games[gameId]
	if !ok {
		return errors.New("No game found.")
	}

	body := new(struct {
		Session uint64 `json:"session"`
	})

	if err := c.Bind(body); err != nil {
		return err
	}

	if err := game.Unregister(body.Session); err != nil {
		return err
	}

	return c.String(200, "ok")
}

func move(c echo.Context) error {
	gameId := c.Param("id")

	game, ok := games[gameId]
	if !ok {
		return errors.New("No game found.")
	}

	body := new(struct {
		Session  uint64 `json:"session"`
		FromInt8 int8   `json:"from"`
		ToInt8   int8   `json:"to"`
	})

	if err := c.Bind(body); err != nil {
		return err
	}

	from := LocationFromInt8(body.FromInt8)
	to := LocationFromInt8(body.ToInt8)

	if err := game.Move(body.Session, from, to); err != nil {
		return err
	}

	return c.String(200, "ok")
}

func promote(c echo.Context) error {
	gameId := c.Param("id")

	game, ok := games[gameId]
	if !ok {
		return errors.New("No game found.")
	}

	body := new(struct {
		Session uint64    `json:"session"`
		To      PieceType `json:"to"`
	})

	if err := c.Bind(body); err != nil {
		return err
	}

	if err := game.Promote(body.Session, body.To); err != nil {
		return err
	}

	return c.String(200, "ok")
}

func reset(c echo.Context) error {
	gameId := c.Param("id")

	game, ok := games[gameId]
	if !ok {
		return errors.New("No game found.")
	}

	if err := game.Reset(); err != nil {
		return err
	}

	return c.String(200, "ok")
}
