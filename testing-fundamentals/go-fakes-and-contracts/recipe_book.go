package gofakesandcontracts

type RecipeBook interface {
	GetRecipes() ([]Recipe, error)
	AddRecipes(...Recipe) error
}

// This is not correct!
