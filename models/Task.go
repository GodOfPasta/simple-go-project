package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Task struct {
	HashKey     uuid.UUID   `json:"hash_key"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Created     time.Time   `json:"created"`
	Updated     pq.NullTime `json:"updated"`
	Deadline    time.Time   `json:"deadline"`
	Closed      bool        `json:"closed"`
}
