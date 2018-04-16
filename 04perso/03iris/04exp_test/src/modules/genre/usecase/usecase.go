package usecase

import "github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/src/modules/genre/model"

// GenreUsecase the usecases for genre
type GenreUsecase interface {
	SaveGenre(*model.Genre) (*model.Genre, error)
	UpdateGenre(string, *model.Genre) (*model.Genre, error)
	GetByID(string) (*model.Genre, error)
}
