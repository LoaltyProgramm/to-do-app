package handlers

import (
	"net/http"
	"time"
	"fmt"
	"io"
	"encoding/json"
	"strconv"

	"github.com/LoaltyProgramm/to-do-app/internal/models"
	"github.com/LoaltyProgramm/to-do-app/internal/utils"
	"github.com/LoaltyProgramm/to-do-app/internal/service"
)

type TaskHandler struct {
	taskService service.TasksService
}

func NewTaskHandlers(taskService service.TasksService) TaskHandler {
	return TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) NextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeate := r.FormValue("repeat")

	now, err := time.Parse(utils.Layout, nowStr)
	if err != nil {
		fmt.Printf("Error parse now time: %v", err)
	}

	next, err := utils.NextDate(now, date, repeate)
	if err != nil {
		fmt.Printf("Couldn't find the next date: %v", err)
	}

	w.Write([]byte(next))
}

func (h *TaskHandler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	switch r.Method {
	case http.MethodGet:
		task, err := h.taskService.GetTaskByID(id)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":err.Error()})
			return
		}
		utils.WriteJson(w, task)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error parse body", http.StatusBadRequest)
		}
		defer r.Body.Close()

		var task models.Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":"Unmarshal error"})
			return
		}

		if task.Title == "" {
			utils.WriteJson(w, map[string]string{"error":"The title cannot be empty"})
			return
		}

		err = utils.CheckDate(&task)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":"The date was not validated"})
			return
		}

		id, err := h.taskService.CreateTask(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.WriteJson(w, map[string]any{"id": id})
	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":err.Error()})
			return
		}
		defer r.Body.Close()

		var task models.Task

		err = json.Unmarshal(body, &task)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":err.Error()})
			return
		}

		err = h.taskService.UpdateTask(task)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":err.Error()})
			return
		}
		utils.WriteJson(w, task)
	case http.MethodDelete:
		if id == "" {
			utils.WriteJson(w, map[string]string{"error":"id is not empty"})
			return
		}

		_, err := strconv.Atoi(id)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":"Invalid id. The ID is passed as an integer"})
			return
		}

		err = h.taskService.DeleteTaskByID(id)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":err.Error()})
			return
		}
		utils.WriteJson(w, struct{}{})
	}
}

func (h *TaskHandler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	valueSearch := r.FormValue("search")
	limit := 50
	if valueSearch == "" {
		tasks, err := h.taskService.ListTasks(limit)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":"Error select from table"})
			return
		}

		utils.WriteJson(w, models.TasksResp{
			Tasks: tasks,
		})
		return
	}

	if valueSearch != "" {
		if utils.CheckingTheDateUsingATemplate(valueSearch) {
			layoutInput := "02.01.2006"
			dateTime, _ := time.Parse(layoutInput, valueSearch)
			date := dateTime.Format(utils.Layout)

			tasks, err := h.taskService.FindTasksByDate(date, 50)
			if err != nil {
				utils.WriteJson(w, map[string]string{"error":err.Error()})
				return
			}

			utils.WriteJson(w, models.TasksResp{
				Tasks: tasks,
			})
			return
		}

		tasks, err := h.taskService.SearchTasks(valueSearch, 50)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error":err.Error()})
			return
		}

		utils.WriteJson(w, models.TasksResp{
			Tasks: tasks,
		})
		return
	}
}

func (h *TaskHandler) ComplitedHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	task, err := h.taskService.GetTaskByID(id)
	if err != nil {
		utils.WriteJson(w, map[string]string{"error": err.Error()})
		return
	}

	if task.Repeat == "" {
		h.taskService.DeleteTaskByID(id)
		utils.WriteJson(w, struct{}{})
	}

	if task.Repeat != "" {
		nextDate, err := utils.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error": err.Error()})
			return
		}
		task.Date = nextDate

		err = h.taskService.UpdateTaskDate(task)
		if err != nil {
			utils.WriteJson(w, map[string]string{"error": err.Error()})
			return
		}
		utils.WriteJson(w, struct{}{})
	}
} 

func (h *TaskHandler) InitHandler() {
	http.HandleFunc("/api/nextdate", h.NextDayHandler)
	http.HandleFunc("/api/task", h.TaskHandler)
	http.HandleFunc("/api/tasks", h.TasksHandler)
	http.HandleFunc("/api/task/done", h.ComplitedHandler)
}