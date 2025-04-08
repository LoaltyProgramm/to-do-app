package db

import (
	"fmt"
	"strconv"
)

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (?, ?, ?, ?);
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func GetTasks(limit int) ([]Task, error) {
	query := `SELECT * FROM scheduler ORDER BY date ASC LIMIT ?;`

	tasks := []Task{}

	err := DB.Select(&tasks, query, limit)
	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func GetTask(id string) (Task, error) {
	query := `SELECT * FROM scheduler WHERE id=?;`

	task := Task{}

	if id == "" {
		return Task{}, fmt.Errorf("Id not specified")
	}

	err := DB.Get(&task, query, id)
	if err != nil {
		return Task{}, fmt.Errorf("There is no given task: %v", err)
	}

	return task, nil
}

func UpdateTask(task Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? 
	WHERE id = ?;`

	if task.Id == "" {
		return fmt.Errorf("Id not specified")
	}

	taskInt, err := strconv.Atoi(task.Id)
	if err != nil {
		return fmt.Errorf("The Id can only be a number: %v", err)
	}

	if taskInt > 1000000 {
		return fmt.Errorf("The ID is too large: %v", err)
	}

	if task.Comment == "" || task.Title == "" {
		return fmt.Errorf("Required fields must be filled in: %v", err)
	}

	_, err = DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.Id)
	if err != nil {
		return fmt.Errorf("Error update task: %v", err)
	}

	return nil
}