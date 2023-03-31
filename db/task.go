package db

import (
	"database/sql"

	"github.com/kallepan/go-backend/models"
)

func (db Database) GetAllTasks(page int, page_size int) (*models.TaskList, error) {
	list := &models.TaskList{}

	rows, err := db.CON.Query("SELECT * FROM tasks ORDER BY id,name LIMIT $1 OFFSET $2", page_size, page_size*(page-1))
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return list, err
		}

		list.Tasks = append(list.Tasks, task)
	}

	return list, nil
}

func (db Database) AddTask(task *models.Task) error {
	var id int
	var createdAt string

	query := "INSERT INTO tasks (name, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := db.CON.QueryRow(query, task.Name, task.Description, task.Completed).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	task.ID = id
	task.CreatedAt = createdAt

	return nil
}

func (db Database) GetTask(id int) (models.Task, error) {
	task := models.Task{}

	query := "SELECT * FROM tasks WHERE id = $1"
	row := db.CON.QueryRow(query, id)
	switch err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt); err {
	case sql.ErrNoRows:
		return task, ErrNoMatch
	default:
		return task, err
	}
}

func (db Database) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := db.CON.Exec(query, id)

	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateTask(id int, taskData models.Task) (models.Task, error) {
	task := models.Task{}

	query := "UPDATE tasks SET name = $1, description = $2, completed = $3 WHERE id = $4 RETURNING name, description, created_at, completed"
	row := db.CON.QueryRow(query, taskData.Name, taskData.Description, taskData.Completed, id)
	switch err := row.Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.Completed); err {
	case sql.ErrNoRows:
		return task, ErrNoMatch
	default:
		return task, err
	}
}
