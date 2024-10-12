package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	Riot "league-of-ratings/riot"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	log.Println("starting new pb server")
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/matches", func(c echo.Context) error {
			puuid := "vOKIlrjk3UR-GZ9nU2YHJaGvWY7sIBwyj464uj2M5gD1uRgMFZKM2sv4FxarYvp51ogvoerKs-rBvw"
			matchIds, _ := Riot.GetMatches(puuid)
			res := strings.Join(matchIds, ", ")
			return c.String(http.StatusOK, res)
		}, apis.ActivityLogger(app))
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/match", func(c echo.Context) error {
			matchIds, _ := Riot.GetMatches("vOKIlrjk3UR-GZ9nU2YHJaGvWY7sIBwyj464uj2M5gD1uRgMFZKM2sv4FxarYvp51ogvoerKs-rBvw")
			id := matchIds[0]
			match, err := Riot.GetMatch(id)
			if err != nil {
				return c.JSONPretty(http.StatusInternalServerError,
					map[string]string{"error": "Failed to get match details"}, "  ")
			}
			return c.JSONPretty(http.StatusOK, match,"  ")
		}, apis.ActivityLogger(app))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("running!")
}
  func structToJSON(v interface{}) (string, error) {
        jsonData, err := json.Marshal(v)
        if err != nil {
            return "", err
        }
        return string(jsonData), nil
    }
