package usecase

import (
	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/src/modules/genre/model"
	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/src/modules/genre/repository"
)

// GenreUsecaseImpl ...?
type GenreUsecaseImpl struct {
	GenreRepository repository.GenreRepository
}

// NewGenreUsecase gives a new GenreUsecaseImpl
func NewGenreUsecase(GenreRepository repository.GenreRepository) *GenreUsecaseImpl {
	return &GenreUsecaseImpl{GenreRepository}
}

// SaveGenre -- todo
func (gu *GenreUsecaseImpl) SaveGenre(genre *model.Genre) (*model.Genre, error) {

	if err := gu.GenreRepository.Save(genre); err != nil {
		return nil, err
	}

	return genre, nil
}

// UpdateGenre -- todo
func (gu *GenreUsecaseImpl) UpdateGenre(id string, genre *model.Genre) (*model.Genre, error) {

	if err := gu.GenreRepository.Update(id, genre); err != nil {
		return nil, err
	}

	return genre, nil
}

// GetByID -- todo
func (gu *GenreUsecaseImpl) GetByID(id string) (*model.Genre, error) {
	var genre *model.Genre

	genre, err := gu.GenreRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return genre, nil

}
