package models

import (
	"time"
)

type Task struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SelectTaskList() (*[]Task, error) {
	var db, _ = DbConn()

	sql := "select * from tasks;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tasks []Task
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := Task{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return &tasks, nil
}

func FindTaskById(id uint64) (*Task, error) {
	var db, _ = DbConn()

	sql := "select * from tasks where id = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	task := Task{}
	if err := stmt.QueryRow(id).Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(title string, content string) (*Task, error) {
	var db, _ = DbConn()

	sql := "insert into tasks(title, content) values(?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// insert
	res, err := stmt.Exec(title, content)
	if err != nil {
		return nil, err
	}

	// get inserted task
	id, _ := res.LastInsertId()
	task, _ := FindTaskById((uint64(id)))

	return task, nil
}