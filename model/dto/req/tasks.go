package req

import (
	"time"

	"github.com/google/uuid"
)

type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StatusID    uuid.UUID `json:"status_id"`
	AssignedTo  []uuid.UUID `json:"assigned_to"`
	DueDate    time.Time   `json:"due_date"` // ISO 8601 format
}