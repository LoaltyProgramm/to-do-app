package api

import (
	"encoding/json"
	"fmt"
	"go1f/pkg/db"
	"io"
	"net/http"
	"time"
)

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeate := r.FormValue("repeat")

	now, err := time.Parse(layout, nowStr)
	if err != nil {
		fmt.Printf("Error parse now time: %v", err)
	}

	next, err := NextDate(now, date, repeate)
	if err != nil {
		fmt.Printf("Couldn't find the next date: %v", err)
	}

	w.Write([]byte(next))
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.FormValue("id")
		task, err := db.GetTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}
		writeJson(w, task)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}
		defer r.Body.Close()

		var task db.Task

		err = json.Unmarshal(body, &task)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}

		err = db.UpdateTask(task)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}
		writeJson(w, db.Task{})
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks(w)
}

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
}