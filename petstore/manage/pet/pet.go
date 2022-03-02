// This package exposes functions to manage
// pets for petstore loneliness 2000.

package pet

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Pet struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Animal    string    `json:"animal" db:"animal"`
	Price     int       `json:"price" db:"price"`
	Age       int       `json:"age" db:"age"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Pets []Pet

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
