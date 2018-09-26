package graphql

import (
	"net/http"

	"github.com/avachen2005/voyager/graphql/resolver"
	"github.com/avachen2005/voyager/graphql/schema"
	"github.com/graph-gophers/graphql-go/relay"

	graphql "github.com/graph-gophers/graphql-go"
)

func GraphqlHandler() http.Handler {
	s := graphql.MustParseSchema(schema.String(), &resolver.Resolver{})
	return &relay.Handler{Schema: s}
}
