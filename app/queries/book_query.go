package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books`

	// Send query to database.
	err := q.Get(&books, query)
	if err != nil {
		return books, err
	}

	return books, nil
}

// GetBooksByAuthor method for getting all books by given author.
func (q *BookQueries) GetBooksByAuthor(author string) ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE author = $1`

	// Send query to database.
	err := q.Get(&books, query, author)
	if err != nil {
		return books, err
	}

	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	err := q.Get(&book, query, id)
	if err != nil {
		return book, err
	}

	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(u *models.Book) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(query, u.ID, u.CreatedAt, u.UpdatedAt, u.Title, u.Author, u.BookStatus, u.BookAttrs)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(u *models.Book) error {
	// Define query string.
	query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_attrs = $5 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, u.ID, u.UpdatedAt, u.Title, u.Author, u.BookAttrs)
	if err != nil {
		return err
	}

	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
