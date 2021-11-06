package task

import "time"

type Task struct {
	ID        uint64
	StartedAt *time.Time
}
