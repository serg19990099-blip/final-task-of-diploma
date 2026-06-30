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
		writeError(w, err)
		return
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
