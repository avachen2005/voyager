package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Video struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type VideosResolver struct {
	videos []*Video
}

type VideoArgument struct {
	Id    int64
	Title string
}

func (r *VideosResolver) Id(args int) (graphql.ID, error) {
	return graphql.ID("1"), nil
}

func (r *VideosResolver) Title(args string) (string, error) {
	return "Title", nil
}
