package api

import (
	"encoding/json"
	"net/http"

	model "../model"
	"github.com/julienschmidt/httprouter"
)

// Tasks : List task
func Tasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tasks := model.Tasks{
		model.Task{Name: "Write presentation"},
		model.Task{Name: "Host meetup"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}
