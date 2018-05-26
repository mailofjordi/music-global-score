package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Score struct {
	Id  string `json:"id" xml:"id" form:"id" query:"id"`
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	GlobalScore GlobalScore
}

type GlobalScore struct {
	GlobalScoreProviders GlobalScoreProvider
}

type GlobalScoreProvider struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
}

func main() {
	e := echo.New();

	e.File("/doc", "../../api/oas.yml")
	e.GET("/score/:trackId", getScore)

	e.Logger.Fatal(e.Start(":1323"))
}

func getScore(context echo.Context) error {
	score := Score{
		context.Param("trackId"),
		context.Param("trackId"),
		GlobalScore{
			GlobalScoreProvider{"Spotify"},
		},
	}
	return context.JSON(http.StatusOK,score)
}
