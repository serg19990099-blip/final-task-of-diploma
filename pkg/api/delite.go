package api

import (
	"errors"
	"net/http"

	"final-task-of-diploma/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{})
}
