package resolver

// graphql "github.com/graph-gophers/graphql-go"

type Resolver struct{}

func (r *Resolver) Videos(args VideoArgument) (*VideosResolver, error) {

	return &VideosResolver{}, nil
}

func (r *Resolver) Artists(args ArtistArgument) (*ArtistsResolver, error) {

	return &ArtistsResolver{}, nil
}
