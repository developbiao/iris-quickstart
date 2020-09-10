package repositories

import "iris-quickstart/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	// mock data
	movie := &datamodels.Movie{Name: "信条"}
	return movie.Name
}
