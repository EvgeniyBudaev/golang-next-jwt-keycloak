package entities

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Name      string
	Price     float32
}
