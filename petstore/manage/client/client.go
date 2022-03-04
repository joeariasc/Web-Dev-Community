package client

import (
	"time"

	"github.com/gofrs/uuid"
)

type Client struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	Gender    string    `json:"gender" db:"gender"`
	Age       int       `json:"age" db:"age"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Clients []Client
