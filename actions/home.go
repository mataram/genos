package actions

import "github.com/gobuffalo/buffalo"
import "github.com/mataram/genos/models"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	breadcrumbs := []models.Breadcrumb{}
	breadcrumbs = append(breadcrumbs, models.Breadcrumb{"/", "Home"})
	c.Set("breadcrumbs", breadcrumbs)
	return c.Render(200, r.HTML("index.html"))
}
