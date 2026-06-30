package api

import (
	"errors"
	"net/http"

	"final-task-of-diploma/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "не указан идентификатор")
		return
	}

	err := db.DeleteTask(id)
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
