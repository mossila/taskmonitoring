package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "../model"
	util "../util"
	"github.com/julienschmidt/httprouter"
)

// API Main interface
type API struct {
	Context *model.AppContext
}

// Tasks : GET API/tasks
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

// TaskByID : GET API/task/id
func (a *API) TaskByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID := ps.ByName("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}
	t := a.Context.Repo.TaskByID(id)
	jsonResponse(w, t)
}

// CreateTask : POST API/task
func (a *API) CreateTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var t model.Task
	err := util.Binding(r, &t)
	if err != nil {
		panic(err)
	}
	nt := a.Context.Repo.CreateTask(t)
	jsonResponse(w, nt)
}

// UpdateTask : PUT API/task
func (a *API) UpdateTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var t model.Task
	err := util.Binding(r, &t)
	if err != nil {
		panic(err)
	}
	nt := a.Context.Repo.UpdateTask(t)
	jsonResponse(w, nt)
}
