package req

import (
    "encoding/json"
)

// Default status values
const (
    StatusTodo    = "todo"
    StatusDone    = "done"
    StatusBacklog = "backlog"
)

type StatusUpdate struct {
    StatusName json.RawMessage `json:"status_name"` // Using RawMessage to allow for flexible status name formats
}

func GetDefaultStatuses() []string {
    return []string{StatusTodo, StatusDone, StatusBacklog}
}

func IsValidDefaultStatus(status string) bool {
    for _, defaultStatus := range GetDefaultStatuses() {
        if status == defaultStatus {
            return true
        }
    }
    return false
}