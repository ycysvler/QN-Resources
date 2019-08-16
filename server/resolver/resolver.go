package resolver

type Resolver struct{}

type ResourceQuery struct {
}

func (_ *Resolver) ResourceQuery() (*ResourceQuery, error) {
	return &ResourceQuery{}, nil
}
