package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tris-tux/go-lis/backend/schema"
)

type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) GetAll() ([]schema.Task, error) {
	query := `
		SELECT *
		FROM task
		ORDER BY task_id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	taskList := []schema.Task{}
	for rows.Next() {
		var t schema.Task
		if err := rows.Scan(&t.TaskId, &t.Title, &t.AcctionTime, &t.CreateTime, &t.UpdateTime, &t.IdFinished); err != nil {
			return nil, err
		}
		taskList = append(taskList, t)
	}
	return taskList, nil
}

func (p *Postgres) Insert(task *schema.Task) (int, error) {
	query := `
		INSERT INTO task (task_id, title, acction_time, create_time, update_time, id_finished)
		VALUES(nextval('task_id'), $1, $2, $3)
		RETURNING task_id;
	`

	rows, err := p.DB.Query(query, task.Title, task.AcctionTime, task.CreateTime, task.UpdateTime, convertBoolToBit(task.IdFinished))
	if err != nil {
		return -1, err
	}

	var task_id int
	for rows.Next() {
		if err := rows.Scan(&task_id); err != nil {
			return -1, err
		}
	}
	return task_id, nil
}

// func (p *Postgres) Update(todo *schema.Todo) error {
// 	query := `
// 		UPDATE todo
// 		SET note = $2, done = $3
// 		WHERE id = $1;
// 	`

// 	rows, err := p.DB.Query(query, todo.ID, todo.Note, convertBoolToBit(todo.Done))
// 	if err != nil {
// 		return err
// 	}

// 	var id int
// 	for rows.Next() {
// 		if err := rows.Scan(&id); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (p *Postgres) Delete(task_id int) error {
	query := `
		DELETE FROM task
		WHERE task_id = $1;
	`

	if _, err := p.DB.Exec(query, task_id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	p.DB.Close()
}

func ConnectPostgres() (*Postgres, error) {
	connStr, err := loadPostgresConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func loadPostgresConfig() (string, error) {
	if os.Getenv("DB_HOST") == "" {
		return "", fmt.Errorf("Environment variable DB_HOST must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		return "", fmt.Errorf("Environment variable DB_PORT must be set")
	}
	if os.Getenv("DB_USER") == "" {
		return "", fmt.Errorf("Environment variable DB_USER must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return "", fmt.Errorf("Environment variable DB_DATABASE must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return "", fmt.Errorf("Environment variable DB_PASSWORD must be set")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	return connStr, nil
}

func convertBoolToBit(val bool) int {
	if val {
		return 1
	}
	return 0
}
