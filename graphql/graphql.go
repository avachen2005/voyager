package graphql

import (
	"net/http"

	"github.com/avachen2005/voyager/graphql/resolver"
	"github.com/graph-gophers/graphql-go/relay"

	graphql "github.com/graph-gophers/graphql-go"
)

func GraphqlHandler() http.Handler {
	s := graphql.MustParseSchema(GenerateSchema(), &resolver.Resolver{})
	return &relay.Handler{Schema: s}
}
