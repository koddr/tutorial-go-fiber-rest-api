package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Book struct describe book object.
type Book struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Title      string    `db:"title" json:"title" validate:"required"`
	Author     string    `db:"author" json:"author" validate:"required"`
	BookStatus int       `db:"book_status" json:"book_status"`
	BookAttrs  BookAttrs `db:"book_attrs" json:"book_attrs"`
}

// BookAttrs struct describe book attributes.
type BookAttrs struct {
	Picture     string `json:"picture"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
}

// Value make the BookAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (u BookAttrs) Value() (driver.Value, error) {
	return json.Marshal(u)
}

// Scan make the BookAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (u *BookAttrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &u)
}
