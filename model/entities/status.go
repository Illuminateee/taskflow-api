package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Status struct {
	ID         uuid.UUID `json:"uuid"`
	StatusName json.RawMessage `json:"status_name"` // Using RawMessage to allow for flexible status name formats
	TasksID  uuid.UUID `json:"tasks_id"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
}

func (Status) TableName() string {
	return "status"
}