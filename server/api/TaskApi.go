package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "../model"
	"github.com/julienschmidt/httprouter"
)

// API Main interface
type API struct {
	Context *model.AppContext
}

// Tasks : API/tasks
func (a *API) Tasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks := a.Context.Repo.TasksList()
	jsonResponse(w, tasks)
}
func jsonResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

// TaskByID : API/task/id
func (a *API) TaskByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID := ps.ByName("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}
	t := a.Context.Repo.TaskByID(id)
	jsonResponse(w, t)
}
