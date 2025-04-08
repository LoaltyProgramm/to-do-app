package api

import (
	"fmt"
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
		tasksHandler(w)
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", taskHandler)

}