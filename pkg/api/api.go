package api

import (
	"encoding/json"
	"fmt"
	"go1f/pkg/db"
	"io"
	"net/http"
	"strconv"
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
	id := r.FormValue("id")

	switch r.Method {
	case http.MethodGet:
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
	case http.MethodDelete:
		if id == "" {
			writeJson(w, map[string]string{"error":"id is not empty"})
			return
		}

		_, err := strconv.Atoi(id)
		if err != nil {
			writeJson(w, map[string]string{"error":"Invalid id. The ID is passed as an integer"})
			return
		}

		err = db.DeletTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}
		writeJson(w, struct{}{})
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks(w, r)
}

func ComplitedHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error":err.Error()})
		return
	}

	if task.Repeat == "" {
		db.DeletTask(id)
		writeJson(w, struct{}{})
	} 

	if task.Repeat != "" {
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}
		task.Date = nextDate

		err = db.UpdateDateTask(task)
		if err != nil {
			writeJson(w, map[string]string{"error":err.Error()})
			return
		}
		writeJson(w, struct{}{})
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

}

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", ComplitedHandler)
}