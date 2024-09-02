package main

import (
	"encoding/json"
	"github.com/Trend20/go-rest-apis/pkg/recipes"
	"github.com/gosimple/slug"
	"net/http"
	"regexp"
)

func main() {

	//	initialize a server multiplexer
	mux := http.NewServeMux()

	// Create the Store and Recipe Handler
	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)

	//	register the routes and handlers
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", recipesHandler)
	mux.Handle("/recipes/", recipesHandler)

	//	server listen on port 3000
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		return
	}
}

// homeHandler initialization
type homeHandler struct{}

// define the homeHandler method
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the homepage"))
}

// ERROR HANDLING
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

// recipe store interface
type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}

// RecipesHandler initialization
type RecipesHandler struct {
	store recipeStore
}

func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

// create function handlers for different routes

// CREATE RECIPE
func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	// Recipe object that will be populated from JSON payload
	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	// Convert the name of the recipe into URL friendly string
	resourceID := slug.Make(recipe.Name)

	// Call the store to add the recipe
	if err := h.store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	// Set the status code to 200
	w.WriteHeader(http.StatusOK)
}

// GET ALL RECIPES
func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {}

// GET RECIPE BY ID
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {}

// UPDATE RECIPE
func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}

// DELETE RECIPE
func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

// REGEX CHECKER
var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

// switch the RecipesHandler method based on the HTTP VERB and the route url
func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}
}
