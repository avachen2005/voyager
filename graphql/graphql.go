package graphql

import (
	"fmt"
	"net/http"

	"github.com/kr/pretty"

	"github.com/avachen2005/voyager/graphql/resolver"
	"github.com/avachen2005/voyager/graphql/schema"
	"github.com/graph-gophers/graphql-go/relay"

	graphql "github.com/graph-gophers/graphql-go"
)

func GraphqlHandler() http.Handler {

	fmt.Printf("%# v", pretty.Formatter(schema.String()))

	s := graphql.MustParseSchema(schema.String(), &resolver.Resolver{})
	return &relay.Handler{Schema: s}
}
