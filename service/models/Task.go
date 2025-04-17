package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Task struct {
	HashKey     uuid.UUID   `json:"hash_key"`
	Name        string      `json:"name" validate:"required,min=3,max=30"`
	Description string      `json:"description" validate:"max=90"`
	Created     time.Time   `json:"created"`
	Updated     pq.NullTime `json:"updated"`
	Deadline    time.Time   `json:"deadline" validate:"deadlineValidator"`
	Closed      bool        `json:"closed" validate:"required"`
}
