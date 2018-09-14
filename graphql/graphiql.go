package graphql

import (
	"github.com/gobuffalo/buffalo"
)

func GraphiqlHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("graphiql.html"))
}
