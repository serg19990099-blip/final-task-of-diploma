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
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, errors.New("задача не найдена"))
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, map[string]string{})
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.UpdateDate(id, next)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{})
}
