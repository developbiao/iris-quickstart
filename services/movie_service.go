package services

import (
	"fmt"
	"iris-quickstart/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	movieName := "2020年9月最新出来上映的电影为: " + m.repo.GetMovieName()
	fmt.Println(movieName)
	return movieName
}