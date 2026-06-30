package api

import "net/http"

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		addTaskHandler(w, r)
		return

	case http.MethodGet:
		getTaskHandler(w, r)
		return

	case http.MethodPut:
		editTaskHandler(w, r)
		return

	case http.MethodDelete:
		deleteTaskHandler(w, r)
		return

	default:
		writeError(w, http.StatusMethodNotAllowed, "метод не поддерживается")
		return
	}
}
