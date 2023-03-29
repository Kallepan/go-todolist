package models

import (
	"fmt"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func (t *Task) Bind(r *http.Request) error {
	if t.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (t *Task) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (tl *TaskList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
