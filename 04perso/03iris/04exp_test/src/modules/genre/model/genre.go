package model

// Genre represents a genre of a book
type Genre struct {
	ID    string `bson:"_id"`
	label string `bson:"label"`
}

// Genres represent all genres for all books
type Genres []Genre
