package repository

import "github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/src/modules/genre/model"

// GenreRepository Required funcs to make Genres work
type GenreRepository interface {
	Save(*model.Genre) error
	Update(string, *model.Genre) error
	Delete(string) error
	FindById(string) (*model.Genre, error)
	FindAll() (*model.Genres, error)
}
