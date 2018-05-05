package main

import (
	api "./api"
	"github.com/julienschmidt/httprouter"
)

// NewRouter : Construct new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/task", api.Tasks)
	router.GET("/tasks", api.Tasks)
	return router
}
