package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"final-task-of-diploma/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, errors.New("задача не найдена"))
		return
	}

	writeJSON(w, task)
}

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, err)
		return
	}

	if task.ID == "" {
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	if task.Title == "" {
		writeError(w, errors.New("не указан заголовок задачи"))
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{})
}
