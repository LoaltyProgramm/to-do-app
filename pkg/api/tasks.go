package api

import (
	"go1f/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJson(w, map[string]string{"error":"Error select from table"})
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

