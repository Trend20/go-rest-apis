package recipes

// Recipe struct
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Individual Ingredient struct
type Ingredient struct {
	Name string `json:"name"`
}
