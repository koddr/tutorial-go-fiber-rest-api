package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/tutorial-go-rest-api-fiber/app/models"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*sqlx.DB
}

// GetBooks func for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	var books []models.Book

	// Send query to database.
	if err := q.Select(&books, `SELECT * FROM books`); err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

// GetBook func for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	var book models.Book

	// Send query to database.
	if err := q.Get(&book, `SELECT * FROM books WHERE id = $1`, id); err != nil {
		return models.Book{}, err
	}

	return book, nil
}

// CreateBook func for creating book by given Book object.
func (q *BookQueries) CreateBook(u *models.Book) error {
	// Send query to database.
	if _, err := q.Exec(
		`INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Title,
		u.Author,
		u.BookStatus,
		u.BookAttrs,
	); err != nil {
		return err
	}

	return nil
}

// UpdateBook func for updating book by given Book object.
func (q *BookQueries) UpdateBook(u *models.Book) error {
	// Send query to database.
	if _, err := q.Exec(
		`UPDATE books SET updated_at = $2, title = $3, author = $4, book_attrs = $5 WHERE id = $1`,
		u.ID,
		u.UpdatedAt,
		u.Title,
		u.Author,
		u.BookAttrs,
	); err != nil {
		return err
	}

	return nil
}

// DeleteBook func for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	// Send query to database.
	if _, err := q.Exec(`DELETE FROM books WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}
