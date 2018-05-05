package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	model "./model"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	c := &model.AppContext{Repo: model.InitialRepo()}
	router := NewRouter(c)
	n := negroni.Classic() // Include Logger, Static, Recovery
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":8080", n))
}

// Index : Http server index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
