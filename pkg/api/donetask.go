package api

import (
	"errors"
	"net/http"
	"time"

	"final-task-of-diploma/pkg/db"
)

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			if errors.Is(err, db.ErrTaskNotFound) {
				writeError(w, http.StatusNotFound, "задача не найдена")
				return
			}

			writeInternalError(w)
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{})
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = db.UpdateDate(id, next)
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
