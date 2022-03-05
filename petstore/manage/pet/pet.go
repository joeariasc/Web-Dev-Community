// This package exposes functions to manage
// pets for petstore loneliness 2000.

package pet

import (
	"time"

	"github.com/gofrs/uuid"
)

type Pet struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Animal    string    `json:"animal" db:"animal"`
	Price     float64   `json:"price" db:"price"`
	Age       int64     `json:"age" db:"age"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Pets []Pet
