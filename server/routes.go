package main

import (
	api "./api"
	model "./model"
	"github.com/julienschmidt/httprouter"
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
	return router
}
