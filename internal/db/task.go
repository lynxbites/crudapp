package db

import (
	"time"
)

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (taskPost TaskPost) Valid() bool {
	if taskPost.Title == "" || taskPost.Description == "" {
		return false
	}
	return true
}

type TaskPut struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (taskPut TaskPut) Valid() bool {
	if taskPut.Title == "" || taskPut.Description == "" || taskPut.Status == "" {
		return false
	}
	if taskPut.Status != "new" && taskPut.Status != "in_progress" && taskPut.Status != "done" {
		return false
	}
	return true
}
