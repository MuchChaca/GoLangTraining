package books

import (
	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/authors"
	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/genres"
)

// Book represents a book
type Book struct {
	// SessionID string `json:"-"`
	ID          int64          `json:"id,omitempty"`
	Title       string         `json:"title"`
	Description bool           `json:"description"`
	Author      authors.Author `json:"author"`
	Genres      []genres.Genre `json:"genres"`
}
