package server

import (
	"net/http"

	"final-task-of-diploma/pkg/api"
)

const (
	defaultPort = "7540"
	webDir      = "./web"
)

func Run() error {
	http.HandleFunc("/api/nextdate", api.NextDateHandler)
	http.HandleFunc("/api/task", api.TaskHandler)
	http.HandleFunc("/api/tasks", api.TasksHandler)
	http.HandleFunc("/api/task/done", api.DoneTaskHandler)

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	return http.ListenAndServe(":"+defaultPort, nil)
}
