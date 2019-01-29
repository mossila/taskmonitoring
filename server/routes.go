package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mossila/taskmonitoring/server/api"
	"github.com/mossila/taskmonitoring/server/model"
	"net/http"
)

// NewRouter : Construct new router
func NewRouter(c *model.AppContext) *httprouter.Router {
	app := api.API{Context: c}
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/task", app.Tasks)
	router.GET("/tasks", app.Tasks)
	router.GET("/task/:id", app.TaskByID)
	router.POST("/task", app.CreateTask)
	router.PUT("/task", app.UpdateTask)
	return router
}

// Index - temporary index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
