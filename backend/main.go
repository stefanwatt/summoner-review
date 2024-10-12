package main

import (
	"log"
	"net/http"

	// Riot "league-of-ratings/riot"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	log.Println("starting new pb server")
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hallo Markus!")
		}, apis.ActivityLogger(app))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("running!")
}
