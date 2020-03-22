package chess

import (
	"errors"
	"os"

	echo "github.com/labstack/echo"
)

var games map[string]*Game = make(map[string]*Game)

func NewServer() *echo.Echo {
	e := echo.New()

	e.Debug = len(os.Getenv("DEBUG")) > 0

	e.Static("/static", "web/static")

	e.File("/", "web/index.html")
	e.File("/:id", "web/game.html")

	e.GET("/game/:id", getGame)

	e.POST("/game/:id/register", registerPlayer)
	e.POST("/game/:id/unregister", unregisterPlayer)

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
