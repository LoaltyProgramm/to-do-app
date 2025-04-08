package api

import (
	"encoding/json"
	"fmt"
	"go1f/pkg/db"
	"io"
	"net/http"
	"time"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parse body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	var task db.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		writeJson(w, map[string]string{"error":"Unmarshal error"})
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{"error":"The title cannot be empty"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, map[string]string{"error":"The date was not validated"})
		return
	}

	id, err := db.AddTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, map[string]any{"id": id})
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(layout)
	}

	t, err := time.Parse(layout, task.Date)
	if err != nil {
		return fmt.Errorf("Error parse string in time.Time: %v", err)
	}

	var next string

	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	
	if comparingDate(t, now) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(layout)
		} else {
			task.Date = next
		}
	}

	return nil
}