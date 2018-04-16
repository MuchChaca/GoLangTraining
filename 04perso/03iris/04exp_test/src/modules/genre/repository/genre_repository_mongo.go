package repository

import (
	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/src/modules/genre/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GenreRepositoryMongo Mongo repository representation for genres
type GenreRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

// NewGenreRepositoryMongo returns a new genreRepositoryMongo
func NewGenreRepositoryMongo(db *mgo.Database, collection string) *GenreRepositoryMongo {
	return &GenreRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

// Save saves the genre passed in param, if something goes wrong, returns the error
func (r *GenreRepositoryMongo) Save(genre *model.Genre) error {
	return r.db.C(r.collection).Insert(genre)
}

// Update updates a genre with the id passed in params, if there is an error, it gets returned
func (r *GenreRepositoryMongo) Update(id string, genre *model.Genre) error {
	return r.db.C(r.collection).UpdateId(id, genre)
	//? return r.db.C(r.collection).Update(bson.M{"id:" id}, genre)
}

// Delete deletes the genre with the matching id
func (r *GenreRepositoryMongo) Delete(id string) error {
	return r.db.C(r.collection).RemoveId("id")
	//? return r.db.C(r.collection).Remove(bson.M{"id": id})
}

// FindByID fetches the genre with the corresponding id
func (r *GenreRepositoryMongo) FindByID(id string) (*model.Genre, error) {
	var genre model.Genre

	//? if err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&genre); err != nil {
	if err := r.db.C(r.collection).FindId(id).One(&genre); err != nil {
		return nil, err
	}
	return &genre, nil
}

// FindAll returns all genres
func (r *GenreRepositoryMongo) FindAll() (*model.Genres, error) {
	var genres model.Genres

	if err := r.db.C(r.collection).Find(bson.M{}).All(&genres); err != nil {
		return nil, err
	}
	return &genres, nil
}
