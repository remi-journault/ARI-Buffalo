package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// Profil handler handle the user's profile page
func ProfileHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("profile.html"))
}
