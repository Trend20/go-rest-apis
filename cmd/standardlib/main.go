package main

import (
	"net/http"
)

func main() {

	//	initialize a server multiplexer
	mux := http.NewServeMux()

	//	register the routes and handlers
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &recipeHandler{})

	//	server listen on port 3000
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		return
	}
}

// homepage initializations
type homeHandler struct{}

// define the homeHandler method
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the homepage"))
}

// recipes initializations
type recipeHandler struct{}

// define the recipeHandler method
func (h *recipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the recipe handler method"))
}
