package api

import (
	"go1f/pkg/db"
	"net/http"
	"time"
)

type TasksResp struct {
	Tasks []db.Task `json:"tasks"`
}

func tasks(w http.ResponseWriter, r *http.Request) {
	valueSearch := r.FormValue("search")
	if valueSearch == "" {
		tasks, err := db.GetTasks(50)
		if err != nil {
			writeJson(w, map[string]string{"error":"Error select from table"})
			return
		}

		writeJson(w, TasksResp{
			Tasks: tasks,
		})
		return
	}

	if valueSearch != "" {
		if isValudDate(valueSearch) {
			layoutInput := "02.01.2006"
			dateTime, _ := time.Parse(layoutInput, valueSearch)
			date := dateTime.Format(layout)

			tasks, err := db.SearchTasksDates(date, 50)
			if err != nil {
				writeJson(w, map[string]string{"error":err.Error()})
				return
			}

			writeJson(w, TasksResp{
				Tasks: tasks,
			})
			return
		}
	}

	
	
}

func isValudDate(data string) bool {
	layoutInput := "02.01.2006"
	_, err := time.Parse(layoutInput, data)
	return err == nil
}

