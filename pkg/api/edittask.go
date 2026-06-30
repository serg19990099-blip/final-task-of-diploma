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
		writeError(w, http.StatusBadRequest, "не указан идентификатор")
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		if errors.Is(err, db.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, "задача не найдена")
			return
		}

		writeInternalError(w)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if task.ID == "" {
		writeError(w, http.StatusBadRequest, "не указан идентификатор")
		return
	}

	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "не указан заголовок задачи")
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		if errors.Is(err, db.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, "задача не найдена")
			return
		}

		writeInternalError(w)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{})
}
