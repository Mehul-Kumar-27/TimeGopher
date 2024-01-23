package models

import "time"

type Task struct {
	Name         string
	Description string
	Completed    bool
	StartTime    time.Time
}
