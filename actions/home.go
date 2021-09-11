package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/project/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("pizzas", []models.Pizza{
		{9, "Margarita"},
		{9.5, "Regina"},
		{9, "Pepina"},
		{10, "Reine"},
		{15, "Triple super kebab"},
		{10, "Tournée du chef"},
	})
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
