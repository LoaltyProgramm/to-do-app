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
		return Task{}, fmt.Errorf("id not specified")
	}

	err := DB.Get(&task, query, id)
	if err != nil {
		return Task{}, fmt.Errorf("there is no given task: %v", err)
	}

	return task, nil
}

func UpdateTask(task Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? 
	WHERE id = ?;`

	if task.Id == "" {
		return fmt.Errorf("id not specified")
	}

	taskInt, err := strconv.Atoi(task.Id)
	if err != nil {
		return fmt.Errorf("the Id can only be a number: %v", err)
	}

	if taskInt > 1000000 {
		return fmt.Errorf("the ID is too large")
	}

	if task.Comment == "" || task.Title == "" {
		return fmt.Errorf("required fields must be filled in")
	}

	_, err = DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.Id)
	if err != nil {
		return fmt.Errorf("error update task: %v", err)
	}

	return nil
}

func UpdateDateTask(task Task) error {
	query := `UPDATE scheduler SET date = ? 
	WHERE id = ?;`

	if task.Id == "" {
		return fmt.Errorf("id not specified")
	}

	taskInt, err := strconv.Atoi(task.Id)
	if err != nil {
		return fmt.Errorf("the Id can only be a number: %v", err)
	}

	if taskInt > 1000000 {
		return fmt.Errorf("the ID is too large")
	}

	if task.Title == "" {
		return fmt.Errorf("required fields must be filled in")
	}

	_, err = DB.Exec(query, task.Date, task.Id)
	if err != nil {
		return fmt.Errorf("error update task: %v", err)
	}

	return nil
}

func DeletTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`

	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("fail delete task for id")
	}

	return nil
}

func SearchTasksDates(date string, limit int) ([]Task, error) {
	query := `SELECT * FROM scheduler WHERE date = ? LIMIT ?;`

	var tasks []Task

	err := DB.Select(&tasks, query, date, limit)
	if err != nil {
		return nil, fmt.Errorf("the request to get rows by date could not be completed: %v", err)
	}

	return tasks, nil
}

func SearchTasks(data string, limit int) ([]Task, error) {
	query := `SELECT * FROM scheduler WHERE title LIKE '%' || ? || '%'
			OR comment LIKE '%' || ? ||'%' LIMIT ?;`

	var tasks []Task

	err := DB.Select(&tasks, query, data, data, limit)
	if err != nil {
		return nil, fmt.Errorf("the request to get rows by date could not be completed: %v", err)
	}

	return tasks, nil
}
