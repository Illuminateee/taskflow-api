package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tasks struct {
	ID          uuid.UUID  `json:"uuid"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	StatusID     Status  `json:"status_id"`
	IsMainTask   bool        `json:"is_main_task"`
	AssignedTo  []User  	`gorm:"many2many:task_assignment;"`
	ParentTask  uuid.UUID   `json:"parent_task"`
	CreatedBy   []User  	`gorm:"many2many:task_users;"`
	DueDate     time.Time   `json:"due_date"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func (Tasks) TableName() string {
	return "tasks"
}