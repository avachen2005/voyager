package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type ArtistResolver struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistsResolver struct {
	artists []*ArtistResolver
}

type ArtistArgument struct {
	Id   int
	Name string
}

func (r *ArtistsResolver) ID(args int) (graphql.ID, error) {
	return graphql.ID(1), nil
}

func (r *ArtistsResolver) NAME(args string) (string, error) {
	return "Name", nil
}
