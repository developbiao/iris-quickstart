package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"iris-quickstart/repositories"
	"iris-quickstart/services"
)

type MoveController struct {

}

func (c *MoveController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	movieResult := movieService.ShowMovieName()
	return mvc.View {
		Name: "movie/index.html",
		Data: movieResult,
	}
}
