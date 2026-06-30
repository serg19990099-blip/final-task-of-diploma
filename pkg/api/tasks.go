package api

import (
	"net/http"

	"final-task-of-diploma/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeInternalError(w)
		return
	}

	writeJSON(w, http.StatusOK, TasksResp{
		Tasks: tasks,
	})
}
