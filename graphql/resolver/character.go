package resolver

import graphql "github.com/graph-gophers/graphql-go"

type CharacterResolver struct {
	Character
}

type Character struct {
	ID   graphql.ID
	Name string
}

func (r *CharacterResolver) ID() graphql.ID {
	return graphql.ID("1")
}

func (r *CharacterResolver) Name() string {
	return "name"
}
