package resolver

type Resolver struct{}

func (r *Resolver) Character() *CharacterResolver {

	ptr := new(string)
	*ptr = "name"

	return &CharacterResolver{}
}
