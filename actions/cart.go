package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// Cart handler handle the user's shopping cart page.
func CartHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("cart.html"))
}
